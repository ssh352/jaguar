package main

import (
	log "github.com/thinkboy/log4go"
	"quant/helper"
)

func init() {
	helper.InitLogFile("pms")
}

func main() {
	log.Info("PMS start ")
	newSmanage()

	var wc chan int
	<-wc
	log.Info("PMS Exit")
}
