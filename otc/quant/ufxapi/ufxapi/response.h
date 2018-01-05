

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
	const char	*account_code;				//�˻����
	const char	*currency_code;				//����
	double total_asset;						//���ʲ�
	double nav;								//�˻���λ��ֵ
	double yesterday_nav;					//���յ�λ��ֵ
	double current_balance;					//��ǰ�ʽ����
	double begin_balance;					//�ڳ��ʽ����
	double futu_deposit_balance;			//�ڻ���֤���˻����
	double occupy_deposit_balance;			//�ڻ�ռ�ñ�֤��
	double futu_asset;						//�ڻ��ʲ�����Լ��ֵ��
	double stock_asset;						//��Ʊ�ʲ�
	double bond_asset;						//ծȯ�ʲ�
	double fund_asset;						//�����ʲ�
	double repo_asset;						//�ع��ʲ�
	double other_asset;						//�����ʲ�
	double fund_share;						//�����ܷݶ�
	double fund_net_asset;					//�����ʲ�
	double payable_balance;					//Ӧ����
	double receivable_balance;				//Ӧ�տ�
}QueryFundAssetResp;


typedef struct QueryPosResp{
	const char *account_code;		//�˻����
	const char *asset_no;			//�ʲ���Ԫ���
	const char *combi_no;			//��ϱ��
	const char *market_no;			//�����г�
	const char *stock_code;			//֤ȯ����
	const char *stock_name;			//֤ȯ����
	const char *stock_type;			//֤ȯ����
	const char *stockholder_id;		//�ɶ�����
	const char *hold_seat;			//�ֲ�ϯλ
	const char *invest_type;		//Ͷ������
	int current_amount;				//��ǰ����
	int enable_amount;				//��������
	double begin_cost;				//�ڳ��ɱ�
	double current_cost;			//��ǰ�ɱ�
	double cost_price;				//�ɱ���
	double last_price;				//���¼�
	double pre_buy_amount;			//��ҵ�����
	double pre_sell_amount;			//���ҵ�����
	double pre_buy_balance;			//��ҵ����
	double pre_sell_balance;		//���ҵ����
	int today_buy_amount;			//������������
	int today_sell_amount;			//������������
	double today_buy_balance;		//����������
	double today_sell_balance;		//�����������
	double today_buy_fee;			//���������
	double today_sell_fee;			//����������
	double floating_profit;			//��ӯ
	double accumulate_profit;		//�ۼ�����
	double total_profit;			//������
	
	void* nextdataptr;
}QueryPosResp;



typedef struct QueryEntrustResp{
	
	int			entrust_date;		//ί������
	int			entrust_time;		//ί��ʱ��
	const char *operator_no;		//����Ա���
	int			batch_no;			//ί������
	int			entrust_no;			//ί�����
	const char *report_no;			//�걨���
	int			extsystem_id;		//������ϵͳ�Զ����
	const char *third_reff;			//������ϵͳ�Զ���˵��
	const char *account_code;		//�˻����
	const char *asset_no;			//�ʲ���Ԫ���
	const char *combi_no;			//��ϱ��
	const char *stockholder_id;		//�ɶ�����
	const char *report_seat;		//�걨ϯλ
	const char *market_no;			//�����г�
	const char *stock_code;			//֤ȯ����
	const char *entrust_direction;	//ί�з���
	const char *price_type;			//ί�м۸�����
	double		entrust_price;		//ί�м۸�
	int			entrust_amount;		//ί������
	double		pre_buy_frozen_balance;	//Ԥ�򶳽���
	double		pre_sell_balance;		//Ԥ�����
	const char *confirm_no;				//ί��ȷ�Ϻ�
	const char *entrust_state;			//ί��״̬
	int			first_deal_time;		//�״γɽ�ʱ��
	int			deal_amount;			//�ɽ�����
	double		deal_balance;			//�ɽ����
	double		deal_price;				//�ɽ�����
	int			deal_times;				//�ֱʳɽ�����
	int			withdraw_amount;		//��������
	const char *withdraw_cause;			//����ԭ��
	const char *position_str;			//��λ��
	const char *exchange_report_no;		//�������걨���

	void* nextentrustptr;
}QueryEntrustResp;


typedef struct QueryAccountResp{
	const char *account_code;			//�˻����
	const char *asset_no;				//�ʲ���Ԫ���
	double		enable_balance_t0;		//T+0�����ʽ�
	double		enable_balance_t1;		//T+1�����ʽ�
	double		current_balance;		//��ǰ�ʽ����
}QueryAccountResp;


typedef struct EntrustPushResp{
	const char *account_code;		//�˻����
	int			batch_no;			//ί������
	const char *operator_no;		//����Ա
	int			business_date;		//ί������
	int			business_time;		//ί��ʱ��
	const char *combi_no;			//��ϱ��
	const char *confirm_no;			//ί��ȷ�Ϻ�
	int			entrust_amount;		//ί������
	int			cancel_amount;		//��������
	const char *entrust_direction;	//ί�з���
	const char *entrust_no;			//ί�б��
	double		entrust_price;		//ί�м۸�
	const char *entrust_status;		//ί��״̬
	int			deal_amount;		//�ɽ�����
	double		deal_balance;		//�ɽ����
	double		deal_price;			//�ɽ�����
	const char *futures_direction;	//��ƽ����
	const char *invest_type;		//Ͷ������
	const char *market_no;			//�����г�
	const char *price_type;			//ί�м۸�����
	const char *report_no;			//�걨���
	const char *report_seat;		//�걨ϯλ
	const char *revoke_cause;		//�ϵ�ԭ��
	const char *stock_code;			//֤ȯ����
	const char *stockholder_id;		//�ɶ�����
	const char *third_reff;			//������ϵͳ�Զ���˵��
	int			extsystem_id;		//������ϵͳ�Զ����
}EntrustPushResp;


typedef struct DealPushResp{
	const char *operator_no;		//����Ա
	int			deal_date;			//�ɽ�����
	int			deal_time;			//�ɽ�ʱ��
	const char *deal_no;			//�ɽ����
	int			batch_no;			//ί������
	int			entrust_no;			//ί�б��
	const char *market_no;			//�����г�
	const char *stock_code;			//֤ȯ����
	const char *account_code;		//�˻����
	const char *combi_no;			//��ϱ��
	const char *stockholder_id;		//�ɶ�����
	const char *report_seat;		//�걨ϯλ
	const char *entrust_direction;	//ί�з���
	const char *futures_direction;	//��ƽ����
	int			entrust_amount;		//ί������
	const char *entrust_status;		//ί��״̬
	int			deal_amount;		//���γɽ�����
	double		deal_price;			//���γɽ��۸�
	double		deal_balance;		//���γɽ����
	double		deal_fee;			//���η���
	double		total_deal_amount;	//�ۼƳɽ�����
	double		total_deal_balance;	//�ۼƳɽ����
	int			cancel_amount;		//��������
	const char *report_direction;	//�걨����
	int			extsystem_id;		//������ϵͳ�Զ����
	const char *third_reff;			//������ϵͳ�Զ���˵��
}DealPushResp;



#endif

