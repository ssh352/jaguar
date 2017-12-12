package main

import (
	"github.com/Workiva/go-datastructures/queue"
	zmq "github.com/pebbe/zmq3"
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
	"os"
	"quant/emsmodule/adapter"
	"quant/emsmodule/algorithm"
	emsbase "quant/emsmodule/base"
	"quant/helper"
	rms "quant/rmsmodule"
	// "time"
)

func init() {
	log.LoadConfiguration(helper.QuantLogConfigFile)
}

type emsModule struct {
	pullAddr  string
	pull      *zmq.Socket
	portQueue *queue.RingBuffer
	conf      *goini.Config
	algorithm algorithm.Admin
	running   bool
}

// NewEMSModule create emsModule
func NewEMSModule() *emsModule {
	ems := emsModule{}
	ems.init()
	return &ems
}

func (ems *emsModule) init() {
	ems.conf = goini.SetConfig(helper.QuantConfigFile)
	ems.running = false
	ems.pullAddr = ems.conf.GetStr(helper.ConfigEMSSessionName, helper.ConfigEMSPullAddr)
	ems.portQueue = queue.NewRingBuffer(uint64(ems.conf.GetInt(helper.ConfigEMSSessionName, helper.ConfigEMSPortQueueLen)))
	ems.algorithm = algorithm.Admin{Portqueue: ems.portQueue}

	ems.algorithm.Init()
	go ems.algorithm.Run()

	ems.listenPort()
}

func (ems *emsModule) listenPort() {
	log.Info("EMS listen to %s", ems.pullAddr)
	var err error
	ems.pull, err = zmq.NewSocket(zmq.PULL)
	if err != nil {
		log.Error("EMS create zmq pull failed.")
		os.Exit(-1)
	}
	err = ems.pull.Bind(ems.pullAddr)
	if err != nil {
		log.Error("EMS create zmq pull bind failed. Addr: %s", ems.pullAddr)
		os.Exit(-1)
	}
}

func (ems *emsModule) release() {
	ems.pull.Close()
}

func (ems *emsModule) stop() {
	ems.running = false
}

func (ems *emsModule) run(wc chan int) {
	ems.running = true
	for ems.running {
		data, _ := ems.pull.Recv(0)
		port := emsbase.Portfolio{ProductInfo: &emsbase.ProductInfo{}, StrategyInfo: &emsbase.StrategyInfo{}}
		err := msgpack.Unmarshal([]byte(data), &port)
		if err != nil {
			log.Error("EMS msgpack unmarshal portfolio failed.")
		} else {
			if Isvaild, _ := rms.CheckPort(&port); Isvaild {
				log.Info("%v", port)
				// arithmetic trade
				ems.portQueue.Put(port)
				// send to OMS
				adapter.PushToOMS(emsbase.PushData{MsgType: helper.PortPushData, Port: port})
			}
		}
	}
	wc <- 0
}

func main() {
	log.Info("EMS Start")
	ems := NewEMSModule()
	wc := make(chan int)
	go ems.run(wc)
	<-wc
	ems.release()
	log.Info("EMS Exit")
}
