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
}

func (s *service) init() {
	log.Info("RMS Service init")
	s.conf = goini.SetConfig(helper.QuantConfigFile)
	s.repAddr = s.conf.GetStr("riskmodule", "rep_addr")
	csp.NewRepService(s.repAddr, s)
	log.Info("RMS Service bind %s", s.repAddr)
}

func (s *service) HandleBReq(breq []byte) (brep []byte) {
	var req csp.Request
	msgpack.Unmarshal(breq, &req)
	if handle, ok := riskfuncrouter[req.CMD]; ok {
		rep := handle(&req)
		brep, _ = msgpack.Marshal(rep)
	} else {
		var rep csp.Response
		csp.SetRepV(&req, &rep)
		errmsg := "RMS Service can't router to '" + req.CMD + "' cmd"
		rep.MSG = errmsg
		rep.RET = -1
		brep, _ = msgpack.Marshal(rep)
		log.Error(errmsg)
	}
	return
}
