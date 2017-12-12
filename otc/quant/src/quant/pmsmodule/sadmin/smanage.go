package sadmin

import (
	log "github.com/thinkboy/log4go"
	"quant/pmsmodule/option"
	"quant/pmsmodule/sbase"
)

var (
	strategyMap = make(map[string]func() sbase.IStrategy)
)

func init() {
	log.Info("Init pmsmodule/sadmin package, register \"DeltaHedge\" strategy")
	strategyMap["DeltaHedge"] = option.NewDeltaHedge
}

func NewStrategy(name string) (sbase.IStrategy, bool) {
	log.Info("New \"%s\" strategy", name)
	NewFunc, ok := strategyMap[name]
	if !ok {
		return nil, false
	}
	return NewFunc(), true
}
