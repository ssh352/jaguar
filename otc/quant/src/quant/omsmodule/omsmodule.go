package main

import (
	log "github.com/thinkboy/log4go"
	"github.com/widuu/goini"
	emsbase "quant/emsmodule/base"
	"quant/helper"
)

var (
	conf          = goini.SetConfig(helper.QuantConfigFile)
	cachedEntrust = make(map[string]emsbase.EntrustPushResp, conf.GetInt(helper.ConfigOMSSessionName, helper.ConfigOMSEntrustLen))
)

func init() {
	log.LoadConfiguration(helper.QuantLogConfigFile)
}

func main() {

	p := newPuller()
	sr := newService()

	chs := make([]chan int, 2)
	for i := 0; i < 2; i++ {
		chs[i] = make(chan int)
	}

	go p.pullOrderResp(chs[0])
	go sr.run(chs[1])

	for _, ch := range chs {
		<-ch
	}
}
