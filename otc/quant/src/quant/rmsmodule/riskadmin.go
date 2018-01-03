package main

import (
	"encoding/json"
	"fmt"
	zmq "github.com/pebbe/zmq3"
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
	emsbase "quant/emsmodule/base"
	"quant/helper"
	"quant/rmsmodule/base"
	"strconv"
	"time"
	"util/csp"
	"util/db"
)

const (
	queryrulesql   string = "select ID,RiskType,InstrumentID,Account,StrategyName,Indicator,`Condition`,Threshold,Action from jgriskrule"
	addrulesql     string = "insert into jgriskrule(ID, RiskType,InstrumentID,Account,StrategyName,Indicator,`Condition`,Threshold,Action, ModifyTimeStamp,CreateTimeStamp) VALUES(?, ?,?,?,?,?,?,?,?,?,?)"
	addrulejoursql string = "insert into jgriskrulejour(Date, Time, OperatorNo, OldRule, NewRule, Remark) VALUES(?,?,?,?,?,?)"
	updaterulesql  string = "update jgriskrule set RiskType = ?,InstrumentID = ?,Account = ?,StrategyName = ?,Indicator = ?,`Condition` = ?,Threshold = ?,Action = ? where ID = ?"
	delrulesql     string = "delete from jgriskrule where ID = '%s'"
)

type riskNoRep struct{ ID string }

func getRiskNo() string {
	var NO int
	for _, r := range riskrules {
		idx, _ := strconv.Atoi(r.ID[1:len(r.ID)])
		if idx > NO {
			NO = idx
		}
	}
	NO = NO + 1
	return fmt.Sprintf("R%04d", NO)
}

type riskAdmin struct {
	conf     *goini.Config
	pushAddr string
	dbop     *db.MysqlWorker
	sqls     chan string
	push     *zmq.Socket
}

func (r *riskAdmin) init() {
	log.Info("RiskAdmin start.")
	r.conf = goini.SetConfig(helper.QuantConfigFile)
	r.connectdb()
	r.cacheRiskRules()
	r.connectToPublisher()
}

func (r *riskAdmin) connectToPublisher() {
	r.pushAddr = r.conf.GetStr("riskmodule", "pull_addr")
	r.push, _ = zmq.NewSocket(zmq.PUSH)
	r.push.Connect(r.pushAddr)
	log.Info("RMS riskadmin connect to msgrouter publish service")
}

func (r *riskAdmin) connectdb() {
	config := db.MysqlConfig{
		MysqlUsernName: r.conf.GetStr(helper.ConfigMysqlSessionName, helper.ConfigMysqlUserName),
		MysqlPwd:       r.conf.GetStr(helper.ConfigMysqlSessionName, helper.ConfigMysqlPwd),
		MysqlURL:       r.conf.GetStr(helper.ConfigMysqlSessionName, helper.ConfigMysqlUrl),
	}
	r.sqls = make(chan string, 1000)
	r.dbop = &db.MysqlWorker{SQLs: r.sqls, MysqlConfig: &config}
	err := r.dbop.Init()
	if err != nil {
		log.Error("RMS connect to mysql fail. mysqlurl: %s. Error:%s.", r.dbop.MysqlURL, err.Error())
	} else {
		log.Info("RMS connect to %s mysql.", r.dbop.MysqlURL)
	}
	go r.dbop.Run()
}

func (r *riskAdmin) cacheRiskRules() {
	rows, err := r.dbop.DB.Query(queryrulesql)
	defer rows.Close()
	if err != nil {
		log.Error(err)
	} else {
		for rows.Next() {
			rule := rmsbase.RiskRule{}
			err := rows.Scan(&rule.ID, &rule.RiskType, &rule.InstrumentID, &rule.Account, &rule.StrategyName,
				&rule.Indicator, &rule.Condition, &rule.Threshold, &rule.Action)
			if err != nil {
				log.Error(err)
			}
			riskrules[rule.ID] = rule
		}
		err = rows.Err()
		if err != nil {
			log.Error(err)
		}
	}
}

func (r *riskAdmin) generateRiskRule(req *csp.Request, riskno string) (rule rmsbase.RiskRule) {
	ts, _ := strconv.ParseFloat(req.PARAMS[6], 32)
	act, _ := strconv.Atoi(req.PARAMS[7])
	rule = rmsbase.RiskRule{
		ID:           riskno,
		RiskType:     req.PARAMS[0],
		InstrumentID: req.PARAMS[1],
		Account:      req.PARAMS[2],
		StrategyName: req.PARAMS[3],
		Indicator:    req.PARAMS[4],
		Condition:    req.PARAMS[5],
		Threshold:    float32(ts),
		Action:       act}
	return
}

func (r *riskAdmin) addRiskRules(req *csp.Request) (rep csp.Response) {
	csp.SetRepV(req, &rep)

	if len(req.PARAMS) != 10 {
		rep.RET = -1
		rep.MSG = fmt.Sprintf("RMS addRiskRules need 10 param, but get %d", len(req.PARAMS))
		log.Error(rep.MSG)
		return
	}

	rule := r.generateRiskRule(req, getRiskNo())

	tx, err := r.dbop.DB.Begin()
	if err != nil {
		rep.RET = -1
		rep.MSG = "RMS DBWorker get context failed. Error:" + err.Error()
		log.Error(rep.MSG)
		return
	}
	// write record into jgriskrule
	_, err1 := tx.Exec(addrulesql, rule.ID, rule.RiskType, rule.InstrumentID, rule.Account, rule.StrategyName,
		rule.Indicator, rule.Condition, rule.Threshold, rule.Action, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	// write record into jgriskrulejour
	remark := req.PARAMS[9]
	if remark == "" {
		remark = fmt.Sprintf("add by %s", req.PARAMS[8])
	}
	srule, _ := json.Marshal(rule)
	_, err2 := tx.Exec(addrulejoursql, time.Now().Format("2006-01-02"), time.Now().Format("15:04:05"), req.PARAMS[8], nil, string(srule), remark)
	if err1 != nil || err2 != nil {
		if err1 != nil {
			err = err1
		} else {
			err = err2
		}
		rep.RET = -1
		rep.MSG = "RMS DBWorker exec fail. " + err.Error()
		log.Error(rep.MSG)
		return
	}
	// commit
	err = tx.Commit()
	if err != nil {
		rep.RET = -1
		rep.MSG = "RMS DBWorker commit fail. " + err.Error()
		log.Error(rep.MSG)
	}
	// cache riskrule in mem
	riskrules[rule.ID] = rule

	// push new riskrule to MsgRouter
	brule, _ := msgpack.Marshal(rule)
	bdat, _ := msgpack.Marshal(csp.PubMsg{Topic: "addRiskRules", Msg: brule})
	log.Info("addRiskRules pub %s", string(bdat))
	r.push.SendBytes(bdat, 1)

	// return risk ID
	rep.DAT, _ = msgpack.Marshal(riskNoRep{ID: rule.ID})
	return
}

func (r *riskAdmin) delRiskRules(req *csp.Request) (rep csp.Response) {
	csp.SetRepV(req, &rep)

	if len(req.PARAMS) != 1 {
		rep.RET = -1
		rep.MSG = fmt.Sprintf("RMS delRiskRules need 1 param, but get %d", len(req.PARAMS))
		log.Error(rep.MSG)
		return
	}
	ID := req.PARAMS[0]
	if _, ok := riskrules[ID]; !ok {
		rep.RET = -1
		rep.MSG = fmt.Sprintf("RMS delRiskRules can't find '%s'", ID)
		return
	}

	r.sqls <- fmt.Sprintf(delrulesql, req.PARAMS[0])
	delete(riskrules, req.PARAMS[0])

	// push deleted riskrule  ID to MsgRouter
	delID, _ := msgpack.Marshal(riskNoRep{ID: req.PARAMS[0]})
	bdat, _ := msgpack.Marshal(csp.PubMsg{Topic: "delRiskRules", Msg: delID})
	log.Info("delRiskRules pub %s", string(bdat))
	r.push.SendBytes(bdat, 1)

	return
}

func (r *riskAdmin) modifyRiskRules(req *csp.Request) (rep csp.Response) {
	csp.SetRepV(req, &rep)

	if len(req.PARAMS) != 11 {
		rep.RET = -1
		rep.MSG = fmt.Sprintf("RMS modifyRiskRules need 11 param, but get %d", len(req.PARAMS))
		log.Error(rep.MSG)
		return
	}
	rule := r.generateRiskRule(req, req.PARAMS[10])
	if _, ok := riskrules[rule.ID]; !ok {
		rep.RET = -1
		rep.MSG = fmt.Sprintf("RMS modifyRiskRules can't find '%s'", rule.ID)
		return
	}
	tx, err := r.dbop.DB.Begin()
	if err != nil {
		rep.RET = -1
		rep.MSG = "RMS DBWorker get context failed. Error:" + err.Error()
		log.Error(rep.MSG)
		return
	}
	_, err = tx.Exec(updaterulesql, rule.RiskType, rule.InstrumentID, rule.Account, rule.StrategyName,
		rule.Indicator, rule.Condition, rule.Threshold, rule.Action, rule.ID)
	if err != nil {
		rep.RET = -1
		rep.MSG = "RMS DBWorker exec fail. " + err.Error()
		log.Error(rep.MSG)
		return
	}
	// commit
	err = tx.Commit()
	if err != nil {
		rep.RET = -1
		rep.MSG = "RMS DBWorker commit fail. " + err.Error()
		log.Error(rep.MSG)
	}
	// cache riskrule in mem
	riskrules[rule.ID] = rule

	// push modified riskrule to MsgRouter
	brule, _ := msgpack.Marshal(rule)
	bdat, _ := msgpack.Marshal(csp.PubMsg{Topic: "modifyRiskRules", Msg: brule})
	log.Info("modifyRiskRules pub %s", string(bdat))
	r.push.SendBytes(bdat, 1)

	return
}

func (r *riskAdmin) getRiskRules(req *csp.Request) (rep csp.Response) {
	csp.SetRepV(req, &rep)
	rep.DAT, _ = msgpack.Marshal(riskrules)
	return
}

func (r *riskAdmin) getEntrustAmountInfo(req *csp.Request) (rep csp.Response) {
	csp.SetRepV(req, &rep)
	return
}

func (r *riskAdmin) getTradePct(req *csp.Request) (rep csp.Response) {
	csp.SetRepV(req, &rep)
	return
}

// CheckPort is used by EMSModule when an execution order passed by
func CheckPort(*emsbase.Portfolio) (bool, error) {
	return true, nil
}
