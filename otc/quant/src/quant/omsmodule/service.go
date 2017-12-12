package main

import (
	zmq "github.com/pebbe/zmq3"
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
	"quant/helper"
	"time"
)

func newService() *service {
	s := service{}
	s.init()
	return &s
}

type service struct {
	conf    *goini.Config
	repAddr string
	rep     *zmq.Socket
	running bool
}

func (s *service) init() {
	s.conf = goini.SetConfig(helper.QuantConfigFile)
	s.repAddr = s.conf.GetStr(helper.ConfigOMSSessionName, helper.ConfigOMSReqAddr)
	s.running = false
	s.rep, _ = zmq.NewSocket(zmq.REP)
	s.rep.Bind(s.repAddr)
}

func (s *service) release() {
	s.rep.Close()
}

func (s *service) stop() {
	s.running = false
}

func (s *service) run(wc chan int) {
	s.running = true
	for s.running {
		msg, _ := s.rep.RecvBytes(0)
		log.Info("oms recv : %s", time.Now().UnixNano()/1e6)
		var req helper.Request
		err := msgpack.Unmarshal(msg, &req)
		if err == nil {
			thirdreff := req.Params[0]
			if e, ok := cachedEntrust[thirdreff]; ok {
				msg, _ := msgpack.Marshal(e)
				log.Info(string(msg))
				s.rep.SendBytes(msg, 0)
			} else {
				s.rep.SendBytes([]byte(""), 0)
				log.Error("OMS service can't find '%s' entrust", thirdreff)
			}
		} else {
			log.Error("OMS service unmarshal request fail. Error: ", err)
		}
	}
	s.release()
	wc <- 1
}
