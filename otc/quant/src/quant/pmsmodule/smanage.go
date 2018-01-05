package main

import (
	"fmt"
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
	"quant/helper"
	"quant/pmsmodule/base"
	"quant/pmsmodule/option"
	"time"
	"util/csp"
	"util/db"
)

var (
	queryrulesql           = "select StrategyName,Author,AnnRet,SR,Vol,Calmar,MDD1,MDD1SD,MDD1ED,MDD2,MDD2SD,MDD2ED,Reamrk,CreateDateTime from jqstrategy"
	strategyMap            = make(map[string]func([]string) (pmsbase.IStrategy, error))
	runningStrategyMap     = make(map[string]pmsbase.IStrategy) // StrategyID string: strategyname_account_securityid
	pmsfuncrouter          = make(map[string]func(*csp.Request) csp.Response)
	strategytempMap        = make(map[string]pmsbase.StrategyTemp)
	runningStrategyInfoMap = make(map[string]*pmsbase.StrategyRunningInfo)
)

func newSmanage() *smanage {
	s := smanage{}
	s.init()
	return &s
}

type smanage struct {
	conf    *goini.Config
	repAddr string
	s       *csp.RepService
	dbop    *db.MysqlWorker
	sqls    chan string
}

func (sm *smanage) init() {

	// Create Rep Server
	log.Info("PMS smanage start")
	sm.conf = goini.SetConfig(helper.QuantConfigFile)
	sm.repAddr = sm.conf.GetStr("pmsmodule", "rep_addr")
	csp.NewRepService(sm.repAddr, sm)
	sm.setRouteMap()
	log.Info("RMS smanage bind %s", sm.repAddr)

	log.Info("PMS smanage register \"DeltaHedge\" strategy")
	strategyMap["DeltaHedge"] = option.NewDeltaHedge

	sm.connectdb()
	sm.cacheStrategyTemp()
	go sm.updateTradeStatus()
}

func (sm *smanage) setRouteMap() {
	// Init function router map
	pmsfuncrouter["newStrategy"] = sm.newStrategy
	pmsfuncrouter["newStrategyBatch"] = sm.newStrategyBatch
	pmsfuncrouter["stopStrategy"] = sm.stopStrategy
	pmsfuncrouter["getStrategyTemp"] = sm.getStrategyTemp
	pmsfuncrouter["getStrategyInfo"] = sm.getStrategyInfo
}

func (sm *smanage) connectdb() {
	config := db.MysqlConfig{
		MysqlUsernName: sm.conf.GetStr(helper.ConfigMysqlSessionName, helper.ConfigMysqlUserName),
		MysqlPwd:       sm.conf.GetStr(helper.ConfigMysqlSessionName, helper.ConfigMysqlPwd),
		MysqlURL:       sm.conf.GetStr(helper.ConfigMysqlSessionName, helper.ConfigMysqlUrl),
	}
	sm.sqls = make(chan string, 1000)
	sm.dbop = &db.MysqlWorker{SQLs: sm.sqls, MysqlConfig: &config}
	err := sm.dbop.Init()
	if err != nil {
		log.Error("PMS smanage connect to mysql fail. mysqlurl: %s. Error:%s.", sm.dbop.MysqlURL, err.Error())
	} else {
		log.Info("PMS smanage connect to %s mysql.", sm.dbop.MysqlURL)
	}
}

func (sm *smanage) cacheStrategyTemp() {
	rows, err := sm.dbop.DB.Query(queryrulesql)
	defer rows.Close()
	if err != nil {
		log.Error(err)
	} else {
		for rows.Next() {
			s := pmsbase.StrategyTemp{}
			err = rows.Scan(&s.StrategyName, &s.Author, &s.AnnRet, &s.SR, &s.Vol, &s.Calmar, &s.MDD1, &s.MDD1SD, &s.MDD1ED, &s.MDD2, &s.MDD2SD, &s.MDD2ED, &s.Reamrk, &s.CreateDateTime)
			if err != nil {
				log.Error(err)
			}
			strategytempMap[s.StrategyName] = s
		}
		err = rows.Err()
		if err != nil {
			log.Error(err)
		}
	}
}

// HandleReq this function will not be called.
func (sm *smanage) HandleReq(req csp.Request) (rep csp.Response) {
	return
}

// HandleBReq deal with monitor routed request
func (sm *smanage) HandleBReq(breq []byte) (brep []byte) {
	log.Info("PMS smanage receive %s ", string(breq))
	var req csp.Request
	msgpack.Unmarshal(breq, &req)
	if handle, ok := pmsfuncrouter[req.CMD]; ok {
		rep := handle(&req)
		brep, _ = msgpack.Marshal(rep)
	} else {
		var rep csp.Response
		csp.SetRepV(&req, &rep)
		rep.MSG = "PMS smanage can't route to '" + req.CMD + "' cmd"
		rep.RET = -1
		brep, _ = msgpack.Marshal(rep)
		log.Error(rep.MSG)
	}
	return
}

func (sm *smanage) updateTradeStatus() {
	for {
		for k, v := range runningStrategyInfoMap {
			if s, ok := runningStrategyMap[k]; ok {
				v.TradeStatus = s.GetTradeStatus()
			} else {
				log.Error("PMS updateTradeStatus can't find %s strategy", k)
			}
			time.Sleep(time.Second)
		}
	}
}

func (sm *smanage) createStrategy(name string, config []string) (string, error) {
	log.Info("PMS smanage new '%s' strategy", name)
	NewFunc, ok := strategyMap[name]
	if !ok {
		// rep.RET = -1
		errormsg := fmt.Sprintf("PMS smanage please register '%' strategy first", name)
		log.Error(errormsg)
		return "", fmt.Errorf(errormsg)
	}

	s, err := NewFunc(config)
	if err != nil {
		errormsg := err.Error()
		log.Error(errormsg)
		return "", fmt.Errorf(errormsg)
	}

	s.Start()
	ID := s.GetID()
	runningStrategyMap[ID] = s
	runningStrategyInfoMap[ID] = &pmsbase.StrategyRunningInfo{StrategyID: ID,
		Account:      s.GetAccount(),
		StrategyName: name,
		SecurityID:   s.GetSecurityID(),
		RunStatus:    "RUNNING",
		TradeStatus:  "-",
	}
	return ID, nil
}

//NewStrategy create strategy by name)
//params[0]: strategyname
//params[1]: RiskAversionRation
//params[2]: TradeCost
//params[3]: NotionalPrincipal
//params[4]: SubQuoteCodes
//params[5]: AdapterName
//params[6]: AccountID
//params[7]: CombiNo
func (sm *smanage) newStrategy(req *csp.Request) (rep csp.Response) {
	csp.SetRepV(req, &rep)

	ID, err := sm.createStrategy(req.PARAMS[0], req.PARAMS)
	if err != nil {
		rep.RET = -1
		rep.MSG = err.Error()
		return
	}
	// return strategyid
	rep.DAT, _ = msgpack.Marshal(pmsbase.StrategyID{ID: ID})
	return
}

func (sm *smanage) newStrategyBatch(req *csp.Request) (rep csp.Response) {
	csp.SetRepV(req, &rep)
	var ret []pmsbase.StrategyID
	for i := 0; i < len(req.PARAMS); i++ {
		var config []string
		err := msgpack.Unmarshal([]byte(req.PARAMS[i]), &config)
		if err != nil {
			rep.RET = -1
			rep.MSG = err.Error()
			log.Error(rep.MSG)
			return
		}
		ID, err := sm.createStrategy(config[0], config)
		if err != nil {
			ret = append(ret, pmsbase.StrategyID{Ret: -1, Msg: err.Error()})
		} else {
			ret = append(ret, pmsbase.StrategyID{ID: ID})
		}
	}
	rep.DAT, _ = msgpack.Marshal(ret)
	return
}

// StopStrategy stop strategy by strategyid
func (sm *smanage) stopStrategy(req *csp.Request) (rep csp.Response) {
	csp.SetRepV(req, &rep)
	SID := req.PARAMS[0]
	if s, ok := runningStrategyMap[SID]; ok {
		s.Stop()
		delete(runningStrategyMap, SID)
		delete(runningStrategyInfoMap, SID)
	} else {
		rep.RET = -1
		rep.MSG = fmt.Sprintf("PMS smanage can't find '%s' strategy", SID)
	}
	return
}

func (sm *smanage) getStrategyTemp(req *csp.Request) (rep csp.Response) {
	csp.SetRepV(req, &rep)
	rep.DAT, _ = msgpack.Marshal(strategytempMap)
	return
}

func (sm *smanage) getStrategyInfo(req *csp.Request) (rep csp.Response) {
	csp.SetRepV(req, &rep)
	rep.DAT, _ = msgpack.Marshal(runningStrategyInfoMap)
	return
}
