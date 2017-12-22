package readSH

import (
	// "fmt"
	"hash/adler32"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/Workiva/go-datastructures/queue"
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
	"quant/helper"
	"strings"
	// "runtime"
	// "os"
	// "runtime/pprof"
)

var (
	rbmd001map   *queue.RingBuffer
	rbmd002map   *queue.RingBuffer
	rbmd004map   *queue.RingBuffer
	hashValueMap map[string]uint32
)

func init() {
	// 初始化ringbuffer
	rbmd001map = queue.NewRingBuffer(9000)
	rbmd002map = queue.NewRingBuffer(9000)
	rbmd004map = queue.NewRingBuffer(9000)
	hashValueMap = make(map[string]uint32, 2000)
}

// GetBuffLen return ringbuffer length
func GetBuffLen() (uint64, uint64, uint64) {
	return rbmd001map.Len(), rbmd002map.Len(), rbmd004map.Len()
}

func updateHs(fd *[]byte, i *int, idx int) bool {
	//判断和上次行情的哈希值是否相等
	code := string((*fd)[*i+7 : *i+13])
	value := adler32.Checksum((*fd)[*i+7 : *i+idx])
	if oldValue, ok := hashValueMap[code]; ok && value == oldValue {
		return false
	}
	hashValueMap[code] = value
	return true
}

// Readfile read Shanghai exchange quote file
func Readfile(conf *goini.Config) {
	filename := conf.GetValue("hqmodule", "sh_filename")
	interval, _ := strconv.Atoi(conf.GetValue("hqmodule", "sh_readfileinterval"))
	quotes := conf.GetStr(helper.ConfigHQSessionName, "sh")
	quotemap := make(map[string]bool, 3)
	for _, q := range strings.Split(quotes, "|") {
		quotemap[q] = true
	}
	_, md001ok := quotemap["md001"]
	_, md002ok := quotemap["md002"]
	_, md004ok := quotemap["md004"]

	var fd []byte
	var l int
	i := 0
	pauseinter := time.Duration(interval) * time.Millisecond
	for {
		fd, _ = ioutil.ReadFile(filename)
		l = len(fd) - 11
		for i = 0; i < l; i++ {
			if fd[i] == 0x0A {
				if fd[i+5] == 0x33 {
					i += 399
					continue
				} else if fd[i+5] == 0x31 {
					if md001ok {
						if updateHs(&fd, &i, 150) {
							rbmd001map.Put(fd[i+7 : i+150])
							i += 149
						}
					}
				} else if fd[i+5] == 0x32 {
					if md002ok {
						if updateHs(&fd, &i, 400) {
							rbmd002map.Put(fd[i+7 : i+400])
							if string(fd[i+7:i+7+6]) == "600000" {
								log.Info("ReadFile: %d", time.Now().UnixNano()/1e6)
							}
							i += 399
						}
					}
				} else if fd[i+5] == 0x34 {
					if md004ok {
						if updateHs(&fd, &i, 424) {
							rbmd004map.Put(fd[i+7 : i+424])
							i += 423
						}
					}
				}
			}
		}
		time.Sleep(pauseinter)
	}
}

// Md001map deal with Md001 data(index quote)
// func Md001map(redisRb *queue.RingBuffer, rb *queue.RingBuffer, mysqlRb *queue.RingBuffer) {
func Md001map(redisChan chan []byte, rb *queue.RingBuffer, mysqlRb *queue.RingBuffer) {
	for {
		if rbmd001map.Len() > 0 {
			msg, _ := rbmd001map.Get()
			marketData := getIndexQuote(msg)
			// _, _ = msgpack.Marshal(&marketData)
			data, _ := msgpack.Marshal(&marketData)
			redisChan <- data
			// redisRb.Put(data)
			rb.Put(data)
		} else {
			time.Sleep(time.Duration(1) * time.Millisecond)
		}
	}
}

// Md002map deal with Md002 data(stock quote)
// func Md002map(redisRb *queue.RingBuffer, rb *queue.RingBuffer, mysqlRb *queue.RingBuffer) {
func Md002map(redisChan chan []byte, rb *queue.RingBuffer, mysqlRb *queue.RingBuffer) {
	for {
		if rbmd002map.Len() > 0 {
			msg, _ := rbmd002map.Get()
			b := msg.([]byte)
			if string(b[0:6]) == "600000" {
				log.Info("Md002map get : %d", time.Now().UnixNano()/1e6)
			}
			marketData := getMkdata(msg, "STOCK")
			if marketData.Code == "600000.SH" {
				log.Info("Md002map getMkdata: %d", time.Now().UnixNano()/1e6)
			}
			// _, _ = msgpack.Marshal(&marketData)
			data, _ := msgpack.Marshal(&marketData)
			if marketData.Code == "600000.SH" {
				log.Info("Md002map Marshal: %d", time.Now().UnixNano()/1e6)
			}
			redisChan <- data
			// redisRb.Put(data)
			// publishQuote(marketData.Code, data)
			rb.Put(data)
			if marketData.Code == "600000.SH" {
				log.Info("Md002map after rb: %d", time.Now().UnixNano()/1e6)
			}
		} else {
			time.Sleep(time.Duration(1) * time.Millisecond)
		}
	}
}

// Md004map deal with Md004 data(fund quote)
// func Md004map(redisRb *queue.RingBuffer, rb *queue.RingBuffer, mysqlRb *queue.RingBuffer) {
func Md004map(redisChan chan []byte, rb *queue.RingBuffer, mysqlRb *queue.RingBuffer) {
	for {
		if rbmd004map.Len() > 0 {
			msg, _ := rbmd004map.Get()
			marketData := getMkdata(msg, "FUND")
			data, _ := msgpack.Marshal(&marketData)
			redisChan <- data
			// redisRb.Put(data)
			rb.Put(data)
		} else {
			time.Sleep(time.Duration(1) * time.Millisecond)
		}
	}
}
