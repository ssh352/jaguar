package omsbase

import (
	zmq "github.com/pebbe/zmq3"
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
	emsbase "quant/emsmodule/base"
	"quant/helper"
)

type Client struct {
	conf    *goini.Config
	reqAddr string
	req     *zmq.Socket
}

// NewClient return oms client
func NewClient() *Client {
	c := &Client{}
	c.init()
	return c
}

// Init connect to oms service
func (c *Client) init() {
	c.conf = goini.SetConfig(helper.QuantConfigFile)
	c.reqAddr = c.conf.GetStr(helper.ConfigOMSSessionName, helper.ConfigOMSReqAddr)
	c.req, _ = zmq.NewSocket(zmq.REQ)
	log.Info("OMS client connect to %s.", c.reqAddr)
	c.req.Connect(c.reqAddr)
}

// GetEntrust retrun entrust from cached entrust
func (c *Client) GetEntrust(rq helper.Request) (entrust emsbase.EntrustPushResp) {
	brq, _ := msgpack.Marshal(rq)
	log.Info(string(brq))

	c.req.SendBytes(brq, 0)
	resp, _ := c.req.RecvBytes(0)
	msgpack.Unmarshal(resp, &entrust)
	return
}
