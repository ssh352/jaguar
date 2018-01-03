package main

import (
	"fmt"
	log "github.com/thinkboy/log4go"
	"quant/helper"
	"time"
)

func init() {
	logfiles := make(map[string]string)
	logfiles["ERROR"] = fmt.Sprintf("monitor_err%s.log", time.Now().Format("2006-01-02"))
	logfiles["INFO"] = fmt.Sprintf("monitor_info%s.log", time.Now().Format("2006-01-02"))
	log.SetLogFiles(logfiles)
	log.LoadConfiguration(helper.QuantLogConfigFile)
}

func main() {
	NewMsgRouter()
	wc := make(chan int)
	<-wc
}
