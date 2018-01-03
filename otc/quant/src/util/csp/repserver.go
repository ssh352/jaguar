package csp

import (
	zmq "github.com/pebbe/zmq3"
	"github.com/vmihailenco/msgpack"
)

func NewRepService(url string, handle IHandleMsg) *RepService {
	s := RepService{URL: url, Handle: handle}
	s.init()
	return &s
}

// RepService is uesd for communication server
type RepService struct {
	URL     string
	Handle  IHandleMsg
	Unpack  bool
	rep     *zmq.Socket
	running bool
}

func (s *RepService) SetUnPack(unpack bool) {
	s.Unpack = unpack
}

func (s *RepService) init() {
	s.running = false
	s.Unpack = true
	s.bind()
	go s.run()
}

func (s *RepService) release() {
	s.rep.Close()
}

func (s *RepService) bind() {
	s.rep, _ = zmq.NewSocket(zmq.REP)
	s.rep.Bind(s.URL)
}

func (s *RepService) Stop() {
	s.running = false
}

func (s *RepService) run() {
	s.running = true
	for s.running {
		msg, _ := s.rep.RecvBytes(0)
		if s.Unpack {
			var req Request
			_ = msgpack.Unmarshal(msg, &req)
			msg, _ = msgpack.Marshal(s.Handle.HandleReq(req))
			s.rep.SendBytes(msg, 0)
		} else {
			s.rep.SendBytes(s.Handle.HandleBReq(msg), 0)
		}
	}
	s.release()
}
