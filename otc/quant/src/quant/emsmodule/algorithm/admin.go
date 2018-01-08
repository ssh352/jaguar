package algorithm

import (
	"github.com/Workiva/go-datastructures/queue"
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
	"quant/emsmodule/adapter/ufxapi"
	emsbase "quant/emsmodule/base"
	"quant/helper"
	"strings"
	"time"
	"util/csp"
)

var (
	emsfuncrouter = make(map[string]func(*csp.Request) csp.Response)
	adaptersMap   = make(map[string]emsbase.ITarde)
	algorithmMap  = make(map[string]iAlgotithm)
	entrusts      = make(map[string]*emsbase.StrategyEntrust, 10000)
	excetions     = make(map[string]*emsbase.ExecutionOrder, 10000)
	traded        = make(map[string]bool, 10000)
)

// NewCommon  is used for create common algorithm pointer.
func newCommon() iAlgotithm {
	log.Info("EMS new common algorithm.")
	c := common{}
	c.init()
	return &c
}

func newTwap() iAlgotithm {
	log.Info("EMS new twap algorithm.")
	t := twap{}
	t.init()
	return &t
}

// Admin manage trade adapters and algorithm trades.
type Admin struct {
	Portqueue *queue.RingBuffer
	conf      *goini.Config
	repAddr   string
}

// Init start trade adapter and algorithm
func (admin *Admin) Init() {
	admin.conf = goini.SetConfig(helper.QuantConfigFile)
	adapters := admin.conf.GetStr(helper.ConfigEMSSessionName, helper.ConfigEMSTradeAdapter)
	for _, name := range strings.Split(adapters, "|") {
		if name == "UFX" {
			var itrade emsbase.ITarde = new(ufxapi.UFX)
			itrade.Init()
			adaptersMap[name] = itrade
		}
	}
	algorithmMap["COMMON"] = newCommon()
	algorithmMap["TWAP"] = newTwap()

	go admin.UpdateExcutionOrder()

	admin.repAddr = admin.conf.GetStr("emsmodule", "rep_addr")
	csp.NewRepService(admin.repAddr, admin)
	admin.setRouteMap()
}

func (admin *Admin) setRouteMap() {
	// Init function router map
	emsfuncrouter["getExecutionOrder"] = admin.getExecutionOrder
	emsfuncrouter["getEntrusts"] = admin.getEntrusts
}

func (admin *Admin) getExecutionOrder(req *csp.Request) (rep csp.Response) {
	csp.SetRepV(req, &rep)
	rep.DAT, _ = msgpack.Marshal(excetions)
	return
}

func (admin *Admin) getEntrusts(req *csp.Request) (rep csp.Response) {
	csp.SetRepV(req, &rep)
	rep.DAT, _ = msgpack.Marshal(entrusts)
	return
}

// HandleBReq deal with monitor routed request
func (admin *Admin) HandleBReq(breq []byte) (brep []byte) {
	log.Info("EMS algo admin receive %s ", string(breq))
	var req csp.Request
	msgpack.Unmarshal(breq, &req)
	if handle, ok := emsfuncrouter[req.CMD]; ok {
		rep := handle(&req)
		brep, _ = msgpack.Marshal(rep)
	} else {
		var rep csp.Response
		csp.SetRepV(&req, &rep)
		rep.MSG = "EMS algo admin can't route to '" + req.CMD + "' cmd"
		rep.RET = -1
		brep, _ = msgpack.Marshal(rep)
		log.Error(rep.MSG)
	}
	return
}

// Run will be called by emsmodule.
func (admin *Admin) Run() {
	for {
		if admin.Portqueue.Len() > 0 {
			port, err := admin.Portqueue.Get()
			if err == nil {
				p := port.(emsbase.Portfolio)
				algo, ok := algorithmMap[p.Algorithm]
				if ok {
					go algo.trade(p)
				} else {
					log.Error("EMS algorithm admin can't find \"%s\" algo.", p.Algorithm)
				}
			}
		} else {
			time.Sleep(time.Millisecond)
		}
	}
}

// UpdateExcutionOrder update cached execution orders
func (admin *Admin) UpdateExcutionOrder() {
	for {
		for _, ent := range entrusts {
			// entrusts key is order ref
			// excetions key is tacticid
			if e, ok := excetions[ent.TacticID]; ok {
				if e.Single {
					// case1 one tacticid -> one instrument
					// ent.EntrustStatus =
					// EntrustStatus    string  //委托状态
					e.DealAmount = ent.Entrust.DealAmount
					e.DealBalance = ent.Entrust.DealBalance
					e.DealPrice = ent.Entrust.DealPrice
				} else {
					// case2 one tacticid -> multi instruments
					_, ok := traded[ent.TacticID]
					if ent.Entrust.EntrustStatus == "7" && !ok {
						e.DealAmount = e.DealAmount + 1
						e.DealBalance = e.DealBalance + ent.Entrust.DealBalance
						traded[ent.TacticID] = true
					}
				}

				if e.EntrustAmount == e.DealAmount {
					e.EntrustStatus = "FINISH"
				}
			}
		}
	}
}
