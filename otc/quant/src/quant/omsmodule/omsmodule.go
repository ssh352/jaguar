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
	helper.InitLogFile("oms")
}

func main() {
	log.Info("OMS start")
	p := newPuller()
	newService()

	ch := make(chan int)
	go p.pullOrderResp(ch)
	<-ch
	log.Info("OMS Exit")
}
