package adapter

import (
	zmq "github.com/pebbe/zmq3"
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
	emsbase "quant/emsmodule/base"
	"quant/helper"
)

var (
	pushAddr string
	push     *zmq.Socket
	conf     *goini.Config
)

func init() {
	log.LoadConfiguration(helper.QuantLogConfigFile)
	conf = goini.SetConfig(helper.QuantConfigFile)
	pushAddr = conf.GetStr(helper.ConfigOMSSessionName, helper.ConfigOMSPullAddr)
	push, _ = zmq.NewSocket(zmq.PUSH)
	push.Connect(pushAddr)
}

// PushToOMS push data to oms
func PushToOMS(dat emsbase.PushData) {
	bDat, err := msgpack.Marshal(dat)
	if err != nil {
		log.Error("EMS push msgpack marshal fail: ", err)
	} else {
		push.SendBytes(bDat, 1)
	}
}
