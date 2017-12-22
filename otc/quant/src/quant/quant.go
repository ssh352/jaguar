package main

import (
	"fmt"
	log "github.com/thinkboy/log4go"
	"github.com/widuu/goini"
	"os"
	"quant/helper"
	"quant/pmsmodule/sadmin"
	"strings"
	"time"
)

func init() {
	logfiles := make(map[string]string)
	logfiles["ERROR"] = fmt.Sprintf("quant_err%s.log", time.Now().Format("2006-01-02"))
	logfiles["INFO"] = fmt.Sprintf("quant_info%s.log", time.Now().Format("2006-01-02"))
	log.SetLogFiles(logfiles)
	log.LoadConfiguration(helper.QuantLogConfigFile)
}

func main() {
	log.Info("Quant start ")
	Conf := goini.SetConfig(helper.QuantConfigFile)
	Stretegies := Conf.GetStr("quant", "strategies")
	StretegyArr := strings.Split(Stretegies, "|")
	WaitChan := make([]chan int, len(StretegyArr))
	time.Sleep(time.Second)
	for idx, sname := range StretegyArr {
		s, ok := sadmin.NewStrategy(sname)
		if !ok {
			log.Error("Can't start \"%s\" strategy", sname)
			os.Exit(-1)
		}
		s.Init()
		go s.Run(WaitChan[idx])
	}
	log.Info("Quant main thread wait for exit")
	for _, ch := range WaitChan {
		<-ch
	}
	log.Info("Quant Exit")
}
