package algorithm

import (
	"fmt"
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
	emsbase "quant/emsmodule/base"
	"quant/helper"
	// omsbase "quant/omsmodule/base"
	"strconv"
	// "time"
	"util/csp"
)

// Common algorithm trade.
type common struct {
	omsclient        *csp.ReqClient
	tradeingStrategy map[string][]string
}

func (c *common) init() {
	// TacticID->(thirdreff1,thirdreff2,...,thirdreffn)
	c.tradeingStrategy = make(map[string][]string, 100)
	conf := goini.SetConfig(helper.QuantConfigFile)
	c.omsclient = csp.NewReqClient(conf.GetStr(helper.ConfigOMSSessionName, helper.ConfigOMSReqAddr))
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

func (c *common) checkPortStatus(TacticID string) {
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

func (c *common) append() {

}
