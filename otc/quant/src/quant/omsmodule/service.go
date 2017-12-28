package main

import (
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
	"quant/helper"
	"util/csp"
)

func newService() *service {
	s := service{}
	s.init()
	return &s
}

type service struct {
	conf    *goini.Config
	repAddr string
	ser     *csp.ServiceRep
}

func (s *service) init() {
	log.Info("Oms Service init")
	s.conf = goini.SetConfig(helper.QuantConfigFile)
	s.repAddr = s.conf.GetStr(helper.ConfigOMSSessionName, helper.ConfigOMSReqAddr)
	s.ser = csp.NewRepService(s.repAddr, s)
	log.Info("Oms Service bind %s", s.repAddr)
}

func (s *service) HandleReq(req csp.Request) csp.Response {
	rep := csp.Response{}
	thirdreff := req.Params[0]
	if e, ok := cachedEntrust[thirdreff]; ok {
		msg, _ := msgpack.Marshal(e)
		rep.Dat = append(rep.Dat, msg)
	}
	return rep
}
