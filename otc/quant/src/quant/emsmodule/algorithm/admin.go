package algorithm

import (
	"github.com/Workiva/go-datastructures/queue"
	log "github.com/thinkboy/log4go"
	"github.com/widuu/goini"
	"quant/emsmodule/adapter/ufxapi"
	emsbase "quant/emsmodule/base"
	"quant/helper"
	"strings"
	"time"
)

var (
	adaptersMap  = make(map[string]emsbase.ITarde)
	algorithmMap = make(map[string]iAlgotithm)
	entrusts     = make(map[string]*emsbase.StrategyEntrust, 10000)
	excetions    = make(map[string]*emsbase.ExecutionOrder, 10000)
)

// NewCommon  is used for create common algorithm pointer.
func newCommon() iAlgotithm {
	log.Info("EMS new common algorithm.")
	common := common{}
	common.init()
	return &common
}

// Admin manage trade adapters and algorithm trades.
type Admin struct {
	Portqueue *queue.RingBuffer
	conf      *goini.Config
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

// StrategyName     string  //策略名称
// TacticID         string  //策略ID
// Algorithm        string  //算法
// OperatorNo       string  //操作员
// StockCode        string  //证券代码
// AccountCode      string  //账户编号
// BusinessTime     int     //委托时间
// EntrustDirection string  //委托方向
// EntrustAmount    int     //委托数量
// DealAmount       int     //成交数量
// DealBalance      float64 //成交金额
// DealPrice        float64 //成交均价
// FuturesDirection string  //开平方向
func (admin *Admin) UpdateExcutionOrder(){
	for{
		for i, v := range entrusts{
			// entrusts key is order ref
			// excetions key is tacticid
			// case1 one tacticid -> one instrument

			// case2 one tacticid -> multi instruments
			
		}
	}
}