package algorithm

import (
	"fmt"
	"quant/emsmodule/base"
	"quant/helper"
	"strconv"
	"util/csp"

	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
)

// twap algorithm trade.
type twap struct {
	omsclient          *csp.ReqClient
	tradeingStrategy   map[string][]string
	conf               *goini.Config
	vollimit           int
	volratio           float64
	tradeunit          int
	waitfortrade       float64
	waitforappendtrade float64
	appendratio        float64
	appendnum          int
}

func (c *twap) init() {
	// TacticID->(thirdreff1,thirdreff2,...,thirdreffn)
	c.tradeingStrategy = make(map[string][]string, 100)
	c.conf = goini.SetConfig(helper.QuantConfigFile)
	c.omsclient = csp.NewReqClient(c.conf.GetStr(helper.ConfigOMSSessionName, helper.ConfigOMSReqAddr))

	c.vollimit = c.conf.GetInt("twap", "vollimit")
	c.volratio = c.conf.GetFloat64("twap", "volratio")
	c.tradeunit = c.conf.GetInt("twap", "tradeunit")
	c.waitfortrade = c.conf.GetFloat64("twap", "waitfortrade")
	c.waitforappendtrade = c.conf.GetFloat64("twap", "waitforappendtrade")
	c.appendratio = c.conf.GetFloat64("twap", "appendratio")
	c.appendnum = c.conf.GetInt("twap", "appendnum")
}

// Trade is called by algorithm/admin.go.
func (c *twap) trade(p emsbase.Portfolio) error {
	var thirdreffs []string

	num := len(p.FutureEntrusts) + len(p.SecurityEntrusts)
	if num > 1{

	}else{
		if len(p.FutureEntrusts) == 0 {
		excetions[p.TacticID] = emsbase.ExecutionOrder{
			StrategyName: p.StrategyName,
			TacticID   : p.TacticID,
			Algorithm        : p.Algorithm,
			OperatorNo       : p.OperatorNo,
			AccountCode : p.AccountID,      
			// StockCode        : p.
			
			BusinessTime     : 
			EntrustDirection :
			EntrustAmount    :        
			FuturesDirection :
		}
	}

	for _, e := range p.SecurityEntrusts {
		itrade, ok := adaptersMap[p.AdapterName]
		if ok {
			if e.Price == 0 {
				mkdat, err := helper.GetQuote(e.StockCode)
				if err == nil {
					if e.BS == emsbase.Buy {
						e.Price = mkdat.AskPrice[0]
					} else {
						e.Price = mkdat.BidPrice[0]
					}
				} else {
					return fmt.Errorf("Common algorithm can't find \"%s\" quote", e.StockCode)
				}
			}
			thirdreff := strconv.FormatInt(e.ID, 10)
			thirdreffs = append(thirdreffs, thirdreff)
			entrusts[thirdreff] = &emsbase.StrategyEntrust{StrategyInfo: p.StrategyInfo, ProductInfo: p.ProductInfo}
			itrade.LimitEntrust(e, p.AccountID, p.CombiNo)
		} else {
			return fmt.Errorf("Common algorithm can't find \"%s\" trade adapter", p.AdapterName)
		}
	}
	c.tradeingStrategy[p.TacticID] = thirdreffs
	c.checkPortStatus(p.TacticID)
	return nil
}

func (c *twap) checkPortStatus(TacticID string) {
	if reffs, ok := c.tradeingStrategy[TacticID]; ok {
		for _, thirdreff := range reffs {
			r := csp.Request{FROM: "EMS",
				TO:  "OMS",
				CMD: "GetEntrust"}
			r.PARAMS = append(r.PARAMS, thirdreff)
			req, _ := msgpack.Marshal(r)
			p := c.omsclient.RequestB(req)
			var rep csp.Response
			err := msgpack.Unmarshal(p, &rep)
			if err != nil {
				log.Error("Common algorithm unmarshal failed. %s", err)
				return
			}
			if se, ok := entrusts[thirdreff]; ok {
				msgpack.Unmarshal(rep.DAT, &se.Entrust)
			}
		}
	} else {
		log.Error("EMS common algorithm get thirdreff fail.")
	}
}

func (c *twap) append() {

}
