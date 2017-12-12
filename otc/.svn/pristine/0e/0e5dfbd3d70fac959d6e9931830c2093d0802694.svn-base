package datacenter

import (
	"github.com/Workiva/go-datastructures/queue"
	log "github.com/thinkboy/log4go"
	"github.com/widuu/goini"
	emsbase "quant/emsmodule/base"
	"quant/helper"
	"time"
	"util/db"
)

const (
	insertEntrustSQL string = "INSERT INTO jqorder(tactic_id,tactic_type,strategy_name, " +
		"prodid,account_code,combi_no,entrust_amount,entrust_direction,entrust_price,market_no, " +
		"price_type,stock_code,stockholder_id,insert_date,insert_time,third_reff,unix_time,capital_type,remark) " +
		"VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	updateEntrustSQL        string = "update jqorder set operator_no = ?, business_date = ?, business_time = ?, entrust_no = ?, entrust_price = ?, entrust_status = ?, invest_type= ?, price_type = ?, report_no = ?, report_seat = ?, stockholder_id = ? where third_reff = ?"
	updateEntrustByTradeSQL string = "update jqorder set entrust_status = ?, deal_amount = ?, deal_balance = ?, deal_price = ? where third_reff = ?"
	insertTradeSQL          string = "insert into jqtrade(account_code,batch_no,cancel_amount,combi_no,deal_amount,deal_balance,deal_date,deal_fee,deal_no,deal_price,deal_time,entrust_amount,entrust_direction,entrust_no,entrust_status,extsystem_id,futures_direction,market_no,operator_no,report_direction,report_seat,stock_code,stockholder_id,third_reff,total_deal_amount,total_deal_balance) " +
		"values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
)

type DBWorker struct {
	dbop            *db.MysqlWorker
	conf            *goini.Config
	sqls            chan string
	Portfolio       *queue.RingBuffer
	EntrustPush     *queue.RingBuffer
	TradePush       *queue.RingBuffer
	trades          *queue.RingBuffer
	pollTimeOut     int
	entrustBatchNum int
	tradeBatchNum   int
}

// NewDBWorker retrun *DBWorker
func NewDBWorker() *DBWorker {
	worker := DBWorker{}
	worker.init()
	return &worker
}

func (r *DBWorker) init() {
	r.conf = goini.SetConfig(helper.QuantConfigFile)
	r.EntrustPush = queue.NewRingBuffer(uint64(r.conf.GetInt(helper.ConfigOMSSessionName, helper.ConfigOMSEntrustLen)))
	r.TradePush = queue.NewRingBuffer(uint64(r.conf.GetInt(helper.ConfigOMSSessionName, helper.ConfigOMSTradeLen)))
	r.Portfolio = queue.NewRingBuffer(uint64(r.conf.GetInt(helper.ConfigEMSSessionName, helper.ConfigEMSPortQueueLen)))
	r.pollTimeOut = r.conf.GetInt(helper.ConfigOMSSessionName, helper.ConfigOMSPollTimeOut)
	r.entrustBatchNum = r.conf.GetInt(helper.ConfigOMSSessionName, helper.ConfigOMSEntrustUpdateBatchNum)
	r.tradeBatchNum = r.conf.GetInt(helper.ConfigOMSSessionName, helper.ConfigOMSTradeInsertBatchNum)
	config := db.MysqlConfig{
		MysqlUsernName: r.conf.GetStr(helper.ConfigMysqlSessionName, helper.ConfigMysqlUserName),
		MysqlPwd:       r.conf.GetStr(helper.ConfigMysqlSessionName, helper.ConfigMysqlPwd),
		MysqlURL:       r.conf.GetStr(helper.ConfigMysqlSessionName, helper.ConfigMysqlUrl),
	}
	r.sqls = make(chan string, r.conf.GetInt(helper.ConfigEMSSessionName, helper.ConfigEMSSqlLen))
	r.trades = queue.NewRingBuffer(uint64(r.conf.GetInt(helper.ConfigOMSSessionName, helper.ConfigOMSTradeLen)))
	r.dbop = &db.MysqlWorker{SQLs: r.sqls, MysqlConfig: &config}
	err := r.dbop.Init()
	if err != nil {
		log.Error("OMS connect to mysql fail. mysqlurl: %s. Error:%s.", r.dbop.MysqlURL, err.Error())
	} else {
		log.Info("OMS connect to %s mysql.", r.dbop.MysqlURL)
	}

	go r.insertEntrust()
	go r.updateEntrust()
	go r.updateEntrustByTrade()
	go r.insertTrade()
}

func (r *DBWorker) insertEntrust() {
	for {
		if r.Portfolio.Len() > 0 {
			tx, err := r.dbop.DB.Begin()
			if err != nil {
				log.Error("OMS DBWorker get context failed. Error: %s", err)
			} else {
				p1, _ := r.Portfolio.Poll(time.Microsecond * time.Duration(r.pollTimeOut))
				p := p1.(emsbase.Portfolio)
				for _, e := range p.SecurityEntrusts {
					_, err := tx.Exec(insertEntrustSQL, p.TacticID, p.TacticType, p.StrategyName, p.ProdID, p.AccountID, p.CombiNo,
						e.Vol, e.BS, e.Price, e.MarkerNo, "", e.TradeCode, p.AccountID, time.Now().Format("2006-01-02"), time.Now().Format("15:04:05"),
						e.ID, time.Now().UnixNano()/1e6, "STOCK", e.Remark)
					if err != nil {
						log.Error("OMS DBWorker exec fail. %s", err)
					}
				}
				err := tx.Commit()
				if err != nil {
					log.Error("OMS DBWorker commit fail. %s", err)
				}
			}
		} else {
			time.Sleep(time.Microsecond * time.Duration(r.pollTimeOut))
		}
	}
}

func (r *DBWorker) updateEntrustByTrade() {
	for {
		if r.trades.Len() > 0 {
			tx, err := r.dbop.DB.Begin()
			if err != nil {
				log.Error("OMS DBWorker get context failed. Error: %s", err)
			} else {
				for i := 0; i < r.tradeBatchNum && r.trades.Len() > 0; i++ {
					t1, _ := r.trades.Poll(time.Microsecond * time.Duration(r.pollTimeOut))
					t := t1.(emsbase.DealPushResp)
					_, err := tx.Exec(updateEntrustByTradeSQL, t.EntrustStatus, t.TotalDealAmount, t.TotalDealBalance, t.TotalDealBalance/float64(t.TotalDealAmount), t.ThirdReff)
					if err != nil {
						log.Error("OMS DBWorker exec fail. %s", err)
					}
				}
				err = tx.Commit()
				if err != nil {
					log.Error("OMS DBWorker commit fail. %s", err)
				}
			}
		} else {
			time.Sleep(time.Microsecond * time.Duration(r.pollTimeOut))
		}
	}
}

func (r *DBWorker) updateEntrust() {
	for {
		if r.EntrustPush.Len() > 0 {
			tx, err := r.dbop.DB.Begin()
			if err != nil {
				log.Error("OMS DBWorker get context failed. Error: %s", err)
			} else {
				for i := 0; i < r.entrustBatchNum && r.EntrustPush.Len() > 0; i++ {
					e1, _ := r.EntrustPush.Poll(time.Microsecond * time.Duration(r.pollTimeOut))
					e := e1.(emsbase.EntrustPushResp)
					_, err = tx.Exec(updateEntrustSQL, e.OperatorNo, e.BusinessDate, e.BusinessTime, e.EntrustNo, e.EntrustPrice, e.EntrustStatus, e.InvestType, e.PriceType, e.ReportNo, e.ReportSeat, e.StockholderID, e.ThirdReff)
					if err != nil {
						log.Error("OMS DBWorker exec fail. %s", err)
					}
				}
				err = tx.Commit()
				if err != nil {
					log.Error("OMS DBWorker commit fail. %s", err)
				}
			}
		} else {
			time.Sleep(time.Microsecond * time.Duration(r.pollTimeOut))
		}
	}
}

func (r *DBWorker) insertTrade() {
	for {
		if r.TradePush.Len() > 0 {
			tx, err := r.dbop.DB.Begin()
			if err != nil {
				log.Error("OMS DBWorker get context failed. Error: %s", err)
			} else {
				for i := 0; i < r.tradeBatchNum && r.TradePush.Len() > 0; i++ {
					t1, _ := r.TradePush.Poll(time.Microsecond * time.Duration(r.pollTimeOut))
					r.trades.Put(t1)
					t := t1.(emsbase.DealPushResp)
					_, err := tx.Exec(insertTradeSQL, t.AccountCode, t.BatchNo, t.CancelAmount, t.CombiNo,
						t.DealAmount, t.DealBalance, t.DealDate, t.DealFee,
						t.DealNo, t.DealPrice, t.DealTime, t.EntrustAmount,
						t.EntrustDirection, t.EntrustNo, t.EntrustStatus,
						t.ExtsystemID, t.FuturesDirection, t.MarketNo, t.OperatorNo,
						t.ReportDirection, t.ReportSeat, t.StockCode, t.StockholderID,
						t.ThirdReff, t.TotalDealAmount, t.TotalDealBalance)
					if err != nil {
						log.Error("OMS DBWorker exec fail. %s", err)
					}
				}
				err = tx.Commit()
				if err != nil {
					log.Error("OMS DBWorker commit fail. %s", err)
				}
			}
		} else {
			time.Sleep(time.Microsecond * time.Duration(r.pollTimeOut))
		}
	}
}
