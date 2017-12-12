package emsbase

// RespError used by communication with adapter
type RespError struct {
	ErrorCode int
	ErrorMsg  string
}

func (e *RespError) Error() string {
	return e.ErrorMsg
}

// LoginResp is the response when login in trade adapter
type LoginResp struct {
	UserToken string
	VersionNo string
	*RespError
}

// PushData is the data structure between oms and ems
type PushData struct {
	MsgType string
	Port    Portfolio
	Entrust EntrustPushResp
	Trade   DealPushResp
}

// StrategyEntrust is the entrust with strategy information
type StrategyEntrust struct {
	*StrategyInfo
	*ProductInfo
	Entrust EntrustPushResp
}

// EntrustPushResp is the entrust push data schema
type EntrustPushResp struct {
	OperatorNo       string  //操作员
	AccountCode      string  //账户编号
	BatchNo          int     //委托批号
	BusinessDate     int     //委托日期
	BusinessTime     int     //委托时间
	CombiNo          string  //组合编号
	ConfirmNo        string  //委托确认号
	EntrustAmount    int     //委托数量
	CancelAmount     int     //撤销数量
	EntrustDirection string  //委托方向
	EntrustNo        string  //委托编号
	EntrustPrice     float64 //委托价格
	EntrustStatus    string  //委托状态
	DealAmount       int     //成交数量
	DealBalance      float64 //成交金额
	DealPrice        float64 //成交均价
	FuturesDirection string  //开平方向
	InvestType       string  //投资类型
	MarketNo         string  //交易市场
	PriceType        string  //委托价格类型
	ReportNo         string  //申报编号
	ReportSeat       string  //申报席位
	RevokeCause      string  //废单原因
	StockCode        string  //证券代码
	StockholderID    string  //股东代码
	ThirdReff        string  //第三方系统自定义说明
	ExtsystemID      int     //第三方系统自定义号
}

// DealPushResp is the trade push data schema
type DealPushResp struct {
	OperatorNo       string  //操作员
	DealDate         int     //成交日期
	DealTime         int     //成交时间
	DealNo           string  //成交编号
	BatchNo          int     //委托批号
	EntrustNo        int     //委托编号
	MarketNo         string  //交易市场
	StockCode        string  //证券代码
	AccountCode      string  //账户编号
	CombiNo          string  //组合编号
	StockholderID    string  //股东代码
	ReportSeat       string  //申报席位
	EntrustDirection string  //委托方向
	FuturesDirection string  //开平方向
	EntrustAmount    int     //委托数量
	EntrustStatus    string  //委托状态
	DealAmount       int     //本次成交数量
	DealPrice        float64 //本次成交价格
	DealBalance      float64 //本次成交金额
	DealFee          float64 //本次费用
	TotalDealAmount  int     //累计成交数量
	TotalDealBalance float64 //累计成交金额
	CancelAmount     int     //撤销数量
	ReportDirection  string  //申报方向
	ExtsystemID      int     //第三方系统自定义号
	ThirdReff        string  //第三方系统自定义说明
}
