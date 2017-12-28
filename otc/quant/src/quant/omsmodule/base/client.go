package omsbase

import (
	log "github.com/thinkboy/log4go"
	"github.com/widuu/goini"
	"quant/helper"
	"util/csp"
)

// Client used for get oms information
type Client struct {
	conf    *goini.Config
	reqAddr string
	reqC    *csp.ReqClient
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
	c.reqC = csp.NewReqClient(c.reqAddr)
	log.Info("OMS client connect to %s.", c.reqAddr)
}

// Request send req to server
func (c *Client) Request(rq csp.Request) (rep csp.Response) {
	return c.reqC.Request(rq)
}

// RequestB send req to server
func (c *Client) RequestB(rq []byte) (rep []byte) {
	return c.reqC.RequestB(rq)
}
