package main

import (
	log "github.com/thinkboy/log4go"
	"quant/helper"
)

func init() {
	helper.InitLogFile("rms")
}

func main() {
	log.Info("RMS start")
	admin := riskAdmin{}
	admin.init()

	var wc chan int
	<-wc
	log.Info("RMS Exit")
}
