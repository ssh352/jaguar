package hqbase

// Marketdata is the quote record
type Marketdata struct {
	Code string
	Time int32 //时间(HHMMSSmmm)
	// '0' 未设置
	// 'Y' 新产品
	// 'R' 交易间，禁止任何交易活动
	// 'P' 休市，例如，午餐休市。无撮合市场内部信息披露。
	// 'B' 停牌
	// 'C' 收市
	// 'D' 停牌
	// 'Z' 产品待删除
	// 'V' 结束交易
	// 'T' 固定价格集合竞价
	// 'W' 暂停，除了自有订单和交易的查询之外，任何交易活动都被禁止。
	// 'X' 停牌,( 'X'和'W'的区别在于'X'时可以撤单)
	// 'I' 盘中集合竞价。
	// 'N' 盘中集合竞价订单溥 平衡
	// 'L' 盘中集合竞价PreOBB
	// 'I' 开市集合竞价
	// 'M' 开市集合况竞价 OBB
	// 'K' 开市集合竞价订单溥平衡(OBB)前期时段
	// 'S' 非交易服务支持
	// 'U' 盘后处理
	// 'F' 盘前处理
	// 'E' 启动
	// 'O' 连续撮合
	// 'Q' 连续交易和集合竞价的波动性中断
	Status              string      // 状态
	PreClose            float64     // 前收盘价
	Open                float64     // 开盘价
	High                float64     // 最高价
	Low                 float64     // 最低价
	Close               float64     // 昨收
	Match               float64     // 最新价
	AskPrice            [10]float64 // 申卖价
	AskVol              [10]int32   // 申卖量
	BidPrice            [10]float64 // 申买价
	BidVol              [10]int32   // 申买量
	NumTrades           int32       // 成交笔数
	Volume              int64       // 成交总量
	Turnover            int64       // 成交总金额
	TotalBidVol         int64       // 委托买入总量
	TotalAskVol         int64       // 委托卖出总量
	WeightedAvgBidPrice float64     // 加权平均委买价格
	WeightedAvgAskPrice float64     // 加权平均委卖价格
	IOPV                float64     // IOPV 净值估值
	YieldToMaturity     float64     // 到期收益率
	HighLimited         float64     // 涨停价
	LowLimited          float64     // 跌停价
}
