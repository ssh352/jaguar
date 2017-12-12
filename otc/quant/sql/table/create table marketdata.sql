create table realtimemarketdata(
	Code 		varchar(32),
	ActionDay 		int,
	TradingDay 	int,
	Time 			int,
	Status 		int,
	PreClose 		double,
	Open 			double,
	High 			double,
	Low 			double,
	Last 			double,
	AskPrice1		double,
	AskPrice2		double,
	AskPrice3		double,
	AskPrice4		double,
	AskPrice5		double,
	AskPrice6		double,
	AskPrice7		double,
	AskPrice8		double,
	AskPrice9		double,
	AskPrice10		double,
	AskVol1		int,
	AskVol2		int,
	AskVol3		int,
	AskVol4		int,
	AskVol5		int,
	AskVol6		int,
	AskVol7		int,
	AskVol8		int,
	AskVol9		int,
	AskVol10		int,
	BidPrice1		double,
	BidPrice2		double,
	BidPrice3		double,
	BidPrice4		double,
	BidPrice5		double,
	BidPrice6		double,
	BidPrice7		double,
	BidPrice8		double,
	BidPrice9		double,
	BidPrice10		double,
	BidVol1		int,
	BidVol2		int,
	BidVol3		int,
	BidVol4		int,
	BidVol5		int,
	BidVol6		int,
	BidVol7		int,
	BidVol8		int,
	BidVol9		int,
	BidVol10		int,
	NumTrades		int,
	Volume 		bigint,
	Turnover 		bigint,
	TotalBidVol 	bigint,
	TotalAskVol 	bigint,
	WeightedAvgBidPrice 	double,
	WeightedAvgAskPrice 	double,
	IOPV 					double,
	YieldToMaturity 		double,
	HighLimited 			double,
	LowLimited 			double,
	Prefix 				varchar(4),
	Syl1 					int,
	Syl2 					int,
	SD2 					int
)ENGINE = MEMORY