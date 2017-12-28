package main

import (
	log "github.com/thinkboy/log4go"
	"github.com/widuu/goini"
	emsbase "quant/emsmodule/base"
	"quant/helper"
	"time"
)

var (
	conf          = goini.SetConfig(helper.QuantConfigFile)
	cachedEntrust = make(map[string]emsbase.EntrustPushResp, conf.GetInt(helper.ConfigOMSSessionName, helper.ConfigOMSEntrustLen))
)

func init() {
	logfiles := make(map[string]string)
	logfiles["ERROR"] = "omsmodule_err%s.log" + time.Now().Format("2006-01-02")
	logfiles["INFO"] = "omsmodule_info%s.log" + time.Now().Format("2006-01-02")
	log.SetLogFiles(logfiles)
	log.LoadConfiguration(helper.QuantLogConfigFile)
}

func main() {

	p := newPuller()
	newService()
	ch := make(chan int)
	go p.pullOrderResp(ch)
	<-ch
}
