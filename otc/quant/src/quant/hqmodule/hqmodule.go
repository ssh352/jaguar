// package hqmodule
package main

import (
	// "hash/adler32"
	"net/http"
	// "net/http"
	// "database/sql"
	"flag"
	"fmt"
	"os"
	// "quant/hqmodule/hqbase"
	"github.com/Workiva/go-datastructures/queue"
	_ "github.com/go-sql-driver/mysql"
	zmq "github.com/pebbe/zmq3"
	log "github.com/thinkboy/log4go"
	"github.com/widuu/goini"
	_ "net/http/pprof"
	"quant/helper"
	"quant/hqmodule/marketdata/readSH"
	"quant/hqmodule/marketdata/vss"
	"runtime"
	// "runtime/pprof"
	// "strconv"
	"strings"
	"sync"
	"time"
	"util/redis"
)

var (
	redisChan chan []byte
	rb        *queue.RingBuffer
	redisRb   *queue.RingBuffer
	mysqlRb   *queue.RingBuffer
	redisPool *redis.ConnPool

	stockCodeMap map[string]string
	conf         *goini.Config
	publisher    *zmq.Socket
)

// func publishQuote(topic string, b []byte) {
// 	_, err := publisher.SendMessage(topic, b)
// 	if topic == "600000.SH" {
// 		log.Info("sendToZMQ: %d", time.Now().UnixNano()/1e6)
// 	}
// 	if err != nil {
// 		log.Error("publisher sendMessage: ", err)
// 	}
// }

// 将行情数据发送到zeroMQ
func sendToZMQ() {
	// defer publisher.Close()
	var topic string
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
			if topic == "600000.SH" {
				log.Info("sendToZMQ before send: %d", time.Now().UnixNano()/1e6)
			}
			_, err := publisher.SendMessage(topic, b)
			if topic == "600000.SH" {
				log.Info("sendToZMQ after send: %d", time.Now().UnixNano()/1e6)
			}

			if err != nil {
				log.Error("publisher sendMessage: ", err)
			}
		} else {
			time.Sleep(time.Duration(1) * time.Millisecond)
		}
	}
}

// func sendToMysql() {
// 	log.Info("-----------sendToMysql-----------")

// 	// Open database connection
// 	db, err := sql.Open("mysql", conf.GetValue("mysql", "mysqlusername")+":"+conf.GetValue("mysql", "mysqlpwd")+"@"+conf.GetValue("mysql", "mysqlurl"))
// 	if err != nil {
// 		log.Error("open database connection: ", err)
// 	}
// 	defer db.Close()

// 	tx, err := db.Begin()
// 	if err != nil {
// 		log.Error("db.Begin err: ", err)
// 	}
// 	replaceStmt, err := tx.Prepare("REPLACE INTO realtimemarketdata(Code,Time,Status,PreClose,Open,High,Low,Last,AskPrice1,AskPrice2,AskPrice3,AskPrice4,AskPrice5,AskVol1,AskVol2,AskVol3,AskVol4,AskVol5,BidPrice1,BidPrice2,BidPrice3,BidPrice4,BidPrice5,BidVol1,BidVol2,BidVol3,BidVol4,BidVol5,NumTrades,Volume,Turnover,TotalBidVol,TotalAskVol,WeightedAvgBidPrice,WeightedAvgAskPrice,IOPV,YieldToMaturity,HighLimited,LowLimited) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
// 	if err != nil {
// 		log.Error("db.Prepare err: ", err)
// 	}

// 	updateStmt, err := tx.Prepare("UPDATE realtimemarketdata SET Time=?,Status=?,High=?,Low=?,Last=?,AskPrice1=?,AskPrice2=?,AskPrice3=?,AskPrice4=?,AskPrice5=?,AskVol1=?,AskVol2=?,AskVol3=?,AskVol4=?,AskVol5=?,BidPrice1=?,BidPrice2=?,BidPrice3=?,BidPrice4=?,BidPrice5=?,BidVol1=?,BidVol2=?,BidVol3=?,BidVol4=?,BidVol5=?,NumTrades=?,Volume=?,Turnover=? WHERE Code=?")
// 	if err != nil {
// 		log.Error("db.Prepare err: ", err)
// 	}

// 	t0 := time.Now()
// 	i := 0
// 	isClear := false
// 	for {
// 		if mysqlRb.Len() > 0 {
// 			msg, _ := mysqlRb.Get()
// 			timeNow := time.Now().Format("2006-01-02 15:04:05")
// 			subTime := timeNow[11:13] + timeNow[14:16] + timeNow[17:19]
// 			subValue, _ := strconv.Atoi(subTime)
// 			if 93000 < subValue && subValue < 93500 && isClear == false {
// 				stockCodeMap = make(map[string]string)
// 				isClear = true
// 			} else if 93000 > subValue || subValue > 93500 {
// 				isClear = false
// 			}

// 			if mkd, ok := msg.(hqbase.Marketdata); ok {
// 				if _, ok = stockCodeMap[mkd.Code]; ok {
// 					_, err := updateStmt.Exec(mkd.Time, mkd.Status, mkd.High, mkd.Low, mkd.Match, mkd.AskPrice[0], mkd.AskPrice[1], mkd.AskPrice[2], mkd.AskPrice[3], mkd.AskPrice[4], mkd.AskVol[0], mkd.AskVol[1], mkd.AskVol[2], mkd.AskVol[3], mkd.AskVol[4], mkd.BidPrice[0], mkd.BidPrice[1], mkd.BidPrice[2], mkd.BidPrice[3], mkd.BidPrice[4], mkd.BidVol[0], mkd.BidVol[1], mkd.BidVol[2], mkd.BidVol[3], mkd.BidVol[4], mkd.NumTrades, mkd.Volume, mkd.Turnover, mkd.Code)
// 					if err != nil {
// 						log.Error("Exec error:", err)
// 					}
// 				} else {
// 					_, err := replaceStmt.Exec(mkd.Code, mkd.Time, mkd.Status, mkd.PreClose, mkd.Open, mkd.High, mkd.Low, mkd.Match, mkd.AskPrice[0], mkd.AskPrice[1], mkd.AskPrice[2], mkd.AskPrice[3], mkd.AskPrice[4], mkd.AskVol[0], mkd.AskVol[1], mkd.AskVol[2], mkd.AskVol[3], mkd.AskVol[4], mkd.BidPrice[0], mkd.BidPrice[1], mkd.BidPrice[2], mkd.BidPrice[3], mkd.BidPrice[4], mkd.BidVol[0], mkd.BidVol[1], mkd.BidVol[2], mkd.BidVol[3], mkd.BidVol[4], mkd.NumTrades, mkd.Volume, mkd.Turnover, mkd.TotalBidVol, mkd.TotalAskVol, mkd.WeightedAvgBidPrice, mkd.WeightedAvgAskPrice, mkd.IOPV, mkd.YieldToMaturity, mkd.HighLimited, mkd.LowLimited)
// 					if err != nil {
// 						log.Error("Exec error:", err)
// 					}
// 					stockCodeMap[mkd.Code] = mkd.Code
// 				}
// 			} else {
// 				log.Error("interface转换Marketdata失败：msg.(hqbase.Marketdata)")
// 			}

// 			i++
// 			if i%1000 == 0 {
// 				tx.Commit()
// 				log.Info("insert %d with %v", i, time.Now().Sub(t0))

// 				tx, err = db.Begin()
// 				if err != nil {
// 					log.Error("db.Begin err: ", err)
// 				}

// 				replaceStmt, err = tx.Prepare("REPLACE INTO realtimemarketdata(Code,Time,Status,PreClose,Open,High,Low,Last,AskPrice1,AskPrice2,AskPrice3,AskPrice4,AskPrice5,AskVol1,AskVol2,AskVol3,AskVol4,AskVol5,BidPrice1,BidPrice2,BidPrice3,BidPrice4,BidPrice5,BidVol1,BidVol2,BidVol3,BidVol4,BidVol5,NumTrades,Volume,Turnover,TotalBidVol,TotalAskVol,WeightedAvgBidPrice,WeightedAvgAskPrice,IOPV,YieldToMaturity,HighLimited,LowLimited) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
// 				if err != nil {
// 					log.Error("db.Prepare err: ", err)
// 				}

// 				updateStmt, err = tx.Prepare("UPDATE realtimemarketdata SET Time=?,Status=?,High=?,Low=?,Last=?,AskPrice1=?,AskPrice2=?,AskPrice3=?,AskPrice4=?,AskPrice5=?,AskVol1=?,AskVol2=?,AskVol3=?,AskVol4=?,AskVol5=?,BidPrice1=?,BidPrice2=?,BidPrice3=?,BidPrice4=?,BidPrice5=?,BidVol1=?,BidVol2=?,BidVol3=?,BidVol4=?,BidVol5=?,NumTrades=?,Volume=?,Turnover=? WHERE Code=?")
// 				if err != nil {
// 					log.Error("db.Prepare err: ", err)
// 				}
// 			}
// 		} else {
// 			time.Sleep(time.Duration(1) * time.Millisecond)
// 		}
// 	}
// }

// 将数据插入redis
func setToRedisForHashMap() {
	for {
		// msg, _ := redisRb.Get()
		// b := msg.([]byte)
		b := <-redisChan
		// b := *b1
		var code string
		if b[18] == 164 {
			code = string(b[9:18])
		} else if b[19] == 164 {
			code = string(b[9:19])
		} else if b[20] == 164 {
			code = string(b[9:20])
		}
		_, err := redisPool.SetHashMapKey(helper.RedisKey, code, string(b))
		if err != nil {
			log.Error("insert to HashMap error: %s", err)
		}
	}
}

func init() {
	rb = queue.NewRingBuffer(10000)
	redisRb = queue.NewRingBuffer(10000)
	redisChan = make(chan []byte, 10000)
	mysqlRb = queue.NewRingBuffer(10000)
	stockCodeMap = make(map[string]string)

	logfiles := make(map[string]string)
	logfiles["ERROR"] = fmt.Sprintf("hqmodule_err%s.log", time.Now().Format("2006-01-02"))
	logfiles["INFO"] = fmt.Sprintf("hqmodule_info%s.log", time.Now().Format("2006-01-02"))
	log.SetLogFiles(logfiles)
	log.LoadConfiguration(helper.QuantLogConfigFile)
	conf = goini.SetConfig(helper.QuantConfigFile)
	var REDIS = map[string]string{
		"host":         conf.GetValue("redis", "redishost"),
		"database":     conf.GetValue("redis", "database"),
		"password":     conf.GetValue("redis", "password"),
		"maxOpenConns": conf.GetValue("redis", "maxOpenConns"),
		"maxIdleConns": conf.GetValue("redis", "maxIdleConns"),
	}
	redisPool = redis.InitRedis(REDIS)
	_, err := redisPool.Do("PING")
	if err != nil {
		log.Error(err)
		panic(err)
	}

	publisher, _ = zmq.NewSocket(zmq.PUB)
	publisher.Bind(conf.GetValue("hqmodule", "publish_addr"))
	log.Info("HQmodule start publish: %s", conf.GetValue("hqmodule", "publish_addr"))
}

func startHqserver() {
	flag.Parse()
	var wg sync.WaitGroup
	wg.Add(1)

	// go sendToMysql()
	for i := 0; i < runtime.NumCPU(); i++ {
		go sendToZMQ()
	}

	for i := 0; i < runtime.NumCPU(); i++ {
		go setToRedisForHashMap()
	}

	vss.InitVss(conf)
	go vss.Heartbeat()
	go vss.ReceiveMarketData(redisChan, rb, mysqlRb)
	// go vss.ReceiveMarketData(redisRb, rb, mysqlRb)

	go func() {
		log.Info(http.ListenAndServe("localhost:6060", nil))
	}()

	time.Sleep(time.Second)
	quotes := conf.GetStr(helper.ConfigHQSessionName, "sh")
	for _, q := range strings.Split(quotes, "|") {
		if q == "md001" {
			// go readSH.Md001map(redisRb, rb, mysqlRb)
			go readSH.Md001map(redisChan, rb, mysqlRb)
		} else if q == "md002" {
			// go readSH.Md002map(redisRb, rb, mysqlRb)
			for i := 0; i < runtime.NumCPU(); i++ {
				go readSH.Md002map(redisChan, rb, mysqlRb)
			}
		} else if q == "md004" {
			// go readSH.Md004map(redisRb, rb, mysqlRb)
			go readSH.Md004map(redisChan, rb, mysqlRb)
		}
	}

	go readSH.Readfile(conf)

	// for i := 0; i < 30; i++ {
	for {
		showBuffLen()
		time.Sleep(time.Second * 1)
	}

	// pprof.StopCPUProfile()
	// os.Exit(0)
	wg.Wait()
	os.Exit(0)
}

func showBuffLen() {
	len1, len2, len3 := readSH.GetBuffLen()
	log.Info("redisChan: %d redisRb: %d, rb: %d , mysqlRb: %d, rbmd001map: %d, rbmd002map: %d, rbmd004map: %d, num of goroutine: %d",
		len(redisChan), redisRb.Len(), rb.Len(), mysqlRb.Len(), len1, len2, len3, runtime.NumGoroutine())
}

func main() {
	startHqserver()
}
