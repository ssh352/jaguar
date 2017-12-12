package readSH

import (
	"fmt"
	"io/ioutil"
	"time"
	//"hash/adler32"
	"quant/hqmodule/hqbase"
	"strconv"

	"github.com/Workiva/go-datastructures/queue"
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
	"runtime"
)

var (
	rbmd001map *queue.RingBuffer
	rbmd002map *queue.RingBuffer
	rbmd004map *queue.RingBuffer

	hashValueMap map[string]uint32
)

func init() {
	// 初始化ringbuffer
	rbmd001map = queue.NewRingBuffer(2000)
	rbmd002map = queue.NewRingBuffer(2000)
	rbmd004map = queue.NewRingBuffer(2000)
	hashValueMap = make(map[string]uint32)
}

func Readfile(conf *goini.Config) {

	sh_filename := conf.GetValue("hqmodule", "sh_filename")
	interval, _ := strconv.Atoi(conf.GetValue("hqmodule", "sh_readfileinterval"))
	for {
		fd, _ := ioutil.ReadFile(sh_filename)
		l := len(fd) - 11
		md001 := 0
		md002 := 0
		md004 := 0
		for i := 0; i < l; i++ {
			if fd[i] == 0x0A {
				if fd[i+5] == 0x33 {
					i += 399
					continue
				} else if fd[i+5] == 0x31 {
					//判断和上次行情的哈希值是否相等
					//					code := string(fd[i+7:i+13])

					//					oldValue, ok := hashValueMap[code]
					//					if(ok){
					//						value := adler32.Checksum(fd[i+7 : i+150])
					//						if value == oldValue {
					//							continue
					//						}
					//						hashValueMap[code] = value
					//					}else{
					//						value := adler32.Checksum(fd[i+7 : i+150])
					//						hashValueMap[code] = value
					//					}

					md001++
					rbmd001map.Put(fd[i+7 : i+150])
					i += 149
				} else if fd[i+5] == 0x32 {
					//判断和上次行情的哈希值是否相等
					//					code := string(fd[i+7:i+13])

					//					oldValue, ok := hashValueMap[code]
					//					if(ok){
					//						value := adler32.Checksum(fd[i+7 : i+400])
					//						if value == oldValue {
					//							continue
					//						}
					//						hashValueMap[code] = value
					//					}else{
					//						value := adler32.Checksum(fd[i+7 : i+400])
					//						hashValueMap[code] = value
					//					}

					md002++
					rbmd002map.Put(fd[i+7 : i+400])
					i += 399
				} else if fd[i+5] == 0x34 {
					//判断和上次行情的哈希值是否相等
					//					code := string(fd[i+7:i+13])

					//					oldValue, ok := hashValueMap[code]
					//					if(ok){
					//						value := adler32.Checksum(fd[i+7 : i+424])
					//						if value == oldValue {
					//							continue
					//						}
					//						hashValueMap[code] = value
					//					}else{
					//						value := adler32.Checksum(fd[i+7 : i+424])
					//						hashValueMap[code] = value
					//					}

					md004++
					rbmd004map.Put(fd[i+7 : i+424])
					i += 423
				} else {
					//fmt.Println("error")
				}

			}

		}

		//log.Info("need to deal num:md001=%d,md002=%d,md004=%d", md001, md002, md004)

		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func Md001map(redisRb *queue.RingBuffer, rb *queue.RingBuffer, mysqlRb *queue.RingBuffer) {
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(seq int) {
			count := 0
			fmt.Println("start ", seq)
			tt1 := time.Now()
			for {
				if rbmd001map.Len() > 0 {
					count++
					msg, _ := rbmd001map.Get()
					b := msg.([]byte)
					code := make([]byte, 6)
					copy(code, b[0:6])

					var volume int64
					volume = 0
					for i := 0; i < 16; i++ {
						if b[16+i] == 49 {
							volume += powb(10, 15-i)
						} else if b[16+i] == 50 {
							volume += 2 * powb(10, 15-i)
						} else if b[16+i] == 51 {
							volume += 3 * powb(10, 15-i)
						} else if b[16+i] == 52 {
							volume += 4 * powb(10, 15-i)
						} else if b[16+i] == 53 {
							volume += 5 * powb(10, 15-i)
						} else if b[16+i] == 54 {
							volume += 6 * powb(10, 15-i)
						} else if b[16+i] == 55 {
							volume += 7 * powb(10, 15-i)
						} else if b[16+i] == 56 {
							volume += 8 * powb(10, 15-i)
						} else if b[16+i] == 57 {
							volume += 9 * powb(10, 15-i)
						} else {
						}
					}
					var amount1 int64
					amount1 = 0
					for i := 0; i < 15; i++ {
						if i == 14 || i == 13 {
							continue
						}
						if b[33+i] == 49 {
							amount1 += powb(10, 14-i)
						} else if b[33+i] == 50 {
							amount1 += 2 * powb(10, 14-i)
						} else if b[33+i] == 51 {
							amount1 += 3 * powb(10, 14-i)
						} else if b[33+i] == 52 {
							amount1 += 4 * powb(10, 14-i)
						} else if b[33+i] == 53 {
							amount1 += 5 * powb(10, 14-i)
						} else if b[33+i] == 54 {
							amount1 += 6 * powb(10, 14-i)
						} else if b[33+i] == 55 {
							amount1 += 7 * powb(10, 14-i)
						} else if b[33+i] == 56 {
							amount1 += 8 * powb(10, 14-i)
						} else if b[33+i] == 57 {
							amount1 += 9 * powb(10, 14-i)
						} else {
						}
					}

					for i := 0; i < 2; i++ {
						if b[47+i] == 49 {
							amount1 += powb(10, 1-i)
						} else if b[47+i] == 50 {
							amount1 += 2 * powb(10, 1-i)
						} else if b[47+i] == 51 {
							amount1 += 3 * powb(10, 1-i)
						} else if b[47+i] == 52 {
							amount1 += 4 * powb(10, 1-i)
						} else if b[47+i] == 53 {
							amount1 += 5 * powb(10, 1-i)
						} else if b[47+i] == 54 {
							amount1 += 6 * powb(10, 1-i)
						} else if b[47+i] == 55 {
							amount1 += 7 * powb(10, 1-i)
						} else if b[47+i] == 56 {
							amount1 += 8 * powb(10, 1-i)
						} else if b[47+i] == 57 {
							amount1 += 9 * powb(10, 1-i)
						} else {
						}
					}

					amount := amount1
					var lastprice1 int64
					lastprice1 = 0
					for i := 0; i < 10; i++ {
						if i == 6 || i == 7 || i == 8 || i == 9 {
							continue
						}
						if b[50+i] == 49 {
							lastprice1 += powb(10, 9-i)
						} else if b[50+i] == 50 {
							lastprice1 += 2 * powb(10, 9-i)
						} else if b[50+i] == 51 {
							lastprice1 += 3 * powb(10, 9-i)
						} else if b[50+i] == 52 {
							lastprice1 += 4 * powb(10, 9-i)
						} else if b[50+i] == 53 {
							lastprice1 += 5 * powb(10, 9-i)
						} else if b[50+i] == 54 {
							lastprice1 += 6 * powb(10, 9-i)
						} else if b[50+i] == 55 {
							lastprice1 += 7 * powb(10, 9-i)
						} else if b[50+i] == 56 {
							lastprice1 += 8 * powb(10, 9-i)
						} else if b[50+i] == 57 {
							lastprice1 += 9 * powb(10, 9-i)
						} else {
						}
					}
					for i := 0; i < 4; i++ {
						if b[57+i] == 49 {
							lastprice1 += powb(10, 3-i)
						} else if b[57+i] == 50 {
							lastprice1 += 2 * powb(10, 3-i)
						} else if b[57+i] == 51 {
							lastprice1 += 3 * powb(10, 3-i)
						} else if b[57+i] == 52 {
							lastprice1 += 4 * powb(10, 3-i)
						} else if b[57+i] == 53 {
							lastprice1 += 5 * powb(10, 3-i)
						} else if b[57+i] == 54 {
							lastprice1 += 6 * powb(10, 3-i)
						} else if b[57+i] == 55 {
							lastprice1 += 7 * powb(10, 3-i)
						} else if b[57+i] == 56 {
							lastprice1 += 8 * powb(10, 3-i)
						} else if b[57+i] == 57 {
							lastprice1 += 9 * powb(10, 3-i)
						} else {
						}
					}

					lastprice := lastprice1
					var open1 int64
					open1 = 0
					for i := 0; i < 10; i++ {
						if i == 6 || i == 7 || i == 8 || i == 9 {
							continue
						}
						if b[62+i] == 49 {
							open1 += powb(10, 9-i)
						} else if b[62+i] == 50 {
							open1 += 2 * powb(10, 9-i)
						} else if b[62+i] == 51 {
							open1 += 3 * powb(10, 9-i)
						} else if b[62+i] == 52 {
							open1 += 4 * powb(10, 9-i)
						} else if b[62+i] == 53 {
							open1 += 5 * powb(10, 9-i)
						} else if b[62+i] == 54 {
							open1 += 6 * powb(10, 9-i)
						} else if b[62+i] == 55 {
							open1 += 7 * powb(10, 9-i)
						} else if b[62+i] == 56 {
							open1 += 8 * powb(10, 9-i)
						} else if b[62+i] == 57 {
							open1 += 9 * powb(10, 9-i)
						} else {
						}
					}
					for i := 0; i < 4; i++ {
						if b[69+i] == 49 {
							open1 += powb(10, 3-i)
						} else if b[69+i] == 50 {
							open1 += 2 * powb(10, 3-i)
						} else if b[69+i] == 51 {
							open1 += 3 * powb(10, 3-i)
						} else if b[69+i] == 52 {
							open1 += 4 * powb(10, 3-i)
						} else if b[69+i] == 53 {
							open1 += 5 * powb(10, 3-i)
						} else if b[69+i] == 54 {
							open1 += 6 * powb(10, 3-i)
						} else if b[69+i] == 55 {
							open1 += 7 * powb(10, 3-i)
						} else if b[69+i] == 56 {
							open1 += 8 * powb(10, 3-i)
						} else if b[69+i] == 57 {
							open1 += 9 * powb(10, 3-i)
						} else {
						}
					}

					open := open1
					var high1 int64
					high1 = 0
					for i := 0; i < 10; i++ {
						if i == 6 || i == 7 || i == 8 || i == 9 {
							continue
						}
						if b[74+i] == 49 {
							high1 += powb(10, 9-i)
						} else if b[74+i] == 50 {
							high1 += 2 * powb(10, 9-i)
						} else if b[74+i] == 51 {
							high1 += 3 * powb(10, 9-i)
						} else if b[74+i] == 52 {
							high1 += 4 * powb(10, 9-i)
						} else if b[74+i] == 53 {
							high1 += 5 * powb(10, 9-i)
						} else if b[74+i] == 54 {
							high1 += 6 * powb(10, 9-i)
						} else if b[74+i] == 55 {
							high1 += 7 * powb(10, 9-i)
						} else if b[74+i] == 56 {
							high1 += 8 * powb(10, 9-i)
						} else if b[74+i] == 57 {
							high1 += 9 * powb(10, 9-i)
						} else {
						}
					}

					for i := 0; i < 4; i++ {
						if b[81+i] == 49 {
							high1 += powb(10, 3-i)
						} else if b[81+i] == 50 {
							high1 += 2 * powb(10, 3-i)
						} else if b[81+i] == 51 {
							high1 += 3 * powb(10, 3-i)
						} else if b[81+i] == 52 {
							high1 += 4 * powb(10, 3-i)
						} else if b[81+i] == 53 {
							high1 += 5 * powb(10, 3-i)
						} else if b[81+i] == 54 {
							high1 += 6 * powb(10, 3-i)
						} else if b[81+i] == 55 {
							high1 += 7 * powb(10, 3-i)
						} else if b[81+i] == 56 {
							high1 += 8 * powb(10, 3-i)
						} else if b[81+i] == 57 {
							high1 += 9 * powb(10, 3-i)
						} else {
						}
					}

					high := high1
					var low1 int64
					low1 = 0
					for i := 0; i < 10; i++ {
						if i == 6 || i == 7 || i == 8 || i == 9 {
							continue
						}
						if b[86+i] == 49 {
							low1 += powb(10, 9-i)
						} else if b[86+i] == 50 {
							low1 += 2 * powb(10, 9-i)
						} else if b[86+i] == 51 {
							low1 += 3 * powb(10, 9-i)
						} else if b[86+i] == 52 {
							low1 += 4 * powb(10, 9-i)
						} else if b[86+i] == 53 {
							low1 += 5 * powb(10, 9-i)
						} else if b[86+i] == 54 {
							low1 += 6 * powb(10, 9-i)
						} else if b[86+i] == 55 {
							low1 += 7 * powb(10, 9-i)
						} else if b[86+i] == 56 {
							low1 += 8 * powb(10, 9-i)
						} else if b[86+i] == 57 {
							low1 += 9 * powb(10, 9-i)
						} else {
						}
					}
					for i := 0; i < 4; i++ {
						if b[93+i] == 49 {
							low1 += powb(10, 3-i)
						} else if b[93+i] == 50 {
							low1 += 2 * powb(10, 3-i)
						} else if b[93+i] == 51 {
							low1 += 3 * powb(10, 3-i)
						} else if b[93+i] == 52 {
							low1 += 4 * powb(10, 3-i)
						} else if b[93+i] == 53 {
							low1 += 5 * powb(10, 3-i)
						} else if b[93+i] == 54 {
							low1 += 6 * powb(10, 3-i)
						} else if b[93+i] == 55 {
							low1 += 7 * powb(10, 3-i)
						} else if b[93+i] == 56 {
							low1 += 8 * powb(10, 3-i)
						} else if b[93+i] == 57 {
							low1 += 9 * powb(10, 3-i)
						} else {
						}
					}

					low := low1
					var tradeprice1 int64
					tradeprice1 = 0
					for i := 0; i < 10; i++ {
						if i == 6 || i == 7 || i == 8 || i == 9 {
							continue
						}
						if b[98+i] == 49 {
							tradeprice1 += powb(10, 9-i)
						} else if b[98+i] == 50 {
							tradeprice1 += 2 * powb(10, 9-i)
						} else if b[98+i] == 51 {
							tradeprice1 += 3 * powb(10, 9-i)
						} else if b[98+i] == 52 {
							tradeprice1 += 4 * powb(10, 9-i)
						} else if b[98+i] == 53 {
							tradeprice1 += 5 * powb(10, 9-i)
						} else if b[98+i] == 54 {
							tradeprice1 += 6 * powb(10, 9-i)
						} else if b[98+i] == 55 {
							tradeprice1 += 7 * powb(10, 9-i)
						} else if b[98+i] == 56 {
							tradeprice1 += 8 * powb(10, 9-i)
						} else if b[98+i] == 57 {
							tradeprice1 += 9 * powb(10, 9-i)
						} else {
						}
					}
					for i := 0; i < 4; i++ {
						if b[105+i] == 49 {
							tradeprice1 += powb(10, 3-i)
						} else if b[105+i] == 50 {
							tradeprice1 += 2 * powb(10, 3-i)
						} else if b[105+i] == 51 {
							tradeprice1 += 3 * powb(10, 3-i)
						} else if b[105+i] == 52 {
							tradeprice1 += 4 * powb(10, 3-i)
						} else if b[105+i] == 53 {
							tradeprice1 += 5 * powb(10, 3-i)
						} else if b[105+i] == 54 {
							tradeprice1 += 6 * powb(10, 3-i)
						} else if b[105+i] == 55 {
							tradeprice1 += 7 * powb(10, 3-i)
						} else if b[105+i] == 56 {
							tradeprice1 += 8 * powb(10, 3-i)
						} else if b[105+i] == 57 {
							tradeprice1 += 9 * powb(10, 3-i)
						} else {
						}
					}

					tradeprice := tradeprice1
					var closepx1 int64
					closepx1 = 0
					for i := 0; i < 10; i++ {
						if i == 6 || i == 7 || i == 8 || i == 9 {
							continue
						}
						if b[110+i] == 49 {
							closepx1 += powb(10, 9-i)
						} else if b[110+i] == 50 {
							closepx1 += 2 * powb(10, 9-i)
						} else if b[110+i] == 51 {
							closepx1 += 3 * powb(10, 9-i)
						} else if b[110+i] == 52 {
							closepx1 += 4 * powb(10, 9-i)
						} else if b[110+i] == 53 {
							closepx1 += 5 * powb(10, 9-i)
						} else if b[110+i] == 54 {
							closepx1 += 6 * powb(10, 9-i)
						} else if b[110+i] == 55 {
							closepx1 += 7 * powb(10, 9-i)
						} else if b[110+i] == 56 {
							closepx1 += 8 * powb(10, 9-i)
						} else if b[110+i] == 57 {
							closepx1 += 9 * powb(10, 9-i)
						} else {
						}
					}

					for i := 0; i < 4; i++ {
						if b[117+i] == 49 {
							closepx1 += powb(10, 3-i)
						} else if b[117+i] == 50 {
							closepx1 += 2 * powb(10, 3-i)
						} else if b[117+i] == 51 {
							closepx1 += 3 * powb(10, 3-i)
						} else if b[117+i] == 52 {
							closepx1 += 4 * powb(10, 3-i)
						} else if b[117+i] == 53 {
							closepx1 += 5 * powb(10, 3-i)
						} else if b[117+i] == 54 {
							closepx1 += 6 * powb(10, 3-i)
						} else if b[117+i] == 55 {
							closepx1 += 7 * powb(10, 3-i)
						} else if b[117+i] == 56 {
							closepx1 += 8 * powb(10, 3-i)
						} else if b[117+i] == 57 {
							closepx1 += 9 * powb(10, 3-i)
						} else {
						}
					}

					closepx := closepx1

					status1 := b[122]
					var status2 int64
					status2 = 0
					for i := 0; i < 3; i++ {
						if b[123+i] == 49 {
							status2 += powb(10, 2-i)
						} else {
						}
					}
					timestamp := make([]byte, 9)
					//copy(timestamp, b[131:143])
					copy(timestamp[0:2], b[131:133])
					copy(timestamp[2:4], b[134:136])
					copy(timestamp[4:6], b[137:139])
					copy(timestamp[6:9], b[140:143])

					debug := false
					if debug {
						fmt.Printf("\n---[code|%v] [volume|%d] [amount|%f] [lastprice|%f]", code, volume, amount, lastprice)
						fmt.Printf(" [open|%f] [high|%f] [low|%f] [tradeprice|%f] [closepx|%f]", open, high, low, tradeprice, closepx)
						fmt.Printf(" [status1|%v] [status2|%d]", status1, status2)
						fmt.Printf("---\n")
					}
					//msgpack序列化
					_code := string(code) + ".SH"

					timestamp_buf, err := strconv.ParseInt(string(timestamp), 10, 32)
					if err != nil {
						log.Error("string to int32 err: ", err)
					}
					_time := int32(timestamp_buf)

					_preClose := float64(lastprice) / 10000.0
					_open := float64(open) / 10000.0
					_high := float64(high) / 10000.0
					_low := float64(low) / 10000.0
					_match := float64(tradeprice) / 10000.0
					_turnover := amount / 100
					debug = false
					if debug {
						fmt.Printf("[code|%v] [time|%d] [preClose|%f]", _code, _time, _preClose)
						fmt.Printf(" [open|%f] [high|%f] [low|%f] [match|%f]", _open, _high, _low, _match)
						fmt.Printf(" [volume|%d] [turnover|%d]\n", volume, _turnover)
					}

					marketData := hqbase.Marketdata{Code: _code, Time: _time, PreClose: _preClose, Open: _open, High: _high, Low: _low, Match: _match, Volume: volume, Turnover: _turnover}

					data, err := msgpack.Marshal(&marketData)
					if err != nil {
						panic(err)
					}

					debug = false
					if debug {
						fmt.Printf("[data|%v]\n", data)
						var mData hqbase.Marketdata
						err = msgpack.Unmarshal(data, &mData)
						if err != nil {
							panic(err)
						}
						fmt.Printf("Ummarshal:[code|%s] [time|%d] [preClose|%f]", mData.Code, mData.Time, mData.PreClose)
						fmt.Printf(" [open|%f] [high|%f] [low|%f] match|%f]", mData.Open, mData.High, mData.Low, mData.Match)
						fmt.Printf(" [volume|%d] [turnover|%d]\n", mData.Volume, mData.Turnover)
					}
					if count == 10 {
						ct2 := time.Now().Sub(tt1)
						fmt.Printf("%d %d %v\n", seq, count, ct2)
					}
					redisRb.Put(data)
					rb.Put(data)
					// mysqlRb.Put(marketData)
				} else {
					time.Sleep(time.Duration(1) * time.Millisecond)
				}

			}

		}(i)
	}
}

func Md002map(redisRb *queue.RingBuffer, rb *queue.RingBuffer, mysqlRb *queue.RingBuffer) {
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(seq int) {
			count := 0
			fmt.Println("start ", seq)
			tt1 := time.Now()
			for {
				if rbmd002map.Len() > 0 {
					count++
					msg, _ := rbmd002map.Get()
					b := msg.([]byte)
					code := make([]byte, 6)
					copy(code, b[0:6])

					var volume int64
					volume = 0
					for i := 0; i < 16; i++ {
						if b[16+i] == 49 {
							volume += powb(10, 15-i)
						} else if b[16+i] == 50 {
							volume += 2 * powb(10, 15-i)
						} else if b[16+i] == 51 {
							volume += 3 * powb(10, 15-i)
						} else if b[16+i] == 52 {
							volume += 4 * powb(10, 15-i)
						} else if b[16+i] == 53 {
							volume += 5 * powb(10, 15-i)
						} else if b[16+i] == 54 {
							volume += 6 * powb(10, 15-i)
						} else if b[16+i] == 55 {
							volume += 7 * powb(10, 15-i)
						} else if b[16+i] == 56 {
							volume += 8 * powb(10, 15-i)
						} else if b[16+i] == 57 {
							volume += 9 * powb(10, 15-i)
						} else {
						}
					}

					var amount int64
					amount = 0
					for i := 0; i < 16; i++ {
						n := 0
						if i == 13 {
							continue
						} else if i > 13 {
							n = i - 1
						} else {
							n = i
						}

						if b[33+i] == 49 {
							amount += powb(10, 14-n)
						} else if b[33+i] == 50 {
							amount += 2 * powb(10, 14-n)
						} else if b[33+i] == 51 {
							amount += 3 * powb(10, 14-n)
						} else if b[33+i] == 52 {
							amount += 4 * powb(10, 14-n)
						} else if b[33+i] == 53 {
							amount += 5 * powb(10, 14-n)
						} else if b[33+i] == 54 {
							amount += 6 * powb(10, 14-n)
						} else if b[33+i] == 55 {
							amount += 7 * powb(10, 14-n)
						} else if b[33+i] == 56 {
							amount += 8 * powb(10, 14-n)
						} else if b[33+i] == 57 {
							amount += 9 * powb(10, 14-n)
						} else {
						}
					}

					var lastprice int64
					lastprice = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[50+i] == 49 {
							lastprice += powb(10, 9-n)
						} else if b[50+i] == 50 {
							lastprice += 2 * powb(10, 9-n)
						} else if b[50+i] == 51 {
							lastprice += 3 * powb(10, 9-n)
						} else if b[50+i] == 52 {
							lastprice += 4 * powb(10, 9-n)
						} else if b[50+i] == 53 {
							lastprice += 5 * powb(10, 9-n)
						} else if b[50+i] == 54 {
							lastprice += 6 * powb(10, 9-n)
						} else if b[50+i] == 55 {
							lastprice += 7 * powb(10, 9-n)
						} else if b[50+i] == 56 {
							lastprice += 8 * powb(10, 9-n)
						} else if b[50+i] == 57 {
							lastprice += 9 * powb(10, 9-n)
						} else {
						}
					}

					var open int64
					open = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[62+i] == 49 {
							open += powb(10, 9-n)
						} else if b[62+i] == 50 {
							open += 2 * powb(10, 9-n)
						} else if b[62+i] == 51 {
							open += 3 * powb(10, 9-n)
						} else if b[62+i] == 52 {
							open += 4 * powb(10, 9-n)
						} else if b[62+i] == 53 {
							open += 5 * powb(10, 9-n)
						} else if b[62+i] == 54 {
							open += 6 * powb(10, 9-n)
						} else if b[62+i] == 55 {
							open += 7 * powb(10, 9-n)
						} else if b[62+i] == 56 {
							open += 8 * powb(10, 9-n)
						} else if b[62+i] == 57 {
							open += 9 * powb(10, 9-n)
						} else {
						}
					}

					var high int64
					high = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[74+i] == 49 {
							high += powb(10, 9-n)
						} else if b[74+i] == 50 {
							high += 2 * powb(10, 9-n)
						} else if b[74+i] == 51 {
							high += 3 * powb(10, 9-n)
						} else if b[74+i] == 52 {
							high += 4 * powb(10, 9-n)
						} else if b[74+i] == 53 {
							high += 5 * powb(10, 9-n)
						} else if b[74+i] == 54 {
							high += 6 * powb(10, 9-n)
						} else if b[74+i] == 55 {
							high += 7 * powb(10, 9-n)
						} else if b[74+i] == 56 {
							high += 8 * powb(10, 9-n)
						} else if b[74+i] == 57 {
							high += 9 * powb(10, 9-n)
						} else {
						}
					}

					var low int64
					low = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[86+i] == 49 {
							low += powb(10, 9-n)
						} else if b[86+i] == 50 {
							low += 2 * powb(10, 9-n)
						} else if b[86+i] == 51 {
							low += 3 * powb(10, 9-n)
						} else if b[86+i] == 52 {
							low += 4 * powb(10, 9-n)
						} else if b[86+i] == 53 {
							low += 5 * powb(10, 9-n)
						} else if b[86+i] == 54 {
							low += 6 * powb(10, 9-n)
						} else if b[86+i] == 55 {
							low += 7 * powb(10, 9-n)
						} else if b[86+i] == 56 {
							low += 8 * powb(10, 9-n)
						} else if b[86+i] == 57 {
							low += 9 * powb(10, 9-n)
						} else {
						}
					}

					var tradeprice int64
					tradeprice = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[98+i] == 49 {
							tradeprice += powb(10, 9-n)
						} else if b[98+i] == 50 {
							tradeprice += 2 * powb(10, 9-n)
						} else if b[98+i] == 51 {
							tradeprice += 3 * powb(10, 9-n)
						} else if b[98+i] == 52 {
							tradeprice += 4 * powb(10, 9-n)
						} else if b[98+i] == 53 {
							tradeprice += 5 * powb(10, 9-n)
						} else if b[98+i] == 54 {
							tradeprice += 6 * powb(10, 9-n)
						} else if b[98+i] == 55 {
							tradeprice += 7 * powb(10, 9-n)
						} else if b[98+i] == 56 {
							tradeprice += 8 * powb(10, 9-n)
						} else if b[98+i] == 57 {
							tradeprice += 9 * powb(10, 9-n)
						} else {
						}
					}

					var closepx int64
					closepx = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[110+i] == 49 {
							closepx += powb(10, 9-n)
						} else if b[110+i] == 50 {
							closepx += 2 * powb(10, 9-n)
						} else if b[110+i] == 51 {
							closepx += 3 * powb(10, 9-n)
						} else if b[110+i] == 52 {
							closepx += 4 * powb(10, 9-n)
						} else if b[110+i] == 53 {
							closepx += 5 * powb(10, 9-n)
						} else if b[110+i] == 54 {
							closepx += 6 * powb(10, 9-n)
						} else if b[110+i] == 55 {
							closepx += 7 * powb(10, 9-n)
						} else if b[110+i] == 56 {
							closepx += 8 * powb(10, 9-n)
						} else if b[110+i] == 57 {
							closepx += 9 * powb(10, 9-n)
						} else {
						}
					}

					var bp1 int64
					bp1 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[122+i] == 49 {
							bp1 += powb(10, 9-n)
						} else if b[122+i] == 50 {
							bp1 += 2 * powb(10, 9-n)
						} else if b[122+i] == 51 {
							bp1 += 3 * powb(10, 9-n)
						} else if b[122+i] == 52 {
							bp1 += 4 * powb(10, 9-n)
						} else if b[122+i] == 53 {
							bp1 += 5 * powb(10, 9-n)
						} else if b[122+i] == 54 {
							bp1 += 6 * powb(10, 9-n)
						} else if b[122+i] == 55 {
							bp1 += 7 * powb(10, 9-n)
						} else if b[122+i] == 56 {
							bp1 += 8 * powb(10, 9-n)
						} else if b[122+i] == 57 {
							bp1 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var bv1 int64
					bv1 = 0
					for i := 0; i < 12; i++ {
						if b[134+i] == 49 {
							bv1 += powb(10, 11-i)
						} else if b[134+i] == 50 {
							bv1 += 2 * powb(10, 11-i)
						} else if b[134+i] == 51 {
							bv1 += 3 * powb(10, 11-i)
						} else if b[134+i] == 52 {
							bv1 += 4 * powb(10, 11-i)
						} else if b[134+i] == 53 {
							bv1 += 5 * powb(10, 11-i)
						} else if b[134+i] == 54 {
							bv1 += 6 * powb(10, 11-i)
						} else if b[134+i] == 55 {
							bv1 += 7 * powb(10, 11-i)
						} else if b[134+i] == 56 {
							bv1 += 8 * powb(10, 11-i)
						} else if b[134+i] == 57 {
							bv1 += 9 * powb(10, 11-i)
						} else {
						}
					}

					var sp1 int64
					sp1 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[147+i] == 49 {
							sp1 += powb(10, 9-n)
						} else if b[147+i] == 50 {
							sp1 += 2 * powb(10, 9-n)
						} else if b[147+i] == 51 {
							sp1 += 3 * powb(10, 9-n)
						} else if b[147+i] == 52 {
							sp1 += 4 * powb(10, 9-n)
						} else if b[147+i] == 53 {
							sp1 += 5 * powb(10, 9-n)
						} else if b[147+i] == 54 {
							sp1 += 6 * powb(10, 9-n)
						} else if b[147+i] == 55 {
							sp1 += 7 * powb(10, 9-n)
						} else if b[147+i] == 56 {
							sp1 += 8 * powb(10, 9-n)
						} else if b[147+i] == 57 {
							sp1 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var sv1 int64
					sv1 = 0
					for i := 0; i < 12; i++ {
						if b[159+i] == 49 {
							sv1 += powb(10, 11-i)
						} else if b[159+i] == 50 {
							sv1 += 2 * powb(10, 11-i)
						} else if b[159+i] == 51 {
							sv1 += 3 * powb(10, 11-i)
						} else if b[159+i] == 52 {
							sv1 += 4 * powb(10, 11-i)
						} else if b[159+i] == 53 {
							sv1 += 5 * powb(10, 11-i)
						} else if b[159+i] == 54 {
							sv1 += 6 * powb(10, 11-i)
						} else if b[159+i] == 55 {
							sv1 += 7 * powb(10, 11-i)
						} else if b[159+i] == 56 {
							sv1 += 8 * powb(10, 11-i)
						} else if b[159+i] == 57 {
							sv1 += 9 * powb(10, 11-i)
						} else {
						}
					}

					var bp2 int64
					bp2 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[172+i] == 49 {
							bp2 += powb(10, 9-n)
						} else if b[172+i] == 50 {
							bp2 += 2 * powb(10, 9-n)
						} else if b[172+i] == 51 {
							bp2 += 3 * powb(10, 9-n)
						} else if b[172+i] == 52 {
							bp2 += 4 * powb(10, 9-n)
						} else if b[172+i] == 53 {
							bp2 += 5 * powb(10, 9-n)
						} else if b[172+i] == 54 {
							bp2 += 6 * powb(10, 9-n)
						} else if b[172+i] == 55 {
							bp2 += 7 * powb(10, 9-n)
						} else if b[172+i] == 56 {
							bp2 += 8 * powb(10, 9-n)
						} else if b[172+i] == 57 {
							bp2 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var bv2 int64
					bv2 = 0
					for i := 0; i < 12; i++ {
						if b[184+i] == 49 {
							bv2 += powb(10, 11-i)
						} else if b[184+i] == 50 {
							bv2 += 2 * powb(10, 11-i)
						} else if b[184+i] == 51 {
							bv2 += 3 * powb(10, 11-i)
						} else if b[184+i] == 52 {
							bv2 += 4 * powb(10, 11-i)
						} else if b[184+i] == 53 {
							bv2 += 5 * powb(10, 11-i)
						} else if b[184+i] == 54 {
							bv2 += 6 * powb(10, 11-i)
						} else if b[184+i] == 55 {
							bv2 += 7 * powb(10, 11-i)
						} else if b[184+i] == 56 {
							bv2 += 8 * powb(10, 11-i)
						} else if b[184+i] == 57 {
							bv2 += 9 * powb(10, 11-i)
						} else {
						}
					}

					var sp2 int64
					sp2 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[197+i] == 49 {
							sp2 += powb(10, 9-n)
						} else if b[197+i] == 50 {
							sp2 += 2 * powb(10, 9-n)
						} else if b[197+i] == 51 {
							sp2 += 3 * powb(10, 9-n)
						} else if b[197+i] == 52 {
							sp2 += 4 * powb(10, 9-n)
						} else if b[197+i] == 53 {
							sp2 += 5 * powb(10, 9-n)
						} else if b[197+i] == 54 {
							sp2 += 6 * powb(10, 9-n)
						} else if b[197+i] == 55 {
							sp2 += 7 * powb(10, 9-n)
						} else if b[197+i] == 56 {
							sp2 += 8 * powb(10, 9-n)
						} else if b[197+i] == 57 {
							sp2 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var sv2 int64
					sv2 = 0
					for i := 0; i < 12; i++ {
						if b[209+i] == 49 {
							sv2 += powb(10, 11-i)
						} else if b[209+i] == 50 {
							sv2 += 2 * powb(10, 11-i)
						} else if b[209+i] == 51 {
							sv2 += 3 * powb(10, 11-i)
						} else if b[209+i] == 52 {
							sv2 += 4 * powb(10, 11-i)
						} else if b[209+i] == 53 {
							sv2 += 5 * powb(10, 11-i)
						} else if b[209+i] == 54 {
							sv2 += 6 * powb(10, 11-i)
						} else if b[209+i] == 55 {
							sv2 += 7 * powb(10, 11-i)
						} else if b[209+i] == 56 {
							sv2 += 8 * powb(10, 11-i)
						} else if b[209+i] == 57 {
							sv2 += 9 * powb(10, 11-i)
						} else {
						}
					}

					var bp3 int64
					bp3 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[222+i] == 49 {
							bp3 += powb(10, 9-n)
						} else if b[222+i] == 50 {
							bp3 += 2 * powb(10, 9-n)
						} else if b[222+i] == 51 {
							bp3 += 3 * powb(10, 9-n)
						} else if b[222+i] == 52 {
							bp3 += 4 * powb(10, 9-n)
						} else if b[222+i] == 53 {
							bp3 += 5 * powb(10, 9-n)
						} else if b[222+i] == 54 {
							bp3 += 6 * powb(10, 9-n)
						} else if b[222+i] == 55 {
							bp3 += 7 * powb(10, 9-n)
						} else if b[222+i] == 56 {
							bp3 += 8 * powb(10, 9-n)
						} else if b[222+i] == 57 {
							bp3 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var bv3 int64
					bv3 = 0
					for i := 0; i < 12; i++ {
						if b[234+i] == 49 {
							bv3 += powb(10, 11-i)
						} else if b[234+i] == 50 {
							bv3 += 2 * powb(10, 11-i)
						} else if b[234+i] == 51 {
							bv3 += 3 * powb(10, 11-i)
						} else if b[234+i] == 52 {
							bv3 += 4 * powb(10, 11-i)
						} else if b[234+i] == 53 {
							bv3 += 5 * powb(10, 11-i)
						} else if b[234+i] == 54 {
							bv3 += 6 * powb(10, 11-i)
						} else if b[234+i] == 55 {
							bv3 += 7 * powb(10, 11-i)
						} else if b[234+i] == 56 {
							bv3 += 8 * powb(10, 11-i)
						} else if b[234+i] == 57 {
							bv3 += 9 * powb(10, 11-i)
						} else {
						}
					}

					var sp3 int64
					sp3 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[247+i] == 49 {
							sp3 += powb(10, 9-n)
						} else if b[247+i] == 50 {
							sp3 += 2 * powb(10, 9-n)
						} else if b[247+i] == 51 {
							sp3 += 3 * powb(10, 9-n)
						} else if b[247+i] == 52 {
							sp3 += 4 * powb(10, 9-n)
						} else if b[247+i] == 53 {
							sp3 += 5 * powb(10, 9-n)
						} else if b[247+i] == 54 {
							sp3 += 6 * powb(10, 9-n)
						} else if b[247+i] == 55 {
							sp3 += 7 * powb(10, 9-n)
						} else if b[247+i] == 56 {
							sp3 += 8 * powb(10, 9-n)
						} else if b[247+i] == 57 {
							sp3 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var sv3 int64
					sv3 = 0
					for i := 0; i < 12; i++ {
						if b[259+i] == 49 {
							sv3 += powb(10, 11-i)
						} else if b[259+i] == 50 {
							sv3 += 2 * powb(10, 11-i)
						} else if b[259+i] == 51 {
							sv3 += 3 * powb(10, 11-i)
						} else if b[259+i] == 52 {
							sv3 += 4 * powb(10, 11-i)
						} else if b[259+i] == 53 {
							sv3 += 5 * powb(10, 11-i)
						} else if b[259+i] == 54 {
							sv3 += 6 * powb(10, 11-i)
						} else if b[259+i] == 55 {
							sv3 += 7 * powb(10, 11-i)
						} else if b[259+i] == 56 {
							sv3 += 8 * powb(10, 11-i)
						} else if b[259+i] == 57 {
							sv3 += 9 * powb(10, 11-i)
						} else {
						}
					}

					var bp4 int64
					bp4 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[272+i] == 49 {
							bp4 += powb(10, 9-n)
						} else if b[272+i] == 50 {
							bp4 += 2 * powb(10, 9-n)
						} else if b[272+i] == 51 {
							bp4 += 3 * powb(10, 9-n)
						} else if b[272+i] == 52 {
							bp4 += 4 * powb(10, 9-n)
						} else if b[272+i] == 53 {
							bp4 += 5 * powb(10, 9-n)
						} else if b[272+i] == 54 {
							bp4 += 6 * powb(10, 9-n)
						} else if b[272+i] == 55 {
							bp4 += 7 * powb(10, 9-n)
						} else if b[272+i] == 56 {
							bp4 += 8 * powb(10, 9-n)
						} else if b[272+i] == 57 {
							bp4 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var bv4 int64
					bv4 = 0
					for i := 0; i < 12; i++ {
						if b[284+i] == 49 {
							bv4 += powb(10, 11-i)
						} else if b[284+i] == 50 {
							bv4 += 2 * powb(10, 11-i)
						} else if b[284+i] == 51 {
							bv4 += 3 * powb(10, 11-i)
						} else if b[284+i] == 52 {
							bv4 += 4 * powb(10, 11-i)
						} else if b[284+i] == 53 {
							bv4 += 5 * powb(10, 11-i)
						} else if b[284+i] == 54 {
							bv4 += 6 * powb(10, 11-i)
						} else if b[284+i] == 55 {
							bv4 += 7 * powb(10, 11-i)
						} else if b[284+i] == 56 {
							bv4 += 8 * powb(10, 11-i)
						} else if b[284+i] == 57 {
							bv4 += 9 * powb(10, 11-i)
						} else {
						}
					}

					var sp4 int64
					sp4 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[297+i] == 49 {
							sp4 += powb(10, 9-n)
						} else if b[297+i] == 50 {
							sp4 += 2 * powb(10, 9-n)
						} else if b[297+i] == 51 {
							sp4 += 3 * powb(10, 9-n)
						} else if b[297+i] == 52 {
							sp4 += 4 * powb(10, 9-n)
						} else if b[297+i] == 53 {
							sp4 += 5 * powb(10, 9-n)
						} else if b[297+i] == 54 {
							sp4 += 6 * powb(10, 9-n)
						} else if b[297+i] == 55 {
							sp4 += 7 * powb(10, 9-n)
						} else if b[297+i] == 56 {
							sp4 += 8 * powb(10, 9-n)
						} else if b[297+i] == 57 {
							sp4 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var sv4 int64
					sv4 = 0
					for i := 0; i < 12; i++ {
						if b[309+i] == 49 {
							sv4 += powb(10, 11-i)
						} else if b[309+i] == 50 {
							sv4 += 2 * powb(10, 11-i)
						} else if b[309+i] == 51 {
							sv4 += 3 * powb(10, 11-i)
						} else if b[309+i] == 52 {
							sv4 += 4 * powb(10, 11-i)
						} else if b[309+i] == 53 {
							sv4 += 5 * powb(10, 11-i)
						} else if b[309+i] == 54 {
							sv4 += 6 * powb(10, 11-i)
						} else if b[309+i] == 55 {
							sv4 += 7 * powb(10, 11-i)
						} else if b[309+i] == 56 {
							sv4 += 8 * powb(10, 11-i)
						} else if b[309+i] == 57 {
							sv4 += 9 * powb(10, 11-i)
						} else {
						}
					}
					var bp5 int64
					bp5 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[322+i] == 49 {
							bp5 += powb(10, 9-n)
						} else if b[322+i] == 50 {
							bp5 += 2 * powb(10, 9-n)
						} else if b[322+i] == 51 {
							bp5 += 3 * powb(10, 9-n)
						} else if b[322+i] == 52 {
							bp5 += 4 * powb(10, 9-n)
						} else if b[322+i] == 53 {
							bp5 += 5 * powb(10, 9-n)
						} else if b[322+i] == 54 {
							bp5 += 6 * powb(10, 9-n)
						} else if b[322+i] == 55 {
							bp5 += 7 * powb(10, 9-n)
						} else if b[322+i] == 56 {
							bp5 += 8 * powb(10, 9-n)
						} else if b[322+i] == 57 {
							bp5 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var bv5 int64
					bv5 = 0
					for i := 0; i < 12; i++ {
						if b[334+i] == 49 {
							bv5 += powb(10, 11-i)
						} else if b[334+i] == 50 {
							bv5 += 2 * powb(10, 11-i)
						} else if b[334+i] == 51 {
							bv5 += 3 * powb(10, 11-i)
						} else if b[334+i] == 52 {
							bv5 += 4 * powb(10, 11-i)
						} else if b[334+i] == 53 {
							bv5 += 5 * powb(10, 11-i)
						} else if b[334+i] == 54 {
							bv5 += 6 * powb(10, 11-i)
						} else if b[334+i] == 55 {
							bv5 += 7 * powb(10, 11-i)
						} else if b[334+i] == 56 {
							bv5 += 8 * powb(10, 11-i)
						} else if b[334+i] == 57 {
							bv5 += 9 * powb(10, 11-i)
						} else {
						}
					}

					var sp5 int64
					sp5 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[347+i] == 49 {
							sp5 += powb(10, 9-n)
						} else if b[347+i] == 50 {
							sp5 += 2 * powb(10, 9-n)
						} else if b[347+i] == 51 {
							sp5 += 3 * powb(10, 9-n)
						} else if b[347+i] == 52 {
							sp5 += 4 * powb(10, 9-n)
						} else if b[347+i] == 53 {
							sp5 += 5 * powb(10, 9-n)
						} else if b[347+i] == 54 {
							sp5 += 6 * powb(10, 9-n)
						} else if b[347+i] == 55 {
							sp5 += 7 * powb(10, 9-n)
						} else if b[347+i] == 56 {
							sp5 += 8 * powb(10, 9-n)
						} else if b[347+i] == 57 {
							sp5 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var sv5 int64
					sv5 = 0
					for i := 0; i < 12; i++ {
						if b[359+i] == 49 {
							sv5 += powb(10, 11-i)
						} else if b[359+i] == 50 {
							sv5 += 2 * powb(10, 11-i)
						} else if b[359+i] == 51 {
							sv5 += 3 * powb(10, 11-i)
						} else if b[359+i] == 52 {
							sv5 += 4 * powb(10, 11-i)
						} else if b[359+i] == 53 {
							sv5 += 5 * powb(10, 11-i)
						} else if b[359+i] == 54 {
							sv5 += 6 * powb(10, 11-i)
						} else if b[359+i] == 55 {
							sv5 += 7 * powb(10, 11-i)
						} else if b[359+i] == 56 {
							sv5 += 8 * powb(10, 11-i)
						} else if b[359+i] == 57 {
							sv5 += 9 * powb(10, 11-i)
						} else {
						}
					}

					status1 := b[372]

					var status2 int64
					status2 = 0
					for i := 0; i < 3; i++ {
						if b[373+i] == 49 {
							status2 += powb(10, 2-i)
						} else {
						}
					}
					timestamp := make([]byte, 9)
					//copy(timestamp, b[381:393])
					copy(timestamp[0:2], b[381:383])
					copy(timestamp[2:4], b[384:386])
					copy(timestamp[4:6], b[387:389])
					copy(timestamp[6:9], b[390:393])

					debug := false
					if debug {
						fmt.Printf("\n---[code|%v] [volume|%d] [amount|%d] [lastprice|%d]", code, volume, amount, lastprice)
						fmt.Printf(" [open|%d] [high|%d] [low|%d] [tradeprice|%d] [closepx|%d]", open, high, low, tradeprice, closepx)
						fmt.Printf(" [bp1|%d] [bv1|%d] [sp1|%d] [sv1|%d]", bp1, bv1, sp1, sv1)
						fmt.Printf(" [bp2|%d] [bv2|%d] [sp2|%d] [sv2|%d]", bp2, bv2, sp2, sv2)
						fmt.Printf(" [bp3|%d] [bv3|%d] [sp3|%d] [sv3|%d]", bp3, bv3, sp3, sv3)
						fmt.Printf(" [bp4|%d] [bv4|%d] [sp4|%d] [sv4|%d]", bp4, bv4, sp4, sv4)
						fmt.Printf(" [bp5|%d] [bv5|%d] [sp5|%d] [sv5|%d]", bp5, bv5, sp5, sv5)
						fmt.Printf(" [status1|%v] [status2|%d]", status1, status2)
						fmt.Printf(" [timestamp|%v]", timestamp)
						fmt.Printf("---\n")
					}
					//msgpack序列化
					_code := string(code) + ".SH"

					timestamp_buf, err := strconv.ParseInt(string(timestamp), 10, 32)
					if err != nil {
						log.Error("string to int32 err: ", err)
					}
					_time := int32(timestamp_buf)

					var _status string
					if status1 == 0x43 { //C => I
						_status = "I"
					} else if status1 == 0x54 { //T => O
						_status = "O"
					} else if status1 == 0x50 { //P => B
						_status = "B"
					}

					_preClose := float64(lastprice) / 1000.00
					_open := float64(open) / 1000.00
					_high := float64(high) / 1000.00
					_low := float64(low) / 1000.00
					_match := float64(tradeprice) / 1000.00
					var _askPrice = [10]float64{float64(sp1) / 1000.0, float64(sp2) / 1000.0, float64(sp3) / 1000.0, float64(sp4) / 1000.0, float64(sp5) / 1000.0}
					var _askVol = [10]int32{int32(sv1), int32(sv2), int32(sv3), int32(sv4), int32(sv5)}
					var _bidPrice = [10]float64{float64(bp1) / 1000.0, float64(bp2) / 1000.0, float64(bp3) / 1000.0, float64(bp4) / 1000.0, float64(bp5) / 1000.0}
					var _bidVol = [10]int32{int32(bv1), int32(bv2), int32(bv3), int32(bv4), int32(bv5)}
					_turnover := amount / 100

					debug = false
					if debug {
						fmt.Printf("[code|%s] [time|%d] [status|%s] [preClose|%f]", _code, _time, _status, _preClose)
						fmt.Printf(" [open|%f] [high|%f] [low|%f] [match|%f]", _open, _high, _low, _match)
						fmt.Printf(" [askPrice|%v] [askVol|%v] [bidPrice|%v] [bidVol|%v]", _askPrice, _askVol, _bidPrice, _bidVol)
						fmt.Printf(" [volume|%d] [turnover|%d]\n", volume, _turnover)
					}

					marketData := hqbase.Marketdata{Code: _code, Time: _time, Status: _status, PreClose: _preClose, Open: _open, High: _high, Low: _low, Match: _match, AskPrice: _askPrice, AskVol: _askVol, BidPrice: _bidPrice, BidVol: _bidVol, Volume: volume, Turnover: _turnover}

					data, err := msgpack.Marshal(&marketData)
					if err != nil {
						panic(err)
					}

					debug = false
					if debug {
						fmt.Printf("data|%v\n", data)
						var mData hqbase.Marketdata
						err = msgpack.Unmarshal(data, &mData)
						if err != nil {
							panic(err)
						}
						fmt.Printf("Ummarshal:[code|%v] [time|%d] [status|%v] [preClose|%f]", mData.Code, mData.Time, mData.Status, mData.PreClose)
						fmt.Printf(" [open|%f] [high|%f] [low|%f] [match|%f]", mData.Open, mData.High, mData.Low, mData.Match)
						fmt.Printf(" [askPrice|%v] [askVol|%v] [bidPrice|%v] [bidVol|%v]", mData.AskPrice, mData.AskVol, mData.BidPrice, mData.BidVol)
						fmt.Printf(" [volume|%d] [turnover|%d]\n", mData.Volume, mData.Turnover)
					}
					if count == 10 {
						ct2 := time.Now().Sub(tt1)
						fmt.Printf("%d %d %v\n", seq, count, ct2)
					}
					redisRb.Put(data)
					rb.Put(data)
					// mysqlRb.Pu
					t(marketData)
				} else {
					time.Sleep(time.Duration(1) * time.Millisecond)
				}

			}

		}(i)
	}

}

func Md004map(redisRb *queue.RingBuffer, rb *queue.RingBuffer, mysqlRb *queue.RingBuffer) {
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(seq int) {
			count := 0
			fmt.Println("start ", seq)
			tt1 := time.Now()
			for {
				if rbmd004map.Len() > 0 {
					count++
					msg, _ := rbmd004map.Get()
					b := msg.([]byte)
					code := make([]byte, 6)
					copy(code, b[0:6])

					var volume int64
					volume = 0
					for i := 0; i < 16; i++ {
						if b[16+i] == 49 {
							volume += powb(10, 15-i)
						} else if b[16+i] == 50 {
							volume += 2 * powb(10, 15-i)
						} else if b[16+i] == 51 {
							volume += 3 * powb(10, 15-i)
						} else if b[16+i] == 52 {
							volume += 4 * powb(10, 15-i)
						} else if b[16+i] == 53 {
							volume += 5 * powb(10, 15-i)
						} else if b[16+i] == 54 {
							volume += 6 * powb(10, 15-i)
						} else if b[16+i] == 55 {
							volume += 7 * powb(10, 15-i)
						} else if b[16+i] == 56 {
							volume += 8 * powb(10, 15-i)
						} else if b[16+i] == 57 {
							volume += 9 * powb(10, 15-i)
						} else {
						}
					}
					var amount int64
					amount = 0
					for i := 0; i < 16; i++ {
						n := 0
						if i == 13 {
							continue
						} else if i > 13 {
							n = i - 1
						} else {
							n = i
						}

						if b[33+i] == 49 {
							amount += powb(10, 14-n)
						} else if b[33+i] == 50 {
							amount += 2 * powb(10, 14-n)
						} else if b[33+i] == 51 {
							amount += 3 * powb(10, 14-n)
						} else if b[33+i] == 52 {
							amount += 4 * powb(10, 14-n)
						} else if b[33+i] == 53 {
							amount += 5 * powb(10, 14-n)
						} else if b[33+i] == 54 {
							amount += 6 * powb(10, 14-n)
						} else if b[33+i] == 55 {
							amount += 7 * powb(10, 14-n)
						} else if b[33+i] == 56 {
							amount += 8 * powb(10, 14-n)
						} else if b[33+i] == 57 {
							amount += 9 * powb(10, 14-n)
						} else {
						}
					}

					var lastprice int64
					lastprice = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[50+i] == 49 {
							lastprice += powb(10, 9-n)
						} else if b[50+i] == 50 {
							lastprice += 2 * powb(10, 9-n)
						} else if b[50+i] == 51 {
							lastprice += 3 * powb(10, 9-n)
						} else if b[50+i] == 52 {
							lastprice += 4 * powb(10, 9-n)
						} else if b[50+i] == 53 {
							lastprice += 5 * powb(10, 9-n)
						} else if b[50+i] == 54 {
							lastprice += 6 * powb(10, 9-n)
						} else if b[50+i] == 55 {
							lastprice += 7 * powb(10, 9-n)
						} else if b[50+i] == 56 {
							lastprice += 8 * powb(10, 9-n)
						} else if b[50+i] == 57 {
							lastprice += 9 * powb(10, 9-n)
						} else {
						}
					}

					var open int64
					open = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[62+i] == 49 {
							open += powb(10, 9-n)
						} else if b[62+i] == 50 {
							open += 2 * powb(10, 9-n)
						} else if b[62+i] == 51 {
							open += 3 * powb(10, 9-n)
						} else if b[62+i] == 52 {
							open += 4 * powb(10, 9-n)
						} else if b[62+i] == 53 {
							open += 5 * powb(10, 9-n)
						} else if b[62+i] == 54 {
							open += 6 * powb(10, 9-n)
						} else if b[62+i] == 55 {
							open += 7 * powb(10, 9-n)
						} else if b[62+i] == 56 {
							open += 8 * powb(10, 9-n)
						} else if b[62+i] == 57 {
							open += 9 * powb(10, 9-n)
						} else {
						}
					}

					var high int64
					high = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[74+i] == 49 {
							high += powb(10, 9-n)
						} else if b[74+i] == 50 {
							high += 2 * powb(10, 9-n)
						} else if b[74+i] == 51 {
							high += 3 * powb(10, 9-n)
						} else if b[74+i] == 52 {
							high += 4 * powb(10, 9-n)
						} else if b[74+i] == 53 {
							high += 5 * powb(10, 9-n)
						} else if b[74+i] == 54 {
							high += 6 * powb(10, 9-n)
						} else if b[74+i] == 55 {
							high += 7 * powb(10, 9-n)
						} else if b[74+i] == 56 {
							high += 8 * powb(10, 9-n)
						} else if b[74+i] == 57 {
							high += 9 * powb(10, 9-n)
						} else {
						}
					}

					var low int64
					low = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[86+i] == 49 {
							low += powb(10, 9-n)
						} else if b[86+i] == 50 {
							low += 2 * powb(10, 9-n)
						} else if b[86+i] == 51 {
							low += 3 * powb(10, 9-n)
						} else if b[86+i] == 52 {
							low += 4 * powb(10, 9-n)
						} else if b[86+i] == 53 {
							low += 5 * powb(10, 9-n)
						} else if b[86+i] == 54 {
							low += 6 * powb(10, 9-n)
						} else if b[86+i] == 55 {
							low += 7 * powb(10, 9-n)
						} else if b[86+i] == 56 {
							low += 8 * powb(10, 9-n)
						} else if b[86+i] == 57 {
							low += 9 * powb(10, 9-n)
						} else {
						}
					}

					var tradeprice int64
					tradeprice = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[98+i] == 49 {
							tradeprice += powb(10, 9-n)
						} else if b[98+i] == 50 {
							tradeprice += 2 * powb(10, 9-n)
						} else if b[98+i] == 51 {
							tradeprice += 3 * powb(10, 9-n)
						} else if b[98+i] == 52 {
							tradeprice += 4 * powb(10, 9-n)
						} else if b[98+i] == 53 {
							tradeprice += 5 * powb(10, 9-n)
						} else if b[98+i] == 54 {
							tradeprice += 6 * powb(10, 9-n)
						} else if b[98+i] == 55 {
							tradeprice += 7 * powb(10, 9-n)
						} else if b[98+i] == 56 {
							tradeprice += 8 * powb(10, 9-n)
						} else if b[98+i] == 57 {
							tradeprice += 9 * powb(10, 9-n)
						} else {
						}
					}

					var closepx int64
					closepx = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[110+i] == 49 {
							closepx += powb(10, 9-n)
						} else if b[110+i] == 50 {
							closepx += 2 * powb(10, 9-n)
						} else if b[110+i] == 51 {
							closepx += 3 * powb(10, 9-n)
						} else if b[110+i] == 52 {
							closepx += 4 * powb(10, 9-n)
						} else if b[110+i] == 53 {
							closepx += 5 * powb(10, 9-n)
						} else if b[110+i] == 54 {
							closepx += 6 * powb(10, 9-n)
						} else if b[110+i] == 55 {
							closepx += 7 * powb(10, 9-n)
						} else if b[110+i] == 56 {
							closepx += 8 * powb(10, 9-n)
						} else if b[110+i] == 57 {
							closepx += 9 * powb(10, 9-n)
						} else {
						}
					}

					var bp1 int64
					bp1 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[122+i] == 49 {
							bp1 += powb(10, 9-n)
						} else if b[122+i] == 50 {
							bp1 += 2 * powb(10, 9-n)
						} else if b[122+i] == 51 {
							bp1 += 3 * powb(10, 9-n)
						} else if b[122+i] == 52 {
							bp1 += 4 * powb(10, 9-n)
						} else if b[122+i] == 53 {
							bp1 += 5 * powb(10, 9-n)
						} else if b[122+i] == 54 {
							bp1 += 6 * powb(10, 9-n)
						} else if b[122+i] == 55 {
							bp1 += 7 * powb(10, 9-n)
						} else if b[122+i] == 56 {
							bp1 += 8 * powb(10, 9-n)
						} else if b[122+i] == 57 {
							bp1 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var bv1 int64
					bv1 = 0
					for i := 0; i < 12; i++ {
						if b[134+i] == 49 {
							bv1 += powb(10, 11-i)
						} else if b[134+i] == 50 {
							bv1 += 2 * powb(10, 11-i)
						} else if b[134+i] == 51 {
							bv1 += 3 * powb(10, 11-i)
						} else if b[134+i] == 52 {
							bv1 += 4 * powb(10, 11-i)
						} else if b[134+i] == 53 {
							bv1 += 5 * powb(10, 11-i)
						} else if b[134+i] == 54 {
							bv1 += 6 * powb(10, 11-i)
						} else if b[134+i] == 55 {
							bv1 += 7 * powb(10, 11-i)
						} else if b[134+i] == 56 {
							bv1 += 8 * powb(10, 11-i)
						} else if b[134+i] == 57 {
							bv1 += 9 * powb(10, 11-i)
						} else {
						}
					}

					var sp1 int64
					sp1 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[147+i] == 49 {
							sp1 += powb(10, 9-n)
						} else if b[147+i] == 50 {
							sp1 += 2 * powb(10, 9-n)
						} else if b[147+i] == 51 {
							sp1 += 3 * powb(10, 9-n)
						} else if b[147+i] == 52 {
							sp1 += 4 * powb(10, 9-n)
						} else if b[147+i] == 53 {
							sp1 += 5 * powb(10, 9-n)
						} else if b[147+i] == 54 {
							sp1 += 6 * powb(10, 9-n)
						} else if b[147+i] == 55 {
							sp1 += 7 * powb(10, 9-n)
						} else if b[147+i] == 56 {
							sp1 += 8 * powb(10, 9-n)
						} else if b[147+i] == 57 {
							sp1 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var sv1 int64
					sv1 = 0
					for i := 0; i < 12; i++ {
						if b[159+i] == 49 {
							sv1 += powb(10, 11-i)
						} else if b[159+i] == 50 {
							sv1 += 2 * powb(10, 11-i)
						} else if b[159+i] == 51 {
							sv1 += 3 * powb(10, 11-i)
						} else if b[159+i] == 52 {
							sv1 += 4 * powb(10, 11-i)
						} else if b[159+i] == 53 {
							sv1 += 5 * powb(10, 11-i)
						} else if b[159+i] == 54 {
							sv1 += 6 * powb(10, 11-i)
						} else if b[159+i] == 55 {
							sv1 += 7 * powb(10, 11-i)
						} else if b[159+i] == 56 {
							sv1 += 8 * powb(10, 11-i)
						} else if b[159+i] == 57 {
							sv1 += 9 * powb(10, 11-i)
						} else {
						}
					}

					var bp2 int64
					bp2 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[172+i] == 49 {
							bp2 += powb(10, 9-n)
						} else if b[172+i] == 50 {
							bp2 += 2 * powb(10, 9-n)
						} else if b[172+i] == 51 {
							bp2 += 3 * powb(10, 9-n)
						} else if b[172+i] == 52 {
							bp2 += 4 * powb(10, 9-n)
						} else if b[172+i] == 53 {
							bp2 += 5 * powb(10, 9-n)
						} else if b[172+i] == 54 {
							bp2 += 6 * powb(10, 9-n)
						} else if b[172+i] == 55 {
							bp2 += 7 * powb(10, 9-n)
						} else if b[172+i] == 56 {
							bp2 += 8 * powb(10, 9-n)
						} else if b[172+i] == 57 {
							bp2 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var bv2 int64
					bv2 = 0
					for i := 0; i < 12; i++ {
						if b[184+i] == 49 {
							bv2 += powb(10, 11-i)
						} else if b[184+i] == 50 {
							bv2 += 2 * powb(10, 11-i)
						} else if b[184+i] == 51 {
							bv2 += 3 * powb(10, 11-i)
						} else if b[184+i] == 52 {
							bv2 += 4 * powb(10, 11-i)
						} else if b[184+i] == 53 {
							bv2 += 5 * powb(10, 11-i)
						} else if b[184+i] == 54 {
							bv2 += 6 * powb(10, 11-i)
						} else if b[184+i] == 55 {
							bv2 += 7 * powb(10, 11-i)
						} else if b[184+i] == 56 {
							bv2 += 8 * powb(10, 11-i)
						} else if b[184+i] == 57 {
							bv2 += 9 * powb(10, 11-i)
						} else {
						}
					}

					var sp2 int64
					sp2 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[197+i] == 49 {
							sp2 += powb(10, 9-n)
						} else if b[197+i] == 50 {
							sp2 += 2 * powb(10, 9-n)
						} else if b[197+i] == 51 {
							sp2 += 3 * powb(10, 9-n)
						} else if b[197+i] == 52 {
							sp2 += 4 * powb(10, 9-n)
						} else if b[197+i] == 53 {
							sp2 += 5 * powb(10, 9-n)
						} else if b[197+i] == 54 {
							sp2 += 6 * powb(10, 9-n)
						} else if b[197+i] == 55 {
							sp2 += 7 * powb(10, 9-n)
						} else if b[197+i] == 56 {
							sp2 += 8 * powb(10, 9-n)
						} else if b[197+i] == 57 {
							sp2 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var sv2 int64
					sv2 = 0
					for i := 0; i < 12; i++ {
						if b[209+i] == 49 {
							sv2 += powb(10, 11-i)
						} else if b[209+i] == 50 {
							sv2 += 2 * powb(10, 11-i)
						} else if b[209+i] == 51 {
							sv2 += 3 * powb(10, 11-i)
						} else if b[209+i] == 52 {
							sv2 += 4 * powb(10, 11-i)
						} else if b[209+i] == 53 {
							sv2 += 5 * powb(10, 11-i)
						} else if b[209+i] == 54 {
							sv2 += 6 * powb(10, 11-i)
						} else if b[209+i] == 55 {
							sv2 += 7 * powb(10, 11-i)
						} else if b[209+i] == 56 {
							sv2 += 8 * powb(10, 11-i)
						} else if b[209+i] == 57 {
							sv2 += 9 * powb(10, 11-i)
						} else {
						}
					}

					var bp3 int64
					bp3 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[222+i] == 49 {
							bp3 += powb(10, 9-n)
						} else if b[222+i] == 50 {
							bp3 += 2 * powb(10, 9-n)
						} else if b[222+i] == 51 {
							bp3 += 3 * powb(10, 9-n)
						} else if b[222+i] == 52 {
							bp3 += 4 * powb(10, 9-n)
						} else if b[222+i] == 53 {
							bp3 += 5 * powb(10, 9-n)
						} else if b[222+i] == 54 {
							bp3 += 6 * powb(10, 9-n)
						} else if b[222+i] == 55 {
							bp3 += 7 * powb(10, 9-n)
						} else if b[222+i] == 56 {
							bp3 += 8 * powb(10, 9-n)
						} else if b[222+i] == 57 {
							bp3 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var bv3 int64
					bv3 = 0
					for i := 0; i < 12; i++ {
						if b[234+i] == 49 {
							bv3 += powb(10, 11-i)
						} else if b[234+i] == 50 {
							bv3 += 2 * powb(10, 11-i)
						} else if b[234+i] == 51 {
							bv3 += 3 * powb(10, 11-i)
						} else if b[234+i] == 52 {
							bv3 += 4 * powb(10, 11-i)
						} else if b[234+i] == 53 {
							bv3 += 5 * powb(10, 11-i)
						} else if b[234+i] == 54 {
							bv3 += 6 * powb(10, 11-i)
						} else if b[234+i] == 55 {
							bv3 += 7 * powb(10, 11-i)
						} else if b[234+i] == 56 {
							bv3 += 8 * powb(10, 11-i)
						} else if b[234+i] == 57 {
							bv3 += 9 * powb(10, 11-i)
						} else {
						}
					}

					var sp3 int64
					sp3 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[247+i] == 49 {
							sp3 += powb(10, 9-n)
						} else if b[247+i] == 50 {
							sp3 += 2 * powb(10, 9-n)
						} else if b[247+i] == 51 {
							sp3 += 3 * powb(10, 9-n)
						} else if b[247+i] == 52 {
							sp3 += 4 * powb(10, 9-n)
						} else if b[247+i] == 53 {
							sp3 += 5 * powb(10, 9-n)
						} else if b[247+i] == 54 {
							sp3 += 6 * powb(10, 9-n)
						} else if b[247+i] == 55 {
							sp3 += 7 * powb(10, 9-n)
						} else if b[247+i] == 56 {
							sp3 += 8 * powb(10, 9-n)
						} else if b[247+i] == 57 {
							sp3 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var sv3 int64
					sv3 = 0
					for i := 0; i < 12; i++ {
						if b[259+i] == 49 {
							sv3 += powb(10, 11-i)
						} else if b[259+i] == 50 {
							sv3 += 2 * powb(10, 11-i)
						} else if b[259+i] == 51 {
							sv3 += 3 * powb(10, 11-i)
						} else if b[259+i] == 52 {
							sv3 += 4 * powb(10, 11-i)
						} else if b[259+i] == 53 {
							sv3 += 5 * powb(10, 11-i)
						} else if b[259+i] == 54 {
							sv3 += 6 * powb(10, 11-i)
						} else if b[259+i] == 55 {
							sv3 += 7 * powb(10, 11-i)
						} else if b[259+i] == 56 {
							sv3 += 8 * powb(10, 11-i)
						} else if b[259+i] == 57 {
							sv3 += 9 * powb(10, 11-i)
						} else {
						}
					}

					var bp4 int64
					bp4 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[272+i] == 49 {
							bp4 += powb(10, 9-n)
						} else if b[272+i] == 50 {
							bp4 += 2 * powb(10, 9-n)
						} else if b[272+i] == 51 {
							bp4 += 3 * powb(10, 9-n)
						} else if b[272+i] == 52 {
							bp4 += 4 * powb(10, 9-n)
						} else if b[272+i] == 53 {
							bp4 += 5 * powb(10, 9-n)
						} else if b[272+i] == 54 {
							bp4 += 6 * powb(10, 9-n)
						} else if b[272+i] == 55 {
							bp4 += 7 * powb(10, 9-n)
						} else if b[272+i] == 56 {
							bp4 += 8 * powb(10, 9-n)
						} else if b[272+i] == 57 {
							bp4 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var bv4 int64
					bv4 = 0
					for i := 0; i < 12; i++ {
						if b[284+i] == 49 {
							bv4 += powb(10, 11-i)
						} else if b[284+i] == 50 {
							bv4 += 2 * powb(10, 11-i)
						} else if b[284+i] == 51 {
							bv4 += 3 * powb(10, 11-i)
						} else if b[284+i] == 52 {
							bv4 += 4 * powb(10, 11-i)
						} else if b[284+i] == 53 {
							bv4 += 5 * powb(10, 11-i)
						} else if b[284+i] == 54 {
							bv4 += 6 * powb(10, 11-i)
						} else if b[284+i] == 55 {
							bv4 += 7 * powb(10, 11-i)
						} else if b[284+i] == 56 {
							bv4 += 8 * powb(10, 11-i)
						} else if b[284+i] == 57 {
							bv4 += 9 * powb(10, 11-i)
						} else {
						}
					}

					var sp4 int64
					sp4 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[297+i] == 49 {
							sp4 += powb(10, 9-n)
						} else if b[297+i] == 50 {
							sp4 += 2 * powb(10, 9-n)
						} else if b[297+i] == 51 {
							sp4 += 3 * powb(10, 9-n)
						} else if b[297+i] == 52 {
							sp4 += 4 * powb(10, 9-n)
						} else if b[297+i] == 53 {
							sp4 += 5 * powb(10, 9-n)
						} else if b[297+i] == 54 {
							sp4 += 6 * powb(10, 9-n)
						} else if b[297+i] == 55 {
							sp4 += 7 * powb(10, 9-n)
						} else if b[297+i] == 56 {
							sp4 += 8 * powb(10, 9-n)
						} else if b[297+i] == 57 {
							sp4 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var sv4 int64
					sv4 = 0
					for i := 0; i < 12; i++ {
						if b[309+i] == 49 {
							sv4 += powb(10, 11-i)
						} else if b[309+i] == 50 {
							sv4 += 2 * powb(10, 11-i)
						} else if b[309+i] == 51 {
							sv4 += 3 * powb(10, 11-i)
						} else if b[309+i] == 52 {
							sv4 += 4 * powb(10, 11-i)
						} else if b[309+i] == 53 {
							sv4 += 5 * powb(10, 11-i)
						} else if b[309+i] == 54 {
							sv4 += 6 * powb(10, 11-i)
						} else if b[309+i] == 55 {
							sv4 += 7 * powb(10, 11-i)
						} else if b[309+i] == 56 {
							sv4 += 8 * powb(10, 11-i)
						} else if b[309+i] == 57 {
							sv4 += 9 * powb(10, 11-i)
						} else {
						}
					}
					var bp5 int64
					bp5 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[322+i] == 49 {
							bp5 += powb(10, 9-n)
						} else if b[322+i] == 50 {
							bp5 += 2 * powb(10, 9-n)
						} else if b[322+i] == 51 {
							bp5 += 3 * powb(10, 9-n)
						} else if b[322+i] == 52 {
							bp5 += 4 * powb(10, 9-n)
						} else if b[322+i] == 53 {
							bp5 += 5 * powb(10, 9-n)
						} else if b[322+i] == 54 {
							bp5 += 6 * powb(10, 9-n)
						} else if b[322+i] == 55 {
							bp5 += 7 * powb(10, 9-n)
						} else if b[322+i] == 56 {
							bp5 += 8 * powb(10, 9-n)
						} else if b[322+i] == 57 {
							bp5 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var bv5 int64
					bv5 = 0
					for i := 0; i < 12; i++ {
						if b[334+i] == 49 {
							bv5 += powb(10, 11-i)
						} else if b[334+i] == 50 {
							bv5 += 2 * powb(10, 11-i)
						} else if b[334+i] == 51 {
							bv5 += 3 * powb(10, 11-i)
						} else if b[334+i] == 52 {
							bv5 += 4 * powb(10, 11-i)
						} else if b[334+i] == 53 {
							bv5 += 5 * powb(10, 11-i)
						} else if b[334+i] == 54 {
							bv5 += 6 * powb(10, 11-i)
						} else if b[334+i] == 55 {
							bv5 += 7 * powb(10, 11-i)
						} else if b[334+i] == 56 {
							bv5 += 8 * powb(10, 11-i)
						} else if b[334+i] == 57 {
							bv5 += 9 * powb(10, 11-i)
						} else {
						}
					}

					var sp5 int64
					sp5 = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[347+i] == 49 {
							sp5 += powb(10, 9-n)
						} else if b[347+i] == 50 {
							sp5 += 2 * powb(10, 9-n)
						} else if b[347+i] == 51 {
							sp5 += 3 * powb(10, 9-n)
						} else if b[347+i] == 52 {
							sp5 += 4 * powb(10, 9-n)
						} else if b[347+i] == 53 {
							sp5 += 5 * powb(10, 9-n)
						} else if b[347+i] == 54 {
							sp5 += 6 * powb(10, 9-n)
						} else if b[347+i] == 55 {
							sp5 += 7 * powb(10, 9-n)
						} else if b[347+i] == 56 {
							sp5 += 8 * powb(10, 9-n)
						} else if b[347+i] == 57 {
							sp5 += 9 * powb(10, 9-n)
						} else {
						}
					}

					var sv5 int64
					sv5 = 0
					for i := 0; i < 12; i++ {
						if b[359+i] == 49 {
							sv5 += powb(10, 11-i)
						} else if b[359+i] == 50 {
							sv5 += 2 * powb(10, 11-i)
						} else if b[359+i] == 51 {
							sv5 += 3 * powb(10, 11-i)
						} else if b[359+i] == 52 {
							sv5 += 4 * powb(10, 11-i)
						} else if b[359+i] == 53 {
							sv5 += 5 * powb(10, 11-i)
						} else if b[359+i] == 54 {
							sv5 += 6 * powb(10, 11-i)
						} else if b[359+i] == 55 {
							sv5 += 7 * powb(10, 11-i)
						} else if b[359+i] == 56 {
							sv5 += 8 * powb(10, 11-i)
						} else if b[359+i] == 57 {
							sv5 += 9 * powb(10, 11-i)
						} else {
						}
					}

					var precloseiopv int64
					precloseiopv = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[372+i] == 49 {
							precloseiopv += powb(10, 9-n)
						} else if b[372+i] == 50 {
							precloseiopv += 2 * powb(10, 9-n)
						} else if b[372+i] == 51 {
							precloseiopv += 3 * powb(10, 9-n)
						} else if b[372+i] == 52 {
							precloseiopv += 4 * powb(10, 9-n)
						} else if b[372+i] == 53 {
							precloseiopv += 5 * powb(10, 9-n)
						} else if b[372+i] == 54 {
							precloseiopv += 6 * powb(10, 9-n)
						} else if b[372+i] == 55 {
							precloseiopv += 7 * powb(10, 9-n)
						} else if b[372+i] == 56 {
							precloseiopv += 8 * powb(10, 9-n)
						} else if b[372+i] == 57 {
							precloseiopv += 9 * powb(10, 9-n)
						} else {
						}
					}

					var iopv int64
					iopv = 0
					for i := 0; i < 11; i++ {
						n := 0
						if i == 7 {
							continue
						} else if i > 7 {
							n = i - 1
						} else {
							n = i
						}

						if b[384+i] == 49 {
							iopv += powb(10, 9-n)
						} else if b[384+i] == 50 {
							iopv += 2 * powb(10, 9-n)
						} else if b[384+i] == 51 {
							iopv += 3 * powb(10, 9-n)
						} else if b[384+i] == 52 {
							iopv += 4 * powb(10, 9-n)
						} else if b[384+i] == 53 {
							iopv += 5 * powb(10, 9-n)
						} else if b[384+i] == 54 {
							iopv += 6 * powb(10, 9-n)
						} else if b[384+i] == 55 {
							iopv += 7 * powb(10, 9-n)
						} else if b[384+i] == 56 {
							iopv += 8 * powb(10, 9-n)
						} else if b[384+i] == 57 {
							iopv += 9 * powb(10, 9-n)
						} else {
						}
					}

					status1 := b[396]

					var status2 int64
					status2 = 0
					for i := 0; i < 3; i++ {
						if b[397+i] == 49 {
							status2 += powb(10, 2-i)
						} else {
						}
					}
					timestamp := make([]byte, 9)
					//copy(timestamp, b[405:417])
					copy(timestamp[0:2], b[405:407])
					copy(timestamp[2:4], b[408:410])
					copy(timestamp[4:6], b[411:413])
					copy(timestamp[6:9], b[414:417])

					debug := false
					temp := string(timestamp)
					if debug {
						fmt.Printf("\n---[code|%v] [volume|%d] [amount|%d] [lastprice|%d]", code, volume, amount, lastprice)
						fmt.Printf(" [open|%d] [high|%d] [low|%d] [tradeprice|%d] [closepx|%d]", open, high, low, tradeprice, closepx)
						fmt.Printf(" [bp1|%d] [bv1|%d] [sp1|%d] [sv1|%d]", bp1, bv1, sp1, sv1)
						fmt.Printf(" [bp2|%d] [bv2|%d] [sp2|%d] [sv2|%d]", bp2, bv2, sp2, sv2)
						fmt.Printf(" [bp3|%d] [bv3|%d] [sp3|%d] [sv3|%d]", bp3, bv3, sp3, sv3)
						fmt.Printf(" [bp4|%d] [bv4|%d] [sp4|%d] [sv4|%d]", bp4, bv4, sp4, sv4)
						fmt.Printf(" [bp5|%d] [bv5|%d] [sp5|%d] [sv5|%d]", bp5, bv5, sp5, sv5)
						fmt.Printf(" [precloseiopv|%d] [iopv|%d]", precloseiopv, iopv)
						fmt.Printf(" [status1|%v] [status2|%d]", status1, status2)
						fmt.Printf(" [timestamp|%s]", temp)
						fmt.Println("---")
					}

					//msgpack序列化
					_code := string(code) + ".SH"

					timestamp_buf, err := strconv.ParseInt(string(timestamp), 10, 32)
					if err != nil {
						log.Error("string to int32 err: ", err)
					}
					_time := int32(timestamp_buf)

					var _status string
					if status1 == 0x43 { //C => I
						_status = "I"
					} else if status1 == 0x54 { //T => O
						_status = "O"
					} else if status1 == 0x50 { //P => B
						_status = "B"
					}

					_preClose := float64(lastprice) / 1000.0
					_open := float64(open) / 1000.0
					_high := float64(high) / 1000.0
					_low := float64(low) / 1000.0
					_match := float64(tradeprice) / 1000.0
					var _askPrice = [10]float64{float64(sp1) / 1000.0, float64(sp2) / 1000.0, float64(sp3) / 1000.0, float64(sp4) / 1000.0, float64(sp5) / 1000.0}
					var _askVol = [10]int32{int32(sv1), int32(sv2), int32(sv3), int32(sv4), int32(sv5)}
					var _bidPrice = [10]float64{float64(bp1) / 1000.0, float64(bp2) / 1000.0, float64(bp3) / 1000.0, float64(bp4) / 1000.0, float64(bp5) / 1000.0}
					var _bidVol = [10]int32{int32(bv1), int32(bv2), int32(bv3), int32(bv4), int32(bv5)}
					_IOPV := float64(iopv) / 1000.0
					_turnover := amount / 100

					debug = false
					if debug {
						fmt.Printf("[code|%v] [time|%d] [status|%v] [preClose|%f]", _code, _time, _status, _preClose)
						fmt.Printf(" [open|%f] [high|%f] [low|%f] [match|%f]", _open, _high, _low, _match)
						fmt.Printf(" [askPrice|%v] [askVol|%v] [bidPrice|%v] [bidVol|%v]", _askPrice, _askVol, _bidPrice, _bidVol)
						fmt.Printf(" [volume|%d] [turnover|%d] [IOPV|%f]\n", volume, _turnover, _IOPV)
					}

					marketData := hqbase.Marketdata{Code: _code, Time: _time, Status: _status, PreClose: _preClose, Open: _open, High: _high, Low: _low, Match: _match, AskPrice: _askPrice, AskVol: _askVol, BidPrice: _bidPrice, BidVol: _bidVol, Volume: volume, Turnover: _turnover, IOPV: _IOPV}

					data, err := msgpack.Marshal(&marketData)
					if err != nil {
						panic(err)
					}

					debug = false
					if debug {
						fmt.Printf("data|%v\n", data)
						var mData hqbase.Marketdata
						err = msgpack.Unmarshal(data, &mData)
						if err != nil {
							panic(err)
						}
						fmt.Printf("Unmarshal:[code|%v] [time|%d] [status|%v] [preClose|%f]", mData.Code, mData.Time, mData.Status, mData.PreClose)
						fmt.Printf(" [open|%f] [high|%f] [low|%f] [match|%f]", mData.Open, mData.High, mData.Low, mData.Match)
						fmt.Printf(" [askPrice|%v] [askVol|%v] [bidPrice|%v] [bidVol|%v]", mData.AskPrice, mData.AskVol, mData.BidPrice, mData.BidVol)
						fmt.Printf(" [volume|%d] [turnover|%d] [IOPV|%f]\n", mData.Volume, mData.Turnover, mData.IOPV)
					}

					if count == 10 {
						ct2 := time.Now().Sub(tt1)
						fmt.Printf("%d %d %v\n", seq, count, ct2)
					}
					redisRb.Put(data)
					rb.Put(data)
					// mysqlRb.Put(marketData)
				} else {
					time.Sleep(time.Duration(1) * time.Millisecond)
				}

			}

		}(i)
	}
}

func powb(x, n int) int64 {
	if n == 0 {
		return 1
	}
	for {
		if (n & 1) != 0 {
			break
		}
		n >>= 1
		x *= x
	}
	result := x
	n >>= 1
	for {
		if n == 0 {
			break
		}
		x *= x
		if (n & 1) != 0 {
			result *= x
		}
		n >>= 1
	}
	return int64(result)
}
