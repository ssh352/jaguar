

#ifndef _RESPONSE_H
#define _RESPONSE_H


typedef struct ConnectRet{
	int ErrorCode;
	char ErrorMsg[128];
}ConnectRet;

typedef struct Result{
	int FuncNo;
	int PacketType;
	int ReturnCode;

	int errorNo;
	const char *errorInfo;
	const char *MsgType;

	int ErrorCode;
	const char *ErrorMsg;          
	const char *MsgDetail;
	int DataCount;
	void *DataSet;
}Result;


typedef struct msginfo{
	const char* p;
	int len;
}msginfo;

typedef struct LoginResp{
	const char *user_token;
	const char *version_no;
}LoginResp;


typedef struct EntrustResp{
	int batch_no;
	int entrust_no;
	int extsystem_id;
}EntrustResp;

typedef struct WithdrawResp{
	int entrust_no;
	const char *MarketNo;
	const char *StockCode;
	const char *SuccessFlag; 
	const char *FailCause;
}WithdrawResp;

typedef struct QueryFundAssetResp{
	const char	*account_code;				//账户编号
	const char	*currency_code;				//币种
	double total_asset;						//总资产
	double nav;								//账户单位净值
	double yesterday_nav;					//昨日单位净值
	double current_balance;					//当前资金余额
	double begin_balance;					//期初资金余额
	double futu_deposit_balance;			//期货保证金账户余额
	double occupy_deposit_balance;			//期货占用保证金
	double futu_asset;						//期货资产（合约价值）
	double stock_asset;						//股票资产
	double bond_asset;						//债券资产
	double fund_asset;						//基金资产
	double repo_asset;						//回购资产
	double other_asset;						//其他资产
	double fund_share;						//基金总份额
	double fund_net_asset;					//基金净资产
	double payable_balance;					//应付款
	double receivable_balance;				//应收款
}QueryFundAssetResp;


typedef struct QueryPosResp{
	const char *account_code;		//账户编号
	const char *asset_no;			//资产单元编号
	const char *combi_no;			//组合编号
	const char *market_no;			//交易市场
	const char *stock_code;			//证券代码
	const char *stock_name;			//证券名称
	const char *stock_type;			//证券类型
	const char *stockholder_id;		//股东代码
	const char *hold_seat;			//持仓席位
	const char *invest_type;		//投资类型
	int current_amount;				//当前数量
	int enable_amount;				//可用数量
	double begin_cost;				//期初成本
	double current_cost;			//当前成本
	double cost_price;				//成本价
	double last_price;				//最新价
	double pre_buy_amount;			//买挂单数量
	double pre_sell_amount;			//卖挂单数量
	double pre_buy_balance;			//买挂单金额
	double pre_sell_balance;		//卖挂单金额
	int today_buy_amount;			//当日买入数量
	int today_sell_amount;			//当日卖出数量
	double today_buy_balance;		//当日买入金额
	double today_sell_balance;		//当日卖出金额
	double today_buy_fee;			//当日买费用
	double today_sell_fee;			//当日卖费用
	double floating_profit;			//浮盈
	double accumulate_profit;		//累计收益
	double total_profit;			//总收益
	
	void* nextdataptr;
}QueryPosResp;



typedef struct QueryEntrustResp{
	
	int			entrust_date;		//委托日期
	int			entrust_time;		//委托时间
	const char *operator_no;		//操作员编号
	int			batch_no;			//委托批号
	int			entrust_no;			//委托序号
	const char *report_no;			//申报编号
	int			extsystem_id;		//第三方系统自定义号
	const char *third_reff;			//第三方系统自定义说明
	const char *account_code;		//账户编号
	const char *asset_no;			//资产单元编号
	const char *combi_no;			//组合编号
	const char *stockholder_id;		//股东代码
	const char *report_seat;		//申报席位
	const char *market_no;			//交易市场
	const char *stock_code;			//证券代码
	const char *entrust_direction;	//委托方向
	const char *price_type;			//委托价格类型
	double		entrust_price;		//委托价格
	int			entrust_amount;		//委托数量
	double		pre_buy_frozen_balance;	//预买冻结金额
	double		pre_sell_balance;		//预卖金额
	const char *confirm_no;				//委托确认号
	const char *entrust_state;			//委托状态
	int			first_deal_time;		//首次成交时间
	int			deal_amount;			//成交数量
	double		deal_balance;			//成交金额
	double		deal_price;				//成交均价
	int			deal_times;				//分笔成交次数
	int			withdraw_amount;		//撤单数量
	const char *withdraw_cause;			//撤单原因
	const char *position_str;			//定位串
	const char *exchange_report_no;		//交易所申报编号

	void* nextentrustptr;
}QueryEntrustResp;


typedef struct QueryAccountResp{
	const char *account_code;			//账户编号
	const char *asset_no;				//资产单元编号
	double		enable_balance_t0;		//T+0可用资金
	double		enable_balance_t1;		//T+1可用资金
	double		current_balance;		//当前资金余额
}QueryAccountResp;


typedef struct EntrustPushResp{
	const char *account_code;		//账户编号
	int			batch_no;			//委托批号
	const char *operator_no;		//操作员
	int			business_date;		//委托日期
	int			business_time;		//委托时间
	const char *combi_no;			//组合编号
	const char *confirm_no;			//委托确认号
	int			entrust_amount;		//委托数量
	int			cancel_amount;		//撤销数量
	const char *entrust_direction;	//委托方向
	const char *entrust_no;			//委托编号
	double		entrust_price;		//委托价格
	const char *entrust_status;		//委托状态
	int			deal_amount;		//成交数量
	double		deal_balance;		//成交金额
	double		deal_price;			//成交均价
	const char *futures_direction;	//开平方向
	const char *invest_type;		//投资类型
	const char *market_no;			//交易市场
	const char *price_type;			//委托价格类型
	const char *report_no;			//申报编号
	const char *report_seat;		//申报席位
	const char *revoke_cause;		//废单原因
	const char *stock_code;			//证券代码
	const char *stockholder_id;		//股东代码
	const char *third_reff;			//第三方系统自定义说明
	int			extsystem_id;		//第三方系统自定义号
}EntrustPushResp;


typedef struct DealPushResp{
	const char *operator_no;		//操作员
	int			deal_date;			//成交日期
	int			deal_time;			//成交时间
	const char *deal_no;			//成交编号
	int			batch_no;			//委托批号
	int			entrust_no;			//委托编号
	const char *market_no;			//交易市场
	const char *stock_code;			//证券代码
	const char *account_code;		//账户编号
	const char *combi_no;			//组合编号
	const char *stockholder_id;		//股东代码
	const char *report_seat;		//申报席位
	const char *entrust_direction;	//委托方向
	const char *futures_direction;	//开平方向
	int			entrust_amount;		//委托数量
	const char *entrust_status;		//委托状态
	int			deal_amount;		//本次成交数量
	double		deal_price;			//本次成交价格
	double		deal_balance;		//本次成交金额
	double		deal_fee;			//本次费用
	double		total_deal_amount;	//累计成交数量
	double		total_deal_balance;	//累计成交金额
	int			cancel_amount;		//撤销数量
	const char *report_direction;	//申报方向
	int			extsystem_id;		//第三方系统自定义号
	const char *third_reff;			//第三方系统自定义说明
}DealPushResp;



#endif

