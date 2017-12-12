package algorithm

import (
	"fmt"
	log "github.com/thinkboy/log4go"
	emsbase "quant/emsmodule/base"
	"quant/helper"
	omsbase "quant/omsmodule/base"
	"strconv"
	// "time"
)

// Common algorithm trade.
type common struct {
	*omsbase.Client
	tradeingStrategy map[string][]string
}

func (c *common) init() {
	// TacticID->(thirdreff1,thirdreff2,...,thirdreffn)
	c.tradeingStrategy = make(map[string][]string, 100)
}

// Trade is called by algorithmadmin.go.
func (c *common) trade(p emsbase.Portfolio) error {
	var thirdreffs []string
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
			entrusts[thirdreff] = emsbase.StrategyEntrust{StrategyInfo: p.StrategyInfo, ProductInfo: p.ProductInfo}
			itrade.LimitEntrust(e, p.AccountID, p.CombiNo)
		} else {
			return fmt.Errorf("Common algorithm can't find \"%s\" trade adapter", p.AdapterName)
		}
	}
	c.tradeingStrategy[p.TacticID] = thirdreffs
	c.checkPortStatus(p.TacticID)
	return nil
}

func (c *common) checkPortStatus(TacticID string) {
	if reffs, ok := c.tradeingStrategy[TacticID]; ok {
		for _, thirdreff := range reffs {
			r := helper.Request{From: "EMS",
				To:  "OMS",
				Cmd: "GetEntrust"}
			r.Params = append(r.Params, thirdreff)
			e := c.GetEntrust(r)
			log.Info("%v", e)
			if se, ok := entrusts[thirdreff]; ok {
				se.Entrust = e
			}
		}
	} else {
		log.Error("EMS common algorithm get thirdreff fail.")
	}
}

func (c *common) append() {

}
