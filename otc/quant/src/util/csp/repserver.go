package csp

import (
	zmq "github.com/pebbe/zmq3"
	log "github.com/thinkboy/log4go"
	// "github.com/vmihailenco/msgpack"
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
	rep     *zmq.Socket
	running bool
}

func (s *RepService) init() {
	s.running = false
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

// Stop quit run function
func (s *RepService) Stop() {
	s.running = false
}

func (s *RepService) run() {
	s.running = true
	for s.running {
		if msg, err := s.rep.RecvBytes(0); err != nil {
			log.Error("RepServer: %s", err.Error())
		} else {
			s.rep.SendBytes(s.Handle.HandleBReq(msg), 0)
		}
	}
	s.release()
}
