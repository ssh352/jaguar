package emsbase

// Portfolio is the structure which is composited by entrusts and stratrgy information.
type Portfolio struct {
	SecurityEntrusts []Entrust
	FutureEntrusts   []Entrust
	*ProductInfo
	*StrategyInfo
}

// StrategyInfo struct contains strategy information
type StrategyInfo struct {
	StrategyName string             // 策略名称
	TacticID     string             // 策略ID
	TacticType   string             // 策略ID
	Algorithm    string             // 算法
	Params       map[string]float64 // 算法交易参数
}

// ProductInfo struct contains product information
type ProductInfo struct {
	ProdID      string // 产品ID
	CombiNo     string // 组合编号
	StockAcc    string // 证券账户
	FutureAcc   string // 期货账户
	AccountID   string // 对接第三方系统账户
	AdapterName string // 交易Adapter
}

// Entrust stands for a entrust information
type Entrust struct {
	MarkerNo      string  // 市场代码
	StockCode     string  // 证券代码（带交易所后缀）
	TradeCode     string  // 证券代码
	Price         float64 // 委托价格
	Vol           int     // 委托数量
	OpenCloseFlag int     // 开平标记 1：开，2：平
	BS            int     // 买卖方向 1：买，2：卖
	TimeStamp     int64   // 订单生成时间戳
	Remark        string  // 备注信息
	ID            int64   // 用于标识订单
}

// EntrustType define the entrust type.
type EntrustType struct {
	OpenCloseFlag int // 开平标记 1：开，2：平
	BS            int // 买卖方向 1：买，2：卖
}

// EntrustTypeMap map the entrusttype with int.
var EntrustTypeMap map[int]EntrustType

func init() {
	EntrustTypeMap = make(map[int]EntrustType)
	EntrustTypeMap[OpenLong] = EntrustType{Open, Buy}
	EntrustTypeMap[OpenShort] = EntrustType{Open, Sell}
	EntrustTypeMap[CloseLong] = EntrustType{Close, Buy}
	EntrustTypeMap[CloseShort] = EntrustType{Close, Sell}
}

// ITarde define the function which should be implemented by trade adapter
type ITarde interface {
	Init() int
	LimitEntrust(e Entrust, AccountCode, ComiNo string)
	QueryPos(AccountCode string, ComiNo string)
	QueryAccount(AccountCode string, ComiNo string)
	QueryEntrustByAcc(AccountCode string, ComiNo string)
	QueryEntrustByEntrustNo(AccountCode string, ComiNo string, EntrustNo int)
}
