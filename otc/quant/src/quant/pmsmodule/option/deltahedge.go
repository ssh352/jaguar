package option

import (
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
	emsbase "quant/emsmodule/base"
	"quant/hqmodule/hqbase"
	"quant/pmsmodule/sbase"
	"time"
)

type deltaHedge struct {
	base               sbase.Sbase
	RiskAversionRation float64
	RiskFreeReturn     float64
	DelteUp            float64
	DelteDown          float64
	HedgeInterval      float64
	StrikePrice        float64
	conf               *goini.Config
}

// NewDeltaHedge generate 'deltaHedge' strategy. Which is used by sbase.IStrategy
func NewDeltaHedge() sbase.IStrategy {
	var IS sbase.IStrategy = new(deltaHedge)
	return IS
}

func (dh *deltaHedge) Init() {
	dh.conf = goini.SetConfig("./conf/strategy/s_deltahedge.ini")

	dh.RiskAversionRation = dh.conf.GetFloat64("deltahedge", "risk_aversion_ratio")
	dh.RiskFreeReturn = dh.conf.GetFloat64("deltahedge", "risk_free_return")
	dh.DelteUp = dh.conf.GetFloat64("deltahedge", "delta_up")
	dh.DelteDown = dh.conf.GetFloat64("deltahedge", "delta_down")
	dh.HedgeInterval = dh.conf.GetFloat64("deltahedge", "hedge_interval")
	dh.StrikePrice = dh.conf.GetFloat64("deltahedge", "strike_price")

	dh.base.StrategyName = "DeltaHedge"
	dh.base.SubQuoteCodes = dh.conf.GetStr("deltahedge", "subcode")
	dh.base.AdapterName = dh.conf.GetStr("deltahedge", "adaptername")
	dh.base.AccountID = dh.conf.GetStr("deltahedge", "account_code")
	dh.base.CombiNo = dh.conf.GetStr("deltahedge", "comi_no")

	dh.base.OMSSubTopic = dh.base.AccountID
	dh.base.Init()
}

func (dh *deltaHedge) generatePortfolio(bs int, mkdat *hqbase.Marketdata) {
	if EType, ok := emsbase.EntrustTypeMap[bs]; ok {
		e := emsbase.Entrust{TradeCode: "600000",
			// Price:         mkdat.Match,
			StockCode:     "600000.SH",
			Vol:           10000,
			MarkerNo:      "1",
			OpenCloseFlag: EType.OpenCloseFlag,
			BS:            EType.BS,
			TimeStamp:     time.Now().UnixNano() / 1e6,
			ID:            time.Now().UnixNano()}

		dh.base.Port = emsbase.Portfolio{ProductInfo: &emsbase.ProductInfo{}, StrategyInfo: &emsbase.StrategyInfo{}}
		dh.base.Port.Algorithm = "COMMON"
		dh.base.Port.SecurityEntrusts = append(dh.base.Port.SecurityEntrusts, e)
	}
}

func (dh *deltaHedge) getDelta(last float64) float64 {
	return 0.75
}

func (dh *deltaHedge) calcSignal(last float64) int {
	ret := emsbase.DoNothing
	if dh.getDelta(last) > dh.DelteUp {
		ret = emsbase.OpenLong
	} else if dh.getDelta(last) < dh.DelteDown {
		ret = emsbase.OpenShort
	}
	return ret
}

func (dh *deltaHedge) Run(waitchan chan int) {
	log.Info("DeltaHedge start recevice quote...")
	for i := 0; true; i++ {
		msg, err := dh.base.HqModuleSub.RecvMessage(0)
		if err != nil {
			log.Error("DeltaHedge hqModuleSub recevice data error: %s", err.Error())
			break
		}
		data := msg[1]

		var mkdat hqbase.Marketdata
		err = msgpack.Unmarshal([]byte(data), &mkdat)
		if err != nil {
			log.Info("timestamp:%d %+v", time.Now().UnixNano()/1e6, err.Error())
		} else {
			log.Info("timestamp:%d %+v", time.Now().UnixNano()/1e6, mkdat)
			if signal := dh.calcSignal(mkdat.Match); signal != emsbase.DoNothing {
				dh.generatePortfolio(signal, &mkdat)
				dh.base.Trade()
				dat, err1 := dh.base.OmsModuleSub.RecvMessage(0)
				if err1 != nil {
					log.Error("DeltaHedge omsModuleSub receive data error: %s", err1.Error())
				} else {
					log.Info("%v", dat)
				}

			}
		}
	}

	dh.base.Release()
	log.Info("before write waitchan")
	waitchan <- 0
	log.Info("after write waitchan")
}
