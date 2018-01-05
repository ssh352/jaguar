package main

import (
	"fmt"
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
	"quant/helper"
	"util/csp"
)

var (
	omsfuncrouter = make(map[string]func(*csp.Request) csp.Response)
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
	log.Info("OMS Service init")
	s.conf = goini.SetConfig(helper.QuantConfigFile)

	// Create Rep Server
	s.repAddr = s.conf.GetStr(helper.ConfigOMSSessionName, helper.ConfigOMSReqAddr)
	csp.NewRepService(s.repAddr, s)
	s.setRouteMap()
	log.Info("OMS Service bind %s", s.repAddr)
}

func (s *service) setRouteMap() {
	omsfuncrouter["GetEntrust"] = s.GetEntrust
}

func (s *service) GetEntrust(req *csp.Request) (rep csp.Response) {
	csp.SetRepV(req, &rep)
	thirdreff := req.PARAMS[0]
	if e, ok := cachedEntrust[thirdreff]; ok {
		msg, _ := msgpack.Marshal(e)
		rep.DAT = msg
	} else {
		rep.RET = -1
		rep.MSG = fmt.Sprintf("OMS Service can't find  thirdreff:'%s' entrust", thirdreff)
	}
	return
}

func (s *service) HandleBReq(breq []byte) (brep []byte) {
	log.Info("OMS service receive %s ", string(breq))
	var req csp.Request
	msgpack.Unmarshal(breq, &req)
	if handle, ok := omsfuncrouter[req.CMD]; ok {
		rep := handle(&req)
		brep, _ = msgpack.Marshal(rep)
	} else {
		var rep csp.Response
		csp.SetRepV(&req, &rep)
		rep.MSG = "OMS service can't route to '" + req.CMD + "' cmd"
		rep.RET = -1
		brep, _ = msgpack.Marshal(rep)
		log.Error(rep.MSG)
	}
	return
}
