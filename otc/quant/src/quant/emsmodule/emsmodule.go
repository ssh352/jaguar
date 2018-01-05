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
	"util/csp"
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
	rmsclient *csp.ReqClient
}

// newEMSModule create emsModule
func newEMSModule() *emsModule {
	ems := emsModule{}
	ems.init()
	return &ems
}

func (ems *emsModule) init() {
	ems.conf = goini.SetConfig(helper.QuantConfigFile)
	ems.running = false
	ems.rmsclient = csp.NewReqClient(ems.conf.GetStr("riskmodule", "rep_addr"))
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

func (ems *emsModule) checkPort(port *emsbase.Portfolio) bool {
	req := csp.Request{
		TO:   "RMS",
		FROM: "EMS",
		CMD:  "checkPort",
	}
	bport, _ := msgpack.Marshal(port)
	req.PARAMS = append(req.PARAMS, string(bport))
	breq, _ := msgpack.Marshal(req)
	brep := ems.rmsclient.RequestB(breq)

	var rep csp.Response
	msgpack.Unmarshal(brep, &rep)
	if rep.RET == 0 {
		return true
	}
	return false
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
			if Isvaild := ems.checkPort(&port); Isvaild {
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
	ems := newEMSModule()
	wc := make(chan int)
	go ems.run(wc)
	<-wc
	ems.release()
	log.Info("EMS Exit")
}
