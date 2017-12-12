package main

import (
	zmq "github.com/pebbe/zmq3"
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
	"os"
	emsbase "quant/emsmodule/base"
	"quant/helper"
	"quant/omsmodule/datacenter"
	// "time"
)

func newPuller() *puller {
	p := puller{}
	p.init()
	return &p
}

// puller handle data pushed from ems
type puller struct {
	pullAddr    string
	publishAddr string
	reqAddr     string
	pull        *zmq.Socket
	conf        *goini.Config
	running     bool
	worker      *datacenter.DBWorker
	cache       *datacenter.Cache
}

func (p *puller) init() {
	p.conf = goini.SetConfig(helper.QuantConfigFile)

	p.pullAddr = p.conf.GetStr(helper.ConfigOMSSessionName, helper.ConfigOMSPullAddr)
	p.publishAddr = p.conf.GetStr(helper.ConfigOMSSessionName, helper.ConfigOMSPublishAddr)
	p.reqAddr = p.conf.GetStr(helper.ConfigOMSSessionName, helper.ConfigOMSReqAddr)

	p.worker = datacenter.NewDBWorker()
	p.cache = datacenter.NewCache(cachedEntrust)

	p.listenOrder()
	p.running = false
}

func (p *puller) listenOrder() {
	log.Info("OMS listen to: %s", p.pullAddr)
	var err error
	p.pull, err = zmq.NewSocket(zmq.PULL)
	if err != nil {
		log.Error("OMS create pull fail.")
		os.Exit(-1)
	}
	err = p.pull.Bind(p.pullAddr)
	if err != nil {
		log.Error("OMS bind to %s fail.", p.publishAddr)
		os.Exit(-1)
	}
}

func (p *puller) stop() {
	p.running = false
}

func (p *puller) pullOrderResp(wc chan int) {
	p.running = true
	for p.running {
		data, _ := p.pull.Recv(0)
		pushdata := emsbase.PushData{Port: emsbase.Portfolio{ProductInfo: &emsbase.ProductInfo{}, StrategyInfo: &emsbase.StrategyInfo{}}}
		err := msgpack.Unmarshal([]byte(data), &pushdata)
		if err != nil {
			log.Error("OMS unmarshal push data fail: ", err)
		} else {
			if pushdata.MsgType == helper.EntrustRespPushData {
				p.worker.EntrustPush.Put(pushdata.Entrust)
				p.cache.EntrustPush.Put(pushdata.Entrust)
			} else if pushdata.MsgType == helper.TradePushData {
				p.worker.TradePush.Put(pushdata.Trade)
				p.cache.TradePush.Put(pushdata.Trade)
			} else if pushdata.MsgType == helper.PortPushData {
				p.worker.Portfolio.Put(pushdata.Port)
				p.cache.Portfolio.Put(pushdata.Port)
			}
		}
	}
	wc <- 1
}
