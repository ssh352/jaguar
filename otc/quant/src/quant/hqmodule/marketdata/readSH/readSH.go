package readSH

import (
	"hash/adler32"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/Workiva/go-datastructures/queue"
	// log "github.com/thinkboy/log4go"
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
	hashValueMap = make(map[string]uint32)
}

// GetBuffLen return ringbuffer length
func GetBuffLen() (uint64, uint64, uint64) {
	return rbmd001map.Len(), rbmd002map.Len(), rbmd004map.Len()
}

func updateHs(fd []byte, i int, idx int) bool {
	//判断和上次行情的哈希值是否相等
	code := string(fd[i+7 : i+13])
	oldValue, ok := hashValueMap[code]
	if ok {
		value := adler32.Checksum(fd[i+7 : i+idx])
		if value == oldValue {
			return false
		}
		hashValueMap[code] = value
	} else {
		hashValueMap[code] = adler32.Checksum(fd[i+7 : i+idx])
	}
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

	for {
		fd, _ := ioutil.ReadFile(filename)
		l := len(fd) - 11
		for i := 0; i < l; i++ {
			if fd[i] == 0x0A {
				if fd[i+5] == 0x33 {
					i += 399
					continue
				} else if fd[i+5] == 0x31 {
					if _, ok := quotemap["md001"]; ok {
						if updated := updateHs(fd, 150, i); updated {
							rbmd001map.Put(fd[i+7 : i+150])
							i += 149
						}
					}
				} else if fd[i+5] == 0x32 {
					if _, ok := quotemap["md002"]; ok {
						if updated := updateHs(fd, 400, i); updated {
							rbmd002map.Put(fd[i+7 : i+400])
							i += 399
						}
					}
				} else if fd[i+5] == 0x34 {
					if _, ok := quotemap["md004"]; ok {
						if updated := updateHs(fd, 424, i); updated {
							rbmd004map.Put(fd[i+7 : i+424])
							i += 423
						}
					}
				}
				time.Sleep(1 * time.Millisecond)
			}
		}
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

// Md001map deal with Md001 data(index quote)
func Md001map(redisRb *queue.RingBuffer, rb *queue.RingBuffer, mysqlRb *queue.RingBuffer) {
	for {
		if rbmd001map.Len() > 0 {
			msg, _ := rbmd001map.Get()
			marketData := getIndexQuote(msg)
			data, err := msgpack.Marshal(&marketData)
			if err != nil {
				panic(err)
			}
			redisRb.Put(data)
			// rb.Put(data)
		} else {
			time.Sleep(time.Duration(1) * time.Microsecond)
		}
	}
}

// Md002map deal with Md002 data(stock quote)
func Md002map(redisRb *queue.RingBuffer, rb *queue.RingBuffer, mysqlRb *queue.RingBuffer) {
	for {
		if rbmd002map.Len() > 0 {
			msg, _ := rbmd002map.Get()
			marketData := getMkdata(msg, "STOCK")
			data, err := msgpack.Marshal(&marketData)
			if err != nil {
				panic(err)
			}
			redisRb.Put(data)
			// rb.Put(data)
		} else {
			time.Sleep(time.Duration(1) * time.Microsecond)
		}
	}
}

// Md004map deal with Md004 data(fund quote)
func Md004map(redisRb *queue.RingBuffer, rb *queue.RingBuffer, mysqlRb *queue.RingBuffer) {
	for {
		if rbmd004map.Len() > 0 {
			msg, _ := rbmd004map.Get()
			marketData := getMkdata(msg, "FUND")
			data, err := msgpack.Marshal(&marketData)
			if err != nil {
				panic(err)
			}
			redisRb.Put(data)
			// rb.Put(data)
		} else {
			time.Sleep(time.Duration(1) * time.Microsecond)
		}
	}
}
