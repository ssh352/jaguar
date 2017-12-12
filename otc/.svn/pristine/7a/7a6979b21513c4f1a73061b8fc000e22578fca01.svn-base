package datacenter

import (
	"github.com/Workiva/go-datastructures/queue"
	// log "github.com/thinkboy/log4go"
	"github.com/widuu/goini"
	emsbase "quant/emsmodule/base"
	"quant/helper"
	"strconv"
	"time"
)

// Cache entrust in mem
type Cache struct {
	conf          *goini.Config
	Portfolio     *queue.RingBuffer
	EntrustPush   *queue.RingBuffer
	TradePush     *queue.RingBuffer
	trades        *queue.RingBuffer
	CachedEntrust map[string]emsbase.EntrustPushResp
	pollTimeOut   int
}

// NewCache return *Cache
func NewCache(m map[string]emsbase.EntrustPushResp) *Cache {
	c := Cache{CachedEntrust: m}
	c.init()
	return &c
}

func (r *Cache) init() {
	r.conf = goini.SetConfig(helper.QuantConfigFile)
	r.EntrustPush = queue.NewRingBuffer(uint64(r.conf.GetInt(helper.ConfigOMSSessionName, helper.ConfigOMSEntrustLen)))
	r.TradePush = queue.NewRingBuffer(uint64(r.conf.GetInt(helper.ConfigOMSSessionName, helper.ConfigOMSTradeLen)))
	r.Portfolio = queue.NewRingBuffer(uint64(r.conf.GetInt(helper.ConfigEMSSessionName, helper.ConfigEMSPortQueueLen)))
	r.pollTimeOut = r.conf.GetInt(helper.ConfigOMSSessionName, helper.ConfigOMSPollTimeOut)
	r.trades = queue.NewRingBuffer(uint64(r.conf.GetInt(helper.ConfigOMSSessionName, helper.ConfigOMSTradeLen)))

	go r.insertEntrust()
	go r.updateEntrust()
	go r.updateEntrustByTrade()
}

func (r *Cache) insertEntrust() {
	for {
		if r.Portfolio.Len() > 0 {
			p1, _ := r.Portfolio.Poll(time.Microsecond * time.Duration(r.pollTimeOut))
			p := p1.(emsbase.Portfolio)
			for _, e := range p.SecurityEntrusts {
				r.cacheEntrust(p, e)
			}
		} else {
			time.Sleep(time.Microsecond * time.Duration(r.pollTimeOut))
		}
	}
}

func (r *Cache) updateEntrustByTrade() {
	for {
		if r.trades.Len() > 0 {
			t1, _ := r.trades.Get()
			t := t1.(emsbase.DealPushResp)
			r.updateCachedEntrustByTrade(t)
		} else {
			time.Sleep(time.Microsecond * time.Duration(r.pollTimeOut))
		}
	}
}

func (r *Cache) updateEntrust() {
	for {
		if r.EntrustPush.Len() > 0 {
			e1, _ := r.EntrustPush.Poll(time.Microsecond * time.Duration(r.pollTimeOut))
			e := e1.(emsbase.EntrustPushResp)
			r.updateCachedEntrust(e)
		} else {
			time.Sleep(time.Microsecond * time.Duration(r.pollTimeOut))
		}
	}
}

func (r *Cache) cacheEntrust(p emsbase.Portfolio, e emsbase.Entrust) {
	entrust := emsbase.EntrustPushResp{}
	entrust.CombiNo = p.CombiNo
	entrust.EntrustAmount = e.Vol
	entrust.EntrustDirection = strconv.Itoa(e.BS)
	entrust.EntrustPrice = e.Price
	entrust.MarketNo = e.MarkerNo
	entrust.StockCode = e.TradeCode
	r.CachedEntrust[strconv.FormatInt(e.ID, 10)] = entrust
}

func (r *Cache) updateCachedEntrust(e emsbase.EntrustPushResp) {
	r.CachedEntrust[e.ThirdReff] = e
}

func (r *Cache) updateCachedEntrustByTrade(t emsbase.DealPushResp) {
	if e, ok := r.CachedEntrust[t.ThirdReff]; ok {
		e.EntrustStatus = t.EntrustStatus
		e.DealAmount = t.TotalDealAmount
		e.DealBalance = t.TotalDealBalance
		e.DealPrice = t.TotalDealBalance / float64(t.TotalDealAmount)
	}
}
