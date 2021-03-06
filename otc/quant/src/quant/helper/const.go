package helper

const (
	// QuantConfigFile is the configure file name for quant project
	QuantConfigFile string = "./conf/quant.ini"
	// QuantLogConfigFile is the log configure file name for quant project
	QuantLogConfigFile string = "./conf/log.xml"
	// RedisKey is the key stored quote data
	RedisKey string = "MarketMap_test"

	EntrustRespPushData string = "ENTRUST"
	PortPushData        string = "ORIGENTRUST"
	TradePushData       string = "TRADE"

	// EMS config define
	ConfigEMSSessionName  string = "emsmodule"
	ConfigEMSPullAddr     string = "pull_addr"
	ConfigEMSTradeAdapter string = "trade_adapter"
	ConfigEMSUfx          string = "ufx"
	ConfigEMSTimeout      string = "timeout"
	ConfigEMSPortQueueLen string = "portqueuelen"
	ConfigEMSSqlLen       string = "sqllen"

	// Mysql config define
	ConfigMysqlSessionName string = "mysql"
	ConfigMysqlUserName    string = "mysqlusername"
	ConfigMysqlPwd         string = "mysqlpwd"
	ConfigMysqlUrl         string = "mysqlurl"

	// OMS config define
	ConfigOMSSessionName           string = "omsmodule"
	ConfigOMSPullAddr              string = "pull_addr"
	ConfigOMSPublishAddr           string = "publish_addr"
	ConfigOMSReqAddr               string = "req_addr"
	ConfigOMSEntrustLen            string = "entrust_sql_len"
	ConfigOMSTradeLen              string = "trade_sql_len"
	ConfigOMSPollTimeOut           string = "rb_pull_time_out"
	ConfigOMSEntrustUpdateBatchNum string = "entrust_update_batch_num"
	ConfigOMSTradeInsertBatchNum   string = "trade_insert_batch_num"

	// HQ config define
	ConfigHQSessionName string = "hqmodule"
	ConfigHQPublishAddr string = "publish_addr"
)
