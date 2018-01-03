package csp

import (
	zmq "github.com/pebbe/zmq3"
	"github.com/vmihailenco/msgpack"
)

// ReqClient send request to `reqAddr`
type ReqClient struct {
	reqAddr string
	req     *zmq.Socket
}

// NewReqClient return oms client
func NewReqClient(url string) *ReqClient {
	c := &ReqClient{reqAddr: url}
	c.init()
	return c
}

// Init connect to oms service
func (c *ReqClient) init() {
	c.req, _ = zmq.NewSocket(zmq.REQ)
	c.req.Connect(c.reqAddr)
}

// Request send req to `reqAddr` server
func (c *ReqClient) Request(rq Request) (rep Response) {
	brq, _ := msgpack.Marshal(rq)
	c.req.SendBytes(brq, 0)
	resp, _ := c.req.RecvBytes(0)
	msgpack.Unmarshal(resp, &rep)
	return
}

// RequestB send req([]byte) to `reqAddr` server
func (c *ReqClient) RequestB(rq []byte) (rep []byte) {
	c.req.SendBytes(rq, 0)
	rep, _ = c.req.RecvBytes(0)
	return
}
