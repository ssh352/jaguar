package sbase

import (
	zmq "github.com/pebbe/zmq3"
	"github.com/satori/go.uuid"
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
	emsbase "quant/emsmodule/base"
	"quant/helper"
	"strings"
)

// IStrategy define the function which should be implemented by strategy
type IStrategy interface {
	Init()
	Run(chan int)
}

// Sbase provide basic function for strategy
type Sbase struct {
	StrategyName         string
	StockAccount         string
	FutureAccount        string
	AccountID            string
	ProdID               string
	CombiNo              string
	AdapterName          string
	SubQuoteCodes        string   //subscribe qoute id, example 000001.SZ|600000.SH
	SubQuoteCodesArr     []string //subscribe qoute id, example arr[0] = 000001.SZ, arr[1] = 600000.SH
	OMSSubTopic          string   //the topic which is used for subscribe oms publish order response.
	HqModulePublishAddr  string   //zmq pub/sub
	OMSModulePublishAddr string   //zmq pub/sub
	OMSModuleQueryAddr   string   //zmq req/rep
	EMSModuleAddr        string   //zmq pull/push

	HqModuleSub   *zmq.Socket
	OmsModuleSub  *zmq.Socket
	OmsModuleRep  *zmq.Socket
	EmsModulePush *zmq.Socket

	Port emsbase.Portfolio
}

// Init subscribe quote from hqmodule[sub]
// Connect to emsmodule[push]
// Subscribe trade response from omsmodule
// Connect to omsmodule service[sub]
func (sb *Sbase) Init() {
	sb.initConfig()
	sb.subQuote()
	sb.connectToEMS()
	sb.subOMSResp()
	sb.connectToOMS()
}

func (sb *Sbase) initConfig() {
	log.Info("Strategy %s init config. Loding ./conf/quant.ini", sb.StrategyName)
	conf := goini.SetConfig(helper.QuantConfigFile)
	sb.HqModulePublishAddr = conf.GetValue("hqmodule", "publish_addr")
	sb.OMSModulePublishAddr = conf.GetValue("omsmodule", "publish_addr")
	sb.OMSModuleQueryAddr = conf.GetValue("omsmodule", "req_addr")
	sb.EMSModuleAddr = conf.GetValue("emsmodule", "pull_addr")
}

func (sb *Sbase) subQuote() {
	log.Info("Strategy %s subscribe quote: %s", sb.StrategyName, sb.HqModulePublishAddr)
	sb.HqModuleSub, _ = zmq.NewSocket(zmq.SUB)
	sb.HqModuleSub.Connect(sb.HqModulePublishAddr)
	sb.setSubQuote()
}

func (sb *Sbase) setSubQuote() {
	sb.SubQuoteCodesArr = strings.Split(sb.SubQuoteCodes, "|")
	for _, filter := range sb.SubQuoteCodesArr {
		log.Info("Strategy %s subscribe: %s", sb.StrategyName, filter)
		sb.HqModuleSub.SetSubscribe(filter)
	}
}

func (sb *Sbase) connectToEMS() {
	log.Info("Strategy %s connect to ems: %s", sb.StrategyName, sb.EMSModuleAddr)
	sb.EmsModulePush, _ = zmq.NewSocket(zmq.PUSH)
	sb.EmsModulePush.Connect(sb.EMSModuleAddr)
}

func (sb *Sbase) subOMSResp() {
	log.Info("Strategy %s subscribe oms: %s", sb.StrategyName, sb.OMSModulePublishAddr)
	sb.OmsModuleSub, _ = zmq.NewSocket(zmq.SUB)
	sb.OmsModuleSub.Connect(sb.OMSModulePublishAddr)
	for _, t := range strings.Split(sb.OMSSubTopic, "|") {
		sb.OmsModuleSub.SetSubscribe(t)
	}
}

func (sb *Sbase) connectToOMS() {
	log.Info("Strategy %s connect to oms: %s", sb.StrategyName, sb.OMSModuleQueryAddr)
	sb.OmsModuleRep, _ = zmq.NewSocket(zmq.REQ)
	sb.OmsModuleRep.Connect(sb.OMSModuleQueryAddr)
}

// Release the connection with hq\ems\oms
func (sb *Sbase) Release() {
	defer sb.HqModuleSub.Close()
	defer sb.EmsModulePush.Close()
	defer sb.OmsModuleSub.Close()
	defer sb.OmsModuleRep.Close()
}

// Trade send portfolio to emsmodule
func (sb *Sbase) Trade() {
	sb.Port.StrategyName = sb.StrategyName
	sb.Port.ProdID = sb.ProdID
	sb.Port.CombiNo = sb.CombiNo
	sb.Port.AccountID = sb.AccountID
	sb.Port.AdapterName = sb.AdapterName
	sb.Port.TacticID = uuid.NewV4().String()
	bPort, err := msgpack.Marshal(sb.Port)
	if err == nil {
		log.Info("%v", sb.Port)
		sb.EmsModulePush.SendBytes(bPort, 1)
	}
}
