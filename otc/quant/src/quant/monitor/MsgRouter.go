package main

import (
	// "fmt"
	zmq "github.com/pebbe/zmq3"
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
	"quant/helper"
	"time"
	"util/csp"
)

func NewMsgRouter() *MsgRouter {
	mr := MsgRouter{}
	mr.init()
	go mr.run()
	return &mr
}

type MsgRouter struct {
	conf      *goini.Config
	repAddr   string
	pubAddr   string
	pullAddr  string
	routermap map[string]csp.IRequestMsg
	s         *csp.RepService
	pubmsg    chan csp.PubMsg
	pull      *zmq.Socket
}

func (m *MsgRouter) init() {
	log.Info("MsgRouter init")
	m.pubmsg = make(chan csp.PubMsg, 1000)
	m.conf = goini.SetConfig(helper.QuantConfigFile)

	// Create Rep Server
	m.repAddr = m.conf.GetStr("monitor", "rep_addr")
	m.s = csp.NewRepService(m.repAddr, m)
	m.s.SetUnPack(false)
	log.Info("MsgRouter bind %s", m.repAddr)

	// Create Publish Server
	m.pubAddr = m.conf.GetStr("monitor", "pub_addr")
	csp.NewPubService(m.pubmsg, m.pubAddr)
	log.Info("MsgRouter publish %s", m.pubAddr)

	// Create Pull Server
	m.pullAddr = m.conf.GetStr("monitor", "pull_addr")
	m.pull, _ = zmq.NewSocket(zmq.PULL)
	err := m.pull.Bind(m.pullAddr)
	if err != nil {
		log.Error("MsgRouter create pull to %s fail.", m.pullAddr)
	}
	log.Info("MsgRouter pull %s", m.pullAddr)

	m.setRouterMap()
}

func (m *MsgRouter) setRouterMap() {
	m.routermap = make(map[string]csp.IRequestMsg, 6)

	rmsServiceAddr := m.conf.GetStr("riskmodule", "rep_addr")
	log.Info("RMS Client connect to %s", rmsServiceAddr)
	m.routermap["RMS"] = csp.NewReqClient(rmsServiceAddr)

	pmsServiceAddr := m.conf.GetStr("pmsmodule", "rep_addr")
	log.Info("PMS Client connect to %s", pmsServiceAddr)
	m.routermap["PMS"] = csp.NewReqClient(pmsServiceAddr)
}

func (m *MsgRouter) HandleBReq(req []byte) (rep []byte) {
	log.Info("MsgRouter recv : %d, %s", time.Now().UnixNano()/1e6, string(req))
	if (5 + int(req[4]&0x0F)) > len(req) {
		var (
			q csp.Request
			p csp.Response
		)
		msgpack.Unmarshal(req, &q)
		csp.SetRepV(&q, &p)
		p.RET = -1
		p.MSG = "Monitor can't route request, because of failing analysis " + string(req)
		return
	}
	To := string(req[5 : 5+int(req[4]&0x0F)])
	if client, ok := m.routermap[To]; ok {
		rep = client.RequestB(req)
	} else {
		log.Error("MsgRouter can't router to %s module", To)
	}
	return
}

func (m *MsgRouter) HandleReq(req csp.Request) (rep csp.Response) {
	// do nothing
	return
}

func (m *MsgRouter) run() {
	for {
		data, _ := m.pull.RecvBytes(0)
		log.Info("MsgRouter recv push msg: %d, %s", time.Now().UnixNano()/1e6, string(data))
		var msg csp.PubMsg
		msgpack.Unmarshal(data, &msg)
		m.pubmsg <- msg
	}
}
