package csp

import (
	zmq "github.com/pebbe/zmq3"
	log "github.com/thinkboy/log4go"
)

// PubMsg publish message
type PubMsg struct {
	Topic string
	Msg   []byte
}

// NewPub create Publisher
func NewPubService(m chan PubMsg, url string) *Publisher {
	p := &Publisher{MsgsChan: m, URL: url}
	p.init()
	return p
}

// Publisher publisher server
type Publisher struct {
	MsgsChan  chan PubMsg
	URL       string
	publisher *zmq.Socket
}

func (p *Publisher) init() {
	p.publisher, _ = zmq.NewSocket(zmq.PUB)
	p.publisher.Bind(p.URL)
	go p.run()
}

func (p *Publisher) run() {
	for {
		mg := <-p.MsgsChan
		_, err := p.publisher.SendMessage(mg.Topic, mg.Msg)
		log.Info("Publisher Topic:%s, Msg:%s", mg.Topic, string(mg.Msg))
		if err != nil {
			log.Error("publisher sendMessage: ", err)
		}
	}
}
