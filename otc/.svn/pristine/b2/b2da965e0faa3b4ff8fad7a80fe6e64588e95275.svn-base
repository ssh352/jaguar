package vss

import (
	"fmt"
	"net"
	"strings"
	"time"

	"quant/hqmodule/hqbase"
	"util/redis"

	"quant/hqmodule/marketdata/vss/struc"

	"github.com/Workiva/go-datastructures/queue"
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"github.com/widuu/goini"
)

var (
	RedisPool *redis.ConnPool
	conn      *net.TCPConn
)

func InitVss(conf *goini.Config) {
	header := struc.NewHeader(1, 92)
	// 初始化vss登录信息
	var logonInfo = map[string]string{
		"SenderCompID":     "CS63                ",
		"TargetCompID":     "GFZB                ",
		"Password":         "GFZB            ",
		"DefaultApplVerID": "1.02                            ",
		// "SenderCompID":     conf.GetValue("hqmodule", "sz_vss_sendercompid"),
		// "TargetCompID":     conf.GetValue("hqmodule", "sz_vss_targetcompid"),
		// "Password":         conf.GetValue("hqmodule", "sz_vss_password"),
		// "DefaultApplVerID": conf.GetValue("hqmodule", "sz_vss_defaultapplverid"),
	}
	fmt.Printf("%v\n", logonInfo)
	logon := struc.NewLogon(logonInfo)
	// 初始化vss服务连接信息
	var CONN = map[string]string{
		"host":     conf.GetValue("hqmodule", "sz_vss_vsshost"),
		"protocol": conf.GetValue("hqmodule", "sz_vss_conntype"),
	}
	fmt.Printf("%v\n", CONN)
	// 获取vss服务连接
	var err error
	conn, err = login(logon, header, CONN)
	checkError(err)
}

// 心跳
func Heartbeat() {
	var err error
	header := struc.NewHeader(3, 0)
	headerlen := 8
	tailerlen := 4
	heartbeatbodylen := 0
	heartbeatmsglen := headerlen + tailerlen + heartbeatbodylen
	sendmsg := make([]byte, heartbeatmsglen)
	copy(sendmsg, header.Marshal())
	copy(sendmsg[8:], ck(sendmsg[:8]))
	log.Info("heartbeatmsg %v\n", sendmsg)
	for {
		time.Sleep(time.Second * time.Duration(60))
		_, err = conn.Write(sendmsg) //发送HTTP请求头
		if err != nil {
			log.Error("%v\n", err)
		}
	}

}

// 接收数据，并处理
func ReceiveMarketData(redisRb *queue.RingBuffer, rb *queue.RingBuffer, mysqlRb *queue.RingBuffer) {
	headerlen := uint(8) // 消息头长度
	tailerlen := uint(4) // 消息尾长度
	for {
		offset := uint(0)
		headermsg := make([]byte, headerlen)
		for {
			if offset < headerlen {
				rl, _ := conn.Read(headermsg[offset:])
				offset += uint(rl)
			} else {
				break
			}
		}

		// 行情类别
		msgtype := bytestoint(headermsg[:4])
		bodylength := bytestoint(headermsg[4:])
		bodymsg := make([]byte, bodylength)
		offset = 0
		for {
			if offset < bodylength {
				rl, _ := conn.Read(bodymsg[offset:])
				offset += uint(rl)
			} else {
				break
			}
		}

		tailmsg := make([]byte, tailerlen)
		offset = 0
		for {
			if offset < tailerlen {
				rl, _ := conn.Read(tailmsg[offset:])
				offset += uint(rl)
			} else {
				break
			}
		}
		headerandbody := make([]byte, headerlen+bodylength)
		copy(headerandbody, headermsg)
		copy(headerandbody[headerlen:], bodymsg)
		cksum := ck(headerandbody)
		for i := uint(0); i < tailerlen; i++ {
			if tailmsg[i] != cksum[i] {
				log.Error("cksum error!!!!!\n")
				break
			}
		}

		//	log.Info("body msg:%v\n", bodymsg)
		/* 消息总长度：行情类别（1） + 消息内容长度 + 消息体结尾：127（1） + 记录末尾：21（1）
		 * 行情类别：11-深圳股票和ETF(300111)；12-深圳指数(309011)；13-深圳盘后定价信息(300611)
		 */
		bodyTotalLen := 1 + len(bodymsg) + 1 + 1
		mk := make([]byte, bodyTotalLen)
		//log.Info("bodymsg:%v\n", bodymsg)

		switch msgtype {
		case 390019:
			//log.Info("390019 市场实时动态\n")
		case 390013:
			//log.Info("390013 证券实时动态\n")
		case 390012:
			//log.Info("390012 公告\n")
		case 300111:
			//log.Info("300111 集中竞价交易业务行情快照\n")
			//	securityid := bodymsg[13:21]
			//	log.Info("!!!!!!!证券代码:%s %v\n", string(securityid), securityid)
			// 设置行情类别
			mk[0] = 0xb
			// 设置消息体
			copy(mk[1:1+len(bodymsg)], bodymsg)
			// 设置消息结尾标志(以127结尾)
			mk[bodyTotalLen-1-1] = 0x7f
			// 设置行情记录结束标志（以21结尾）
			mk[bodyTotalLen-1] = 0x15
			//	securityid1 := mk[14:22]
			//	log.Info("!!!!!!!证券代码1:%s %v\n", string(securityid1), securityid1)

			mkd := parseMarketData(mk)
			md, err := msgpack.Marshal(&mkd)
			if err != nil {
				panic(err)
			}
			redisRb.Put(md)
			rb.Put(md)
			mysqlRb.Put(mkd)
		case 300611:
			//log.Info("300611 盘后定价交易业务行情快照\n")
			//	securityid := bodymsg[13:21]
			//	log.Info("!!!!!!!证券代码:%s %v\n", string(securityid), securityid)
			// 设置行情类别
			mk[0] = 0xd
			// 设置消息体
			copy(mk[1:1+len(bodymsg)], bodymsg)
			// 设置消息结尾标志(以127结尾)
			mk[bodyTotalLen-1-1] = 0x7f
			// 设置行情记录结束标志（以21结尾）
			mk[bodyTotalLen-1] = 0x15
			//	securityid1 := mk[14:22]
			//	log.Info("!!!!!!!证券代码1:%s %v\n", string(securityid1), securityid1)

			mkd := parseMarketData(mk)
			md, err := msgpack.Marshal(&mkd)
			if err != nil {
				panic(err)
			}
			redisRb.Put(md)
			rb.Put(md)
			mysqlRb.Put(mkd)
		case 306311:
			//log.Info("306311 港股实时行情快照\n")
		case 309011:
			//log.Info("309011 指数行情快照\n")
			//	securityid := bodymsg[13:21]
			//	log.Info("!!!!!!!证券代码:%s %v\n", string(securityid), securityid)
			// 设置行情类别
			mk[0] = 0xc
			// 设置消息体
			copy(mk[1:1+len(bodymsg)], bodymsg)
			// 设置消息结尾标志(以127结尾)
			mk[bodyTotalLen-1-1] = 0x7f
			// 设置行情记录结束标志（以21结尾）
			mk[bodyTotalLen-1] = 0x15
			//		securityid1 := mk[14:22]
			//	log.Info("!!!!!!!证券代码1:%s %v\n", string(securityid1), securityid1)

			mkd := parseMarketData(mk)
			md, err := msgpack.Marshal(&mkd)
			if err != nil {
				panic(err)
			}
			redisRb.Put(md)
			rb.Put(md)
			mysqlRb.Put(mkd)
		case 309111:
			//log.Info("309111 成交量统计指标行情\n")
		case 300192:
			//log.Info("300192 集中竞价业务逐笔委托行情\n")
		case 300592:
			//log.Info("300592 协议交易业务逐笔委托行情\n")
		case 300792:
			//log.Info("300792 转融通证券出借业务逐笔委托行情\n")
		case 300191:
			//log.Info("300192 集中竞价业务逐笔成交行情\n")
		case 300591:
			//log.Info("300592 协议交易业务逐笔成交行情\n")
		case 300791:
			//log.Info("300792 转融通证券出借业务逐笔成交行情\n")
		default:
			//log.Info("%d wtf!!!\n", msgtype)
		}
	}

}

// 登录vss服务
func login(logon *struc.Logon, header *struc.Header, CONN map[string]string) (*net.TCPConn, error) {
	headerlen := 8
	tailerlen := 4
	logonbodylen := 92
	logonmsglen := headerlen + tailerlen + logonbodylen
	sendmsg := make([]byte, logonmsglen)
	copy(sendmsg, header.Marshal())
	copy(sendmsg[8:], logon.Marshal())
	copy(sendmsg[100:], ck(sendmsg[:100]))

	log.Info("host:%s, proto:%s\n", CONN["host"], CONN["protocol"])
	tcpAddr, err := net.ResolveTCPAddr(CONN["protocol"], CONN["host"]) //获取一个TCP地址信息,TCPAddr
	if err != nil {
		log.Info("Fatal error: %s\n", err.Error())
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr) //创建一个TCP连接:TCPConn
	if err != nil {
		log.Info("Fatal error: %s\n", err.Error())
		return nil, err
	}
	_, err = conn.Write(sendmsg) //发送HTTP请求头
	if err != nil {
		log.Info("Fatal error: %s\n", err.Error())
		return nil, err
	}
	result := make([]byte, logonmsglen)
	llen, err := conn.Read(result)
	if err != nil {
		log.Info("3 登陆网关失败 %d\n", llen)
		return nil, err
	}
	log.Info("llen %d\n", llen)
	log.Info("%v\n", result)
	if llen > 0 {
		if result[3] == 1 {
			log.Info("登陆网关成功\n")
		} else {
			log.Info("1 登陆网关失败 %s %v\n", result[3], result[3])
		}
	} else {
		log.Info("2 登陆网关失败 %d %v\n", result[3], result[3])
	}
	return conn, nil
}

// 查询redis行情，并解析
func queryRedis(codes string) {
	codelist := strings.Split(codes, ",")
	md, err := RedisPool.GetMHashMapString("MarketMap_test", codelist)
	//	md, err := RedisPool.GetAllHashMapString("MarketMap_test")
	if err != nil {
		log.Error("QueryStock ", err)
	}
	//	log.Info("md:%v\n", md)
	var mkd hqbase.Marketdata
	for i := 0; i < len(md); i++ {
		log.Info("反序列化redis数据库二进制数据！")
		data := md[i]
		bodymsg := []byte(data)
		mkd = parseMarketData(bodymsg)
		b, err := msgpack.Marshal(&mkd)
		if err != nil {
			panic(err)
		}
		log.Info("=========msgpack.b:%v\n", b)
		var marketd hqbase.Marketdata
		err = msgpack.Unmarshal(b, &marketd)
		if err != nil {
			panic(err)
		}
		fmt.Printf("marketd:%+v\n", marketd)
		log.Info("mkd:%+v\n", mkd)
	}
}

func parseMarketData(bodymsg []byte) hqbase.Marketdata {
	var mkd hqbase.Marketdata
	mktype := bytestoint(bodymsg[:1])
	//log.Info("!!!!!!!行情类型：%d\n", mktype)
	timestamp := bytestoint(bodymsg[1:9])
	//log.Info("!!!!!!!时间戳:%d\n", timestamp)
	mkd.Time = int32(timestamp % 1000000000)
	//channelno := bytestoint(bodymsg[9:11])
	//log.Info("!!!!!!!频道代码:%d\n", channelno)
	//mdstreamid := bodymsg[11:14]
	//log.Info("!!!!!!!行情类别:%s %v\n", string(mdstreamid), mdstreamid)
	var securityid []byte
	if bodymsg[20] != 32 && bodymsg[21] != 32 {
		securityid = bodymsg[14:22]
	} else if bodymsg[20] != 32 && bodymsg[21] == 32 {
		securityid = bodymsg[14:21]
	} else {
		securityid = bodymsg[14:20]
	}
	//log.Info("!!!!!!!证券代码:%s %v\n", string(securityid), securityid)
	code := string(securityid) + ".SZ"
	mkd.Code = code
	//securityidsource := bodymsg[22:26]
	//log.Info("!!!!!!!证券代码源:%s %v\n", string(securityidsource), securityidsource)
	//tradingPhaseCode := bodymsg[26:34]
	//log.Info("!!!!!!!产品所处的交易阶段代码:%s %v\n", string(tradingPhaseCode), tradingPhaseCode)
	switch string(bodymsg[26:27]) {
	case "S":
		// 开市前
		mkd.Status = "0"
	case "O":
		// 开盘集合竞价
		mkd.Status = "I"
	case "T":
		// 连续竞价阶段
		mkd.Status = "O"
	case "H":
		// 盘中临时停市
		mkd.Status = "W"
	case "C":
		// 收盘集合竞价
		mkd.Status = "J"
	case "E":
		// 闭市
		mkd.Status = "C"
	}
	preClosePx := bytestoint(bodymsg[34:42])
	//log.Info("!!!!!!!昨收盘:%d\n", preClosePx)
	mkd.PreClose = float64(preClosePx) / 10000.0
	numTrades := bytestoint(bodymsg[42:50])
	//log.Info("!!!!!!!成交笔数:%d\n", numTrades)
	mkd.NumTrades = int32(numTrades)
	totalVolumeTrade := bytestoint(bodymsg[50:58])
	//log.Info("!!!!!!!成交总量:%d\n", totalVolumeTrade)
	mkd.Volume = int64(totalVolumeTrade)
	totalValueTrade := bytestoint(bodymsg[58:66])
	//log.Info("!!!!!!!成交总金额:%d %v\n", totalValueTrade, bodymsg[58:66])
	mkd.Turnover = int64(totalValueTrade) / 10000
	noMdEntries := bytestoint(bodymsg[66:70])
	//log.Info("!!!!!!!行情条目个数:%d\n", noMdEntries)
	// 深圳股票和ETF
	if mktype == 11 {
		byteoffset := uint(0)
		for i := uint(0); i < noMdEntries; i++ {
			mDEntryType := bodymsg[70+byteoffset : 72+byteoffset]
			//log.Info("!!!!!!![%d] 行情条目类别:%s\n", i+1, string(mDEntryType))
			mDPriceLevel := bytestoint(bodymsg[88+byteoffset : 90+byteoffset])
			//log.Info("!!!!!!![%d] 买卖盘档位:%d\n", i+1, mDPriceLevel)
			mDEntryPx := bytestoint(bodymsg[72+byteoffset : 80+byteoffset])
			//log.Info("!!!!!!![%d] 价格:%d\n", i+1, mDEntryPx)
			mDEntrySize := bytestoint(bodymsg[80+byteoffset : 88+byteoffset])
			//log.Info("!!!!!!![%d] 数量:%d\n", i+1, mDEntrySize)
			switch strings.TrimSpace(string(mDEntryType)) {
			case "0":
				// 买入
				switch mDPriceLevel {
				case 1:
					// 买1
					mkd.BidPrice[0] = float64(mDEntryPx) / 1000000.0
					mkd.BidVol[0] = int32(mDEntrySize)
				case 2:
					// 买2
					mkd.BidPrice[1] = float64(mDEntryPx) / 1000000.0
					mkd.BidVol[1] = int32(mDEntrySize)
				case 3:
					// 买3
					mkd.BidPrice[2] = float64(mDEntryPx) / 1000000.0
					mkd.BidVol[2] = int32(mDEntrySize)
				case 4:
					// 买4
					mkd.BidPrice[3] = float64(mDEntryPx) / 1000000.0
					mkd.BidVol[3] = int32(mDEntrySize)
				case 5:
					// 买5
					mkd.BidPrice[4] = float64(mDEntryPx) / 1000000.0
					mkd.BidVol[4] = int32(mDEntrySize)
				}
			case "1":
				// 卖出
				switch mDPriceLevel {
				case 1:
					// 卖1
					mkd.AskPrice[0] = float64(mDEntryPx) / 1000000.0
					mkd.AskVol[0] = int32(mDEntrySize)
				case 2:
					// 卖2
					mkd.AskPrice[1] = float64(mDEntryPx) / 1000000.0
					mkd.AskVol[1] = int32(mDEntrySize)
				case 3:
					// 卖3
					mkd.AskPrice[2] = float64(mDEntryPx) / 1000000.0
					mkd.AskVol[2] = int32(mDEntrySize)
				case 4:
					// 卖4
					mkd.AskPrice[3] = float64(mDEntryPx) / 1000000.0
					mkd.AskVol[3] = int32(mDEntrySize)
				case 5:
					// 卖5
					mkd.AskPrice[4] = float64(mDEntryPx) / 1000000.0
					mkd.AskVol[4] = int32(mDEntrySize)
				}
			case "2":
				// 最近价
				mkd.Match = float64(mDEntryPx) / 1000000.0
			case "4":
				// 开盘价
				mkd.Open = float64(mDEntryPx) / 1000000.0
			case "7":
				// 最高价
				mkd.High = float64(mDEntryPx) / 1000000.0
			case "8":
				// 最低价
				mkd.Low = float64(mDEntryPx) / 1000000.0
			case "x1":
				// 升跌一
			case "x2":
				// 升跌二
			case "x3":
				// 买入汇总（总量及加权平均价）
				mkd.WeightedAvgBidPrice = float64(mDEntryPx) / 1000000.0
			case "x4":
				// 卖出汇总（总量及加权平均价）
				mkd.WeightedAvgAskPrice = float64(mDEntryPx) / 1000000.0
			case "x5":
				// 股票市盈率一
			case "x6":
				// 股票市盈率二
			case "x7":
				// 基金T-1日净值
			case "x8":
				// 基金实时参考净值（包括ETF的IOPV）
				mkd.IOPV = float64(mDEntryPx) / 1000000.0
			case "x9":
				// 权证溢价率
			case "xe":
				// 涨停价
				mkd.HighLimited = float64(mDEntryPx) / 1000000.0
			case "xf":
				// 跌停价
				mkd.LowLimited = float64(mDEntryPx) / 1000000.0
			case "xg":
				// 合约持仓量
			}

			//numberOfOrders := bytestoint(bodymsg[90+byteoffset : 98+byteoffset])
			//log.Info("!!!!!!![%d] 价位总委托笔数:%d\n", i+1, numberOfOrders)
			noOrders := bytestoint(bodymsg[98+byteoffset : 102+byteoffset])
			//log.Info("!!!!!!![%d] 价位揭示委托笔数:%d\n", i+1, noOrders)
			j := uint(0)
			for ; j < noOrders; j++ {
				//orderQty := bytestoint(bodymsg[102+byteoffset+8*j : 110+byteoffset+8*j])
				//log.Info("!!!!!!![%d] [%d] 委托数量:%d\n", i+1, j+1, orderQty)
			}
			byteoffset += j*8 + 32
		}
	}

	// 深圳指数
	if mktype == 12 {
		byteoffset := uint(0)
		for i := uint(0); i < noMdEntries; i++ {
			mDEntryType := bodymsg[70+byteoffset : 72+byteoffset]
			//log.Info("!!!!!!![%d] 行情条目类别:%s\n", i+1, string(mDEntryType))
			mDEntryPx := bytestoint(bodymsg[72+byteoffset : 80+byteoffset])
			//log.Info("!!!!!!![%d] 指数点位:%d\n", i+1, mDEntryPx)
			switch strings.TrimSpace(string(mDEntryType)) {
			case "3":
				// 当前指数
				mkd.Match = float64(mDEntryPx) / 1000000.0
			case "xa":
				// 昨日收盘指数
				mkd.PreClose = float64(mDEntryPx) / 1000000.0
			case "xb":
				// 开盘指数
				mkd.Open = float64(mDEntryPx) / 1000000.0
			case "xc":
				// 最高指数
				mkd.High = float64(mDEntryPx) / 1000000.0
			case "xd":
				// 最低指数
				mkd.Low = float64(mDEntryPx) / 1000000.0
			}
			byteoffset += 10
		}
	}

	// 深圳盘后定价信息
	if mktype == 13 {
		byteoffset := uint(0)
		for i := uint(0); i < noMdEntries; i++ {
			mDEntryType := bodymsg[70+byteoffset : 72+byteoffset]
			//log.Info("!!!!!!![%d] 行情条目类别:%s\n", i+1, string(mDEntryType))
			mDEntryPx := bytestoint(bodymsg[72+byteoffset : 80+byteoffset])
			//log.Info("!!!!!!![%d] 价格:%d\n", i+1, mDEntryPx)
			mDEntrySize := bytestoint(bodymsg[80+byteoffset : 88+byteoffset])
			//log.Info("!!!!!!![%d] 数量:%d\n", i+1, mDEntrySize)
			switch strings.TrimSpace(string(mDEntryType)) {
			case "0":
				// 买入
				mkd.BidPrice[0] = float64(mDEntryPx) / 1000000.0
				mkd.BidVol[0] = int32(mDEntrySize)

			case "1":
				// 卖出
				mkd.AskPrice[0] = float64(mDEntryPx) / 1000000.0
				mkd.AskVol[0] = int32(mDEntrySize)
			}
			byteoffset += 18
		}
	}
	return mkd
}
func ck(msg []byte) []byte {
	sum := uint(0)
	l := len(msg)
	for i := 0; i < l; i++ {
		sum += uint(msg[i])
	}
	ck := uint(sum % 256)
	data := make([]byte, 4)
	data[0] = (byte)((ck & 0xFF000000) >> 24)
	data[1] = (byte)((ck & 0x00FF0000) >> 16)
	data[2] = (byte)((ck & 0x0000FF00) >> 8)
	data[3] = (byte)(ck & 0x000000FF)
	return data
}

func bytestoint(data []byte) (x1 uint) {
	x1 = uint(0)
	l := len(data)
	for i := 0; i < l; i++ {
		if i != l-1 {
			shift := uint(8 * i)
			x1 |= uint(data[l-1-i]) & 0XFF << shift
		} else {
			x1 |= uint(data[i]) & 0XFF
		}
	}
	return
}

func checkError(err error) {
	if err != nil {
		log.Error("Fatal error: %s\n", err.Error())
		panic(err)
	}
}
