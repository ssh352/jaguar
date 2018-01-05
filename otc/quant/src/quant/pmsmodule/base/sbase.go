package pmsbase

import (
	zmq "github.com/pebbe/zmq3"
	"github.com/satori/go.uuid"
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
	"quant/emsmodule/base"
	"quant/helper"
	"quant/hqmodule/base"
	"strings"
	"time"
)

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
	running              bool

	HqModuleSub   *zmq.Socket
	OmsModuleSub  *zmq.Socket
	OmsModuleRep  *zmq.Socket
	EmsModulePush *zmq.Socket

	Port emsbase.Portfolio

	S IStrategy
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
	sb.running = false
}

func (sb *Sbase) initConfig() {
	log.Info("Strategy %s Loding config %s", sb.GetID(), helper.QuantConfigFile)
	conf := goini.SetConfig(helper.QuantConfigFile)
	sb.HqModulePublishAddr = conf.GetValue("hqmodule", "publish_addr")
	sb.OMSModulePublishAddr = conf.GetValue("omsmodule", "publish_addr")
	sb.OMSModuleQueryAddr = conf.GetValue("omsmodule", "req_addr")
	sb.EMSModuleAddr = conf.GetValue("emsmodule", "pull_addr")
}

func (sb *Sbase) subQuote() {
	log.Info("Strategy %s subscribe quote: %s", sb.GetID(), sb.HqModulePublishAddr)
	ctx, _ := zmq.NewContext()
	// ctx.SetTimeOut(500)	未生效！
	sb.HqModuleSub, _ = ctx.NewSocket(zmq.SUB)
	sb.HqModuleSub.Connect(sb.HqModulePublishAddr)
	sb.setSubQuote()
}

func (sb *Sbase) setSubQuote() {
	sb.SubQuoteCodesArr = strings.Split(sb.SubQuoteCodes, "|")
	for _, filter := range sb.SubQuoteCodesArr {
		log.Info("Strategy %s subscribe: %s", sb.GetID(), filter)
		sb.HqModuleSub.SetSubscribe(filter)
	}
}

func (sb *Sbase) connectToEMS() {
	log.Info("Strategy %s connect to ems: %s", sb.GetID(), sb.EMSModuleAddr)
	sb.EmsModulePush, _ = zmq.NewSocket(zmq.PUSH)
	sb.EmsModulePush.Connect(sb.EMSModuleAddr)
}

func (sb *Sbase) subOMSResp() {
	log.Info("Strategy %s subscribe oms: %s", sb.GetID(), sb.OMSModulePublishAddr)
	sb.OmsModuleSub, _ = zmq.NewSocket(zmq.SUB)
	sb.OmsModuleSub.Connect(sb.OMSModulePublishAddr)
	for _, t := range strings.Split(sb.OMSSubTopic, "|") {
		sb.OmsModuleSub.SetSubscribe(t)
	}
}

func (sb *Sbase) connectToOMS() {
	log.Info("Strategy %s connect to oms: %s", sb.GetID(), sb.OMSModuleQueryAddr)
	sb.OmsModuleRep, _ = zmq.NewSocket(zmq.REQ)
	sb.OmsModuleRep.Connect(sb.OMSModuleQueryAddr)
}

// Release the connection with hq\ems\oms
func (sb *Sbase) Release() {
	sb.HqModuleSub.Close()
	sb.EmsModulePush.Close()
	sb.OmsModuleSub.Close()
	sb.OmsModuleRep.Close()
}

// GetID retrun strategy unique id
func (sb *Sbase) GetID() string {
	return (sb.StrategyName + "_" + sb.AccountID + "_" + sb.SubQuoteCodes)
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

// Run function will running in coroutine
func (sb *Sbase) Run() {
	log.Info("%s start running. Recevice quote...", sb.GetID())
	sb.running = true
	for sb.running {
		// TODO replace to timeout instead of time.Sleep
		msg, err := sb.HqModuleSub.RecvMessage(1)
		if err != nil {
			log.Error("%s hqModuleSub recevice data error: %s", sb.GetID(), err.Error())
			break
		}
		if len(msg) == 1 {
			time.Sleep(time.Millisecond)
			continue
		}
		data := msg[1]

		var mkdat hqbase.Marketdata
		msgpack.Unmarshal([]byte(data), &mkdat)
		log.Info("timestamp:%d %+v", time.Now().UnixNano()/1e6, mkdat)
		if signal := sb.S.CalcSignal(&mkdat); signal != emsbase.DoNothing {
			sb.S.GeneratePortfolio(signal, &mkdat)
			sb.Trade()
			dat, err1 := sb.OmsModuleSub.RecvMessage(0)
			if err1 != nil {
				log.Error("%s omsModuleSub receive data error: %s", sb.GetID(), err1.Error())
			} else {
				log.Info("%v", dat)
			}
		}
	}
	log.Info("%s quit.", sb.GetID())
	sb.Release()
}

// Stop strategy
func (sb *Sbase) Stop() {
	sb.running = false
}

// Start strategy
func (sb *Sbase) Start() {
	go sb.Run()
}

func (sb *Sbase) GetAccount() string {
	return sb.AccountID
}

func (sb *Sbase) GetTradeStatus() string {
	return "-"
}

func (sb *Sbase) GetSecurityID() string {
	return sb.SubQuoteCodes
}
