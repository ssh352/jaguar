package option

import (
	"fmt"
	log "github.com/thinkboy/log4go"
	"quant/emsmodule/base"
	"quant/hqmodule/base"
	"quant/pmsmodule/base"
	"strconv"
	"time"
)

// NewDeltaHedge generate 'deltaHedge' strategy. Which is used by sbase.IStrategy
//params[0]: strategyname
//params[1]: RiskAversionRation
//params[2]: TradeCost
//params[3]: NotionalPrincipal
//params[4]: SubQuoteCodes
//params[5]: AdapterName
//params[6]: AccountID
//params[7]: CombiNo
func NewDeltaHedge(params []string) (pmsbase.IStrategy, error) {
	if len(params) != 13 {
		return nil, fmt.Errorf("NewDeltaHedge need 13 params, but get %d", len(params))
	}

	RiskAversionRation, _ := strconv.ParseFloat(params[1], 64)
	TradeCost, _ := strconv.ParseFloat(params[2], 64)
	NotionalPrincipal, _ := strconv.ParseFloat(params[3], 64)

	IS := &deltaHedge{
		//params
		RiskAversionRation: RiskAversionRation,
		TradeCost:          TradeCost,
		NotionalPrincipal:  NotionalPrincipal,
		//sbase params
		Sbase: &pmsbase.Sbase{
			StrategyName:  params[0],
			SubQuoteCodes: params[4],
			AdapterName:   params[5],
			AccountID:     params[6],
			CombiNo:       params[7],
		},
	}
	IS.SInit()
	return IS, nil
}

type DeltaHedgeConfig struct {
	RiskAversionRation float64
	TradeCost          float64
	NotionalPrincipal  float64
	*pmsbase.SbaseConfig
}

type deltaHedge struct {
	*pmsbase.Sbase
	// params
	RiskAversionRation float64
	TradeCost          float64
	NotionalPrincipal  float64
}

func (dh *deltaHedge) SInit() {
	dh.OMSSubTopic = dh.AccountID
	dh.S = dh
	dh.Init()
}

func (dh *deltaHedge) CalcSignal(q *hqbase.Marketdata) int {
	log.Info("deltaHedge CalcSignal")
	ret := emsbase.DoNothing
	// if dh.getDelta(last) > dh.DelteUp {
	// 	ret = emsbase.OpenLong
	// } else if dh.getDelta(last) < dh.DelteDown {
	ret = emsbase.OpenShort
	// }
	return ret
}

func (dh *deltaHedge) GeneratePortfolio(bs int, mkdat *hqbase.Marketdata) {
	log.Info("deltaHedge GeneratePortfolio")
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

		dh.Port = emsbase.Portfolio{ProductInfo: &emsbase.ProductInfo{}, StrategyInfo: &emsbase.StrategyInfo{}}
		dh.Port.Algorithm = "COMMON"
		dh.Port.SecurityEntrusts = append(dh.Port.SecurityEntrusts, e)
	}
}

func (dh *deltaHedge) getDelta(last float64) float64 {
	return 0.75
}
