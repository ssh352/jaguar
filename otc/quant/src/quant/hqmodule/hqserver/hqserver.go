// package hqmodule
package main

import (
	"database/sql"
	"flag"
	"os"
	"quant/hqmodule/hqbase"
	"quant/hqmodule/marketdata/readSH"
	// "quant/hqmodule/marketdata/vss"
	"strconv"
	"sync"
	"time"
	"util/redis"

	"github.com/Workiva/go-datastructures/queue"
	_ "github.com/go-sql-driver/mysql"
	zmq "github.com/pebbe/zmq3"
	log "github.com/thinkboy/log4go"
	"github.com/widuu/goini"
)

var (
	rb        *queue.RingBuffer
	redisRb   *queue.RingBuffer
	mysqlRb   *queue.RingBuffer
	RedisPool *redis.ConnPool

	stockCodeMap map[string]string
	conf         *goini.Config
)

// 将行情数据发送到zeroMQ
func sendToZMQ() {
	log.Info("-----sendToZMQ-----")
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.Bind(conf.GetValue("hqmodule", "publish_addr"))

	log.Info("------publisher start on:%v------\n", conf.GetValue("hqmodule", "publish_addr"))
	//Ensure subscriber connection has time to complete
	//time.Sleep(time.Second)

	var topic string
	//t0 := time.Now()
	//i := 0
	for {
		if rb.Len() > 0 {
			msg, _ := rb.Get()
			b := msg.([]byte)
			if b[18] == 164 {
				topic = string(b[9:18])
			} else if b[19] == 164 {
				topic = string(b[9:19])
			} else if b[20] == 164 {
				topic = string(b[9:20])
			}
			//log.Info("send len %d, topic %s, rb.len %d", len(b), topic, rb.Len())
			_, err := publisher.SendMessage(topic, b)
			if err != nil {
				log.Error("publisher sendMessage: ", err)
			}

			//i++
			//if i%1000 == 0 {
			//	log.Info("put %d with %v", i, time.Now().Sub(t0))
			//}
		} else {
			time.Sleep(time.Duration(1) * time.Millisecond)
		}

	}
}

func sendToMysql() {
	log.Info("-----------sendToMysql-----------")

	// Open database connection
	db, err := sql.Open("mysql", conf.GetValue("mysql", "mysqlusername")+":"+conf.GetValue("mysql", "mysqlpwd")+"@"+conf.GetValue("mysql", "mysqlurl"))
	if err != nil {
		log.Error("open database connection: ", err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Error("db.Begin err: ", err)
	}
	replaceStmt, err := tx.Prepare("REPLACE INTO realtimemarketdata(Code,Time,Status,PreClose,Open,High,Low,Last,AskPrice1,AskPrice2,AskPrice3,AskPrice4,AskPrice5,AskVol1,AskVol2,AskVol3,AskVol4,AskVol5,BidPrice1,BidPrice2,BidPrice3,BidPrice4,BidPrice5,BidVol1,BidVol2,BidVol3,BidVol4,BidVol5,NumTrades,Volume,Turnover,TotalBidVol,TotalAskVol,WeightedAvgBidPrice,WeightedAvgAskPrice,IOPV,YieldToMaturity,HighLimited,LowLimited) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Error("db.Prepare err: ", err)
	}

	updateStmt, err := tx.Prepare("UPDATE realtimemarketdata SET Time=?,Status=?,High=?,Low=?,Last=?,AskPrice1=?,AskPrice2=?,AskPrice3=?,AskPrice4=?,AskPrice5=?,AskVol1=?,AskVol2=?,AskVol3=?,AskVol4=?,AskVol5=?,BidPrice1=?,BidPrice2=?,BidPrice3=?,BidPrice4=?,BidPrice5=?,BidVol1=?,BidVol2=?,BidVol3=?,BidVol4=?,BidVol5=?,NumTrades=?,Volume=?,Turnover=? WHERE Code=?")
	if err != nil {
		log.Error("db.Prepare err: ", err)
	}

	t0 := time.Now()
	i := 0
	isClear := false
	for {
		if mysqlRb.Len() > 0 {
			msg, _ := mysqlRb.Get()
			timeNow := time.Now().Format("2006-01-02 15:04:05")
			subTime := timeNow[11:13] + timeNow[14:16] + timeNow[17:19]
			subValue, _ := strconv.Atoi(subTime)
			if 93000 < subValue && subValue < 93500 && isClear == false {
				stockCodeMap = make(map[string]string)
				isClear = true
			} else if 93000 > subValue || subValue > 93500 {
				isClear = false
			}

			if mkd, ok := msg.(hqbase.Marketdata); ok {
				if _, ok = stockCodeMap[mkd.Code]; ok {
					_, err := updateStmt.Exec(mkd.Time, mkd.Status, mkd.High, mkd.Low, mkd.Match, mkd.AskPrice[0], mkd.AskPrice[1], mkd.AskPrice[2], mkd.AskPrice[3], mkd.AskPrice[4], mkd.AskVol[0], mkd.AskVol[1], mkd.AskVol[2], mkd.AskVol[3], mkd.AskVol[4], mkd.BidPrice[0], mkd.BidPrice[1], mkd.BidPrice[2], mkd.BidPrice[3], mkd.BidPrice[4], mkd.BidVol[0], mkd.BidVol[1], mkd.BidVol[2], mkd.BidVol[3], mkd.BidVol[4], mkd.NumTrades, mkd.Volume, mkd.Turnover, mkd.Code)
					if err != nil {
						log.Error("Exec error:", err)
					}
				} else {
					_, err := replaceStmt.Exec(mkd.Code, mkd.Time, mkd.Status, mkd.PreClose, mkd.Open, mkd.High, mkd.Low, mkd.Match, mkd.AskPrice[0], mkd.AskPrice[1], mkd.AskPrice[2], mkd.AskPrice[3], mkd.AskPrice[4], mkd.AskVol[0], mkd.AskVol[1], mkd.AskVol[2], mkd.AskVol[3], mkd.AskVol[4], mkd.BidPrice[0], mkd.BidPrice[1], mkd.BidPrice[2], mkd.BidPrice[3], mkd.BidPrice[4], mkd.BidVol[0], mkd.BidVol[1], mkd.BidVol[2], mkd.BidVol[3], mkd.BidVol[4], mkd.NumTrades, mkd.Volume, mkd.Turnover, mkd.TotalBidVol, mkd.TotalAskVol, mkd.WeightedAvgBidPrice, mkd.WeightedAvgAskPrice, mkd.IOPV, mkd.YieldToMaturity, mkd.HighLimited, mkd.LowLimited)
					if err != nil {
						log.Error("Exec error:", err)
					}
					stockCodeMap[mkd.Code] = mkd.Code
				}
			} else {
				log.Error("interface转换Marketdata失败：msg.(hqbase.Marketdata)")
			}

			i++
			if i%1000 == 0 {
				tx.Commit()
				log.Info("insert %d with %v", i, time.Now().Sub(t0))

				tx, err = db.Begin()
				if err != nil {
					log.Error("db.Begin err: ", err)
				}

				replaceStmt, err = tx.Prepare("REPLACE INTO realtimemarketdata(Code,Time,Status,PreClose,Open,High,Low,Last,AskPrice1,AskPrice2,AskPrice3,AskPrice4,AskPrice5,AskVol1,AskVol2,AskVol3,AskVol4,AskVol5,BidPrice1,BidPrice2,BidPrice3,BidPrice4,BidPrice5,BidVol1,BidVol2,BidVol3,BidVol4,BidVol5,NumTrades,Volume,Turnover,TotalBidVol,TotalAskVol,WeightedAvgBidPrice,WeightedAvgAskPrice,IOPV,YieldToMaturity,HighLimited,LowLimited) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
				if err != nil {
					log.Error("db.Prepare err: ", err)
				}

				updateStmt, err = tx.Prepare("UPDATE realtimemarketdata SET Time=?,Status=?,High=?,Low=?,Last=?,AskPrice1=?,AskPrice2=?,AskPrice3=?,AskPrice4=?,AskPrice5=?,AskVol1=?,AskVol2=?,AskVol3=?,AskVol4=?,AskVol5=?,BidPrice1=?,BidPrice2=?,BidPrice3=?,BidPrice4=?,BidPrice5=?,BidVol1=?,BidVol2=?,BidVol3=?,BidVol4=?,BidVol5=?,NumTrades=?,Volume=?,Turnover=? WHERE Code=?")
				if err != nil {
					log.Error("db.Prepare err: ", err)
				}
			}
		} else {
			time.Sleep(time.Duration(1) * time.Millisecond)
		}

	}
}

// 将数据插入redis
func setToRedisForHashMap() {
	log.Info("-----------setToRedisForHashMap-----------")
	data := make(map[string]interface{})

	var code string
	for {
		if redisRb.Len() > 0 {
			msg, _ := redisRb.Get()
			//log.Info("redisRB.len: %d", redisRb.Len())
			b := msg.([]byte)
			if b[18] == 164 {
				code = string(b[9:18])
			} else if b[19] == 164 {
				code = string(b[9:19])
			} else if b[20] == 164 {
				code = string(b[9:20])
			}
			data[code] = string(b)
			if len(data) == 200 {
				_, err := RedisPool.SetHashMap("MarketMap_test", data)
				//	log.Info("data:%v\n", data)
				if err != nil {
					log.Error("insert to HashMap error: ", err)
				}
				//log.Info("插入redis数据库成功!插入记录：%v", md)
				for k, _ := range data {
					delete(data, k)
				}
			}
		} else {
			time.Sleep(time.Duration(1) * time.Millisecond)
		}

	}
}

func StartHqserver() {
	flag.Parse()
	var wg sync.WaitGroup
	wg.Add(1)

	log.LoadConfiguration("./conf/quant_log.xml")
	defer log.Close()
	conf = goini.SetConfig("./conf/quant.ini")

	var REDIS = map[string]string{
		"host":         conf.GetValue("redis", "redishost"),
		"database":     conf.GetValue("redis", "database"),
		"password":     conf.GetValue("redis", "password"),
		"maxOpenConns": conf.GetValue("redis", "maxOpenConns"),
		"maxIdleConns": conf.GetValue("redis", "maxIdleConns"),
	}
	RedisPool = redis.InitRedis(REDIS)
	_, err := RedisPool.Do("PING")
	if err != nil {
		log.Error("redis error")
		panic(err)
	}

	rb = queue.NewRingBuffer(9000)
	redisRb = queue.NewRingBuffer(9000)
	mysqlRb = queue.NewRingBuffer(9000)
	stockCodeMap = make(map[string]string)

	// go sendToMysql()

	go sendToZMQ()

	go setToRedisForHashMap()

	// vss.InitVss(conf)

	// go vss.Heartbeat()

	// go vss.ReceiveMarketData(redisRb, rb, mysqlRb)

	go readSH.Readfile(conf)

	go readSH.Md001map(redisRb, rb, mysqlRb)

	go readSH.Md002map(redisRb, rb, mysqlRb)

	go readSH.Md004map(redisRb, rb, mysqlRb)
	wg.Wait()
	os.Exit(0)
}

func main() {
	StartHqserver()
}
