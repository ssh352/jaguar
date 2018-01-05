


#include <vector>
#include <string>
#include <iostream>
#include "t2sdk_interface.h"
#include "ufxutil.h"
#include "response.h"
using namespace std;


void UFX_10001_unPacker(Result* ret, void** resps, IF2UnPacker* responseUnPacker){
	void* resp			= malloc(sizeof(LoginResp));
	resps[0]			= resp;
	//printf("UFX_10001_unPacker---%d: %u\n", 0 , resp);
	LoginResp *ptr		= (LoginResp *)(resp);
	ptr->user_token		= responseUnPacker->GetStr("user_token");
	ptr->version_no		= responseUnPacker->GetStr("version_no");
	ret->DataSet		= ptr;
}

void UFX_91001_unPacker(Result* ret, void** resps, IF2UnPacker* responseUnPacker){
	void* resp			= malloc(sizeof(LoginResp));
	resps[0]			= resp;
	EntrustResp *ptr	= (EntrustResp *)(resp);
	ptr->batch_no		= responseUnPacker->GetInt("batch_no");
	ptr->entrust_no		= responseUnPacker->GetInt("entrust_no");
	ptr->extsystem_id	= responseUnPacker->GetInt("extsystem_id");
	ret->DataSet		= ptr;
}

void UFX_91101_unPacker(Result* ret,  void** resps, IF2UnPacker* responseUnPacker){
	void* resp			= malloc(sizeof(WithdrawResp));
	resps[0]			= resp;
	WithdrawResp *ptr	= (WithdrawResp *)(resp);
	ptr->entrust_no		= responseUnPacker->GetInt("entrust_no");
	ptr->MarketNo		= responseUnPacker->GetStr("market_no");
	ptr->StockCode		= responseUnPacker->GetStr("stockkl_code");
	ptr->SuccessFlag	= responseUnPacker->GetStr("success_flag");
	ptr->FailCause		= responseUnPacker->GetStr("fail_cause");
	ret->DataSet		= ptr;
}


void UFX_35003_unPacker(Result* ret, void** resps, IF2UnPacker* responseUnPacker){
	void* resp					 = malloc(sizeof(QueryFundAssetResp));
	resps[0]					 = resp;

	QueryFundAssetResp *ptr		 = (QueryFundAssetResp *)(resp);
	ptr->account_code			 = responseUnPacker->GetStr("account_code");
	ptr->currency_code			 = responseUnPacker->GetStr("currency_code");
	ptr->total_asset			 = responseUnPacker->GetDouble("total_asset");
	ptr->nav					 = responseUnPacker->GetDouble("nav");
	ptr->yesterday_nav			 = responseUnPacker->GetDouble("yesterday_nav");
	ptr->current_balance		 = responseUnPacker->GetDouble("current_balance");
	ptr->begin_balance			 = responseUnPacker->GetDouble("begin_balance");
	ptr->futu_deposit_balance	 = responseUnPacker->GetDouble("futu_deposit_balance");
	ptr->occupy_deposit_balance	 = responseUnPacker->GetDouble("occupy_deposit_balance");
	ptr->futu_asset				 = responseUnPacker->GetDouble("futu_asset");
	ptr->stock_asset			 = responseUnPacker->GetDouble("stock_asset");
	ptr->bond_asset				 = responseUnPacker->GetDouble("bond_asset");
	ptr->fund_asset				 = responseUnPacker->GetDouble("fund_asset");
	ptr->repo_asset				 = responseUnPacker->GetDouble("repo_asset");
	ptr->other_asset			 = responseUnPacker->GetDouble("other_asset");
	ptr->fund_share				 = responseUnPacker->GetDouble("fund_share");
	ptr->fund_net_asset			 = responseUnPacker->GetDouble("fund_net_asset");
	ptr->payable_balance		 = responseUnPacker->GetDouble("payable_balance");
	ptr->receivable_balance		 = responseUnPacker->GetDouble("receivable_balance");

	ret->DataSet				 = ptr;
}

void UFX_31001_unPacker(Result* ret, void** resps, IF2UnPacker* responseUnPacker){
	void* resp					 = malloc(sizeof(QueryPosResp));
	ret->DataSet				 = resp; // the fisrt element
	for(int i=0; i<ret->DataCount; i++){
		//printf("UFX_31001_unPacker---%d: %u\n", i , resp);
		resps[i]				 = resp;
		QueryPosResp *ptr		 = (QueryPosResp *)(resp);
		ptr->account_code		 = responseUnPacker->GetStr("account_code");
		ptr->asset_no			 = responseUnPacker->GetStr("asset_no");
		ptr->combi_no			 = responseUnPacker->GetStr("combi_no");
		ptr->market_no			 = responseUnPacker->GetStr("market_no");
		ptr->stock_code			 = responseUnPacker->GetStr("stock_code");
		ptr->stock_name			 = responseUnPacker->GetStr("stock_name");
		ptr->stockholder_id		 = responseUnPacker->GetStr("stockholder_id");
		ptr->hold_seat			 = responseUnPacker->GetStr("hold_seat");
		ptr->invest_type		 = responseUnPacker->GetStr("invest_type");
		ptr->current_amount		 = responseUnPacker->GetInt("current_amount");
		ptr->enable_amount		 = responseUnPacker->GetInt("enable_amount");
		ptr->begin_cost			 = responseUnPacker->GetDouble("begin_cost");
		ptr->current_cost		 = responseUnPacker->GetDouble("current_cost");
		ptr->cost_price			 = responseUnPacker->GetDouble("cost_price");
		ptr->last_price			 = responseUnPacker->GetDouble("last_price");
		ptr->pre_buy_amount		 = responseUnPacker->GetDouble("pre_buy_amount");
		ptr->pre_sell_amount	 = responseUnPacker->GetDouble("pre_sell_amount");
		ptr->pre_buy_balance	 = responseUnPacker->GetDouble("pre_buy_balance");
		ptr->pre_sell_balance	 = responseUnPacker->GetDouble("pre_sell_balance");
		ptr->today_buy_amount	 = responseUnPacker->GetInt("today_buy_amount");
		ptr->today_sell_amount	 = responseUnPacker->GetInt("today_sell_amount");
		ptr->today_buy_balance	 = responseUnPacker->GetDouble("today_buy_balance");
		ptr->today_sell_balance	 = responseUnPacker->GetDouble("today_sell_balance");
		ptr->today_buy_fee		 = responseUnPacker->GetDouble("today_buy_fee");
		ptr->today_sell_fee		 = responseUnPacker->GetDouble("today_sell_fee");
		ptr->floating_profit	 = responseUnPacker->GetDouble("floating_profit");
		ptr->accumulate_profit	 = responseUnPacker->GetDouble("accumulate_profit");
		ptr->total_profit		 = responseUnPacker->GetDouble("total_profit");
		
		responseUnPacker->Next();
		resp					 = malloc(sizeof(QueryPosResp));
		ptr->nextdataptr		 = resp;
	}
}


void UFX_32001_unPacker(Result* ret, void** resps, IF2UnPacker* responseUnPacker){
	void* resp					 = malloc(sizeof(QueryEntrustResp));
	ret->DataSet				 = resp; // the fisrt element
	for(int i=0; i<ret->DataCount; i++){
		//printf("UFX_31001_unPacker---%d: %u\n", i , resp);
		resps[i]				 = resp;
		QueryEntrustResp *ptr		 = (QueryEntrustResp *)(resp);
		ptr->entrust_date			 = responseUnPacker->GetInt("entrust_date");
		ptr->entrust_time			 = responseUnPacker->GetInt("entrust_time");
		ptr->operator_no			 = responseUnPacker->GetStr("operator_no");
		ptr->batch_no				 = responseUnPacker->GetInt("batch_no");
		ptr->entrust_no				 = responseUnPacker->GetInt("entrust_no");
		ptr->report_no				 = responseUnPacker->GetStr("report_no");
		ptr->extsystem_id			 = responseUnPacker->GetInt("extsystem_id");
		ptr->third_reff				 = responseUnPacker->GetStr("third_reff");
		ptr->account_code			 = responseUnPacker->GetStr("account_code");
		ptr->asset_no				 = responseUnPacker->GetStr("asset_no");
		ptr->combi_no				 = responseUnPacker->GetStr("combi_no");
		ptr->stockholder_id			 = responseUnPacker->GetStr("stockholder_id");
		ptr->report_seat			 = responseUnPacker->GetStr("report_seat");
		ptr->market_no				 = responseUnPacker->GetStr("market_no");
		ptr->stock_code				 = responseUnPacker->GetStr("stock_code");
		ptr->entrust_direction		 = responseUnPacker->GetStr("entrust_direction");
		ptr->price_type				 = responseUnPacker->GetStr("price_type");
		ptr->entrust_price			 = responseUnPacker->GetDouble("entrust_price");
		ptr->entrust_amount			 = responseUnPacker->GetInt("entrust_amount");
		ptr->pre_buy_frozen_balance	 = responseUnPacker->GetDouble("pre_buy_frozen_balance");
		ptr->pre_sell_balance		 = responseUnPacker->GetDouble("pre_sell_balance");
		ptr->confirm_no				 = responseUnPacker->GetStr("confirm_no");
		ptr->entrust_state			 = responseUnPacker->GetStr("entrust_state");
		ptr->first_deal_time		 = responseUnPacker->GetInt("first_deal_time");
		ptr->deal_amount			 = responseUnPacker->GetInt("deal_amount");
		ptr->deal_balance			 = responseUnPacker->GetDouble("deal_balance");
		ptr->deal_price				 = responseUnPacker->GetDouble("deal_price");
		ptr->deal_times				 = responseUnPacker->GetInt("deal_times");
		ptr->withdraw_amount		 = responseUnPacker->GetInt("withdraw_amount");
		ptr->withdraw_cause			 = responseUnPacker->GetStr("withdraw_cause");
		ptr->position_str			 = responseUnPacker->GetStr("position_str");
		ptr->exchange_report_no		 = responseUnPacker->GetStr("exchange_report_no");

		responseUnPacker->Next();
		resp					 = malloc(sizeof(QueryEntrustResp));
		ptr->nextentrustptr		 = resp;
	}
}


void UFX_34001_unPacker(Result* ret, void** resps, IF2UnPacker* responseUnPacker){
	void* resp				 = malloc(sizeof(QueryAccountResp));
	resps[0]				 = resp;
	QueryAccountResp *ptr	 = (QueryAccountResp *)(resp);
	ptr->account_code		 = responseUnPacker->GetStr("account_code");
	ptr->asset_no			 = responseUnPacker->GetStr("asset_no");
	ptr->enable_balance_t0	 = responseUnPacker->GetDouble("enable_balance_t0");
	ptr->enable_balance_t1	 = responseUnPacker->GetDouble("enable_balance_t1");
	ptr->current_balance	 = responseUnPacker->GetDouble("current_balance");
	ret->DataSet			 = ptr;
}

void UFX_HANDLE_ENTRUST_PUSH(Result* ret, void* resp, IF2UnPacker* lpUnPack){
	resp				 = malloc(sizeof(EntrustPushResp));
	EntrustPushResp *ptr = (EntrustPushResp*)resp;
	ptr->operator_no		 = lpUnPack->GetStr("operator_no");
	ptr->account_code		 = lpUnPack->GetStr("account_code");
	ptr->batch_no			 = lpUnPack->GetInt("batch_no");
	ptr->business_date		 = lpUnPack->GetInt("business_date");
	ptr->business_time		 = lpUnPack->GetInt("business_time");
	ptr->combi_no			 = lpUnPack->GetStr("combi_no");
	ptr->confirm_no			 = lpUnPack->GetStr("confirm_no");
	ptr->entrust_amount		 = lpUnPack->GetInt("entrust_amount");
	ptr->cancel_amount		 = lpUnPack->GetInt("cancel_amount");
	ptr->entrust_direction	 = lpUnPack->GetStr("entrust_direction");
	ptr->entrust_no			 = lpUnPack->GetStr("entrust_no");
	ptr->entrust_price		 = lpUnPack->GetDouble("entrust_price");
	ptr->entrust_status		 = lpUnPack->GetStr("entrust_status");
	ptr->deal_amount		 = lpUnPack->GetInt("deal_amount");
	ptr->deal_balance		 = lpUnPack->GetDouble("deal_balance");
	ptr->deal_price			 = lpUnPack->GetDouble("deal_price");
	ptr->futures_direction	 = lpUnPack->GetStr("futures_direction");
	ptr->invest_type		 = lpUnPack->GetStr("invest_type");
	ptr->market_no			 = lpUnPack->GetStr("market_no");
	ptr->price_type			 = lpUnPack->GetStr("price_type");
	ptr->report_no			 = lpUnPack->GetStr("report_no");
	ptr->report_seat		 = lpUnPack->GetStr("report_seat");
	ptr->revoke_cause		 = lpUnPack->GetStr("revoke_cause");
	ptr->stock_code			 = lpUnPack->GetStr("stock_code");
	ptr->stockholder_id		 = lpUnPack->GetStr("stockholder_id");
	ptr->third_reff			 = lpUnPack->GetStr("third_reff");
	ptr->extsystem_id		 = lpUnPack->GetInt("extsystem_id");

	ret->DataSet			 = ptr;
}

void UFX_HANDLE_DEAL_PUSH(Result* ret, void* resp, IF2UnPacker* lpUnPack){
	resp					 = malloc(sizeof(DealPushResp));
	DealPushResp *ptr		 = (DealPushResp*)resp;

	ptr->operator_no		 = lpUnPack->GetStr("operator_no");
	ptr->deal_date			 = lpUnPack->GetInt("deal_date");
	ptr->deal_time			 = lpUnPack->GetInt("deal_time");
	ptr->deal_no			 = lpUnPack->GetStr("deal_no");
	ptr->batch_no			 = lpUnPack->GetInt("batch_no");
	ptr->entrust_no			 = lpUnPack->GetInt("entrust_no");
	ptr->market_no			 = lpUnPack->GetStr("market_no");
	ptr->stock_code			 = lpUnPack->GetStr("stock_code");
	ptr->account_code		 = lpUnPack->GetStr("account_code");
	ptr->combi_no			 = lpUnPack->GetStr("combi_no");
	ptr->stockholder_id		 = lpUnPack->GetStr("stockholder_id");
	ptr->report_seat		 = lpUnPack->GetStr("report_seat");
	ptr->entrust_direction	 = lpUnPack->GetStr("entrust_direction");
	ptr->futures_direction	 = lpUnPack->GetStr("futures_direction");
	ptr->entrust_amount		 = lpUnPack->GetInt("entrust_amount");
	ptr->entrust_status		 = lpUnPack->GetStr("entrust_status");
	ptr->deal_amount		 = lpUnPack->GetInt("deal_amount");
	ptr->deal_price			 = lpUnPack->GetDouble("deal_price");
	ptr->deal_balance		 = lpUnPack->GetDouble("deal_balance");
	ptr->deal_fee			 = lpUnPack->GetDouble("deal_fee");
	ptr->total_deal_amount	 = lpUnPack->GetDouble("total_deal_amount");
	ptr->total_deal_balance	 = lpUnPack->GetDouble("total_deal_balance");
	ptr->cancel_amount		 = lpUnPack->GetInt("cancel_amount");
	ptr->report_direction	 = lpUnPack->GetStr("report_direction");
	ptr->extsystem_id		 = lpUnPack->GetInt("extsystem_id");
	ptr->third_reff			 = lpUnPack->GetStr("third_reff");

	ret->DataSet			 = ptr;
	
}

void ShowPacket(IF2UnPacker* unPacker){
	printf("---------------ShowPacket-------------\n");
	int i = 0, t = 0, j = 0, k = 0;
	for (i = 0; i < unPacker->GetDatasetCount(); ++i){
		// 设置当前结果集
		unPacker->SetCurrentDatasetByIndex(i);
		// 打印字段
		for (t = 0; t < unPacker->GetColCount(); ++t)		{
			printf("%20s", unPacker->GetColName(t));
		}
		putchar('\n');
		// 打印所有记录
		for (j = 0; j < (int)unPacker->GetRowCount(); ++j){
			// 打印每条记录
			for (k = 0; k < unPacker->GetColCount(); ++k){
				switch (unPacker->GetColType(k))
				{
				case 'I':
					printf("%20d", unPacker->GetIntByIndex(k));
					break;
				case 'C':
					printf("%20c", unPacker->GetCharByIndex(k));
					break;
				case 'S':
					printf("%20s", unPacker->GetStrByIndex(k));
					break;
				case 'F':
					printf("%20f", unPacker->GetDoubleByIndex(k));
					break;
				case 'R':
					{
						int nLength = 0;
						void *lpData = unPacker->GetRawByIndex(k, &nLength);
						// 对2进制数据进行处理
						break;
					}
				default:
					// 未知数据类型
					printf("未知数据类型。\n");
					break;
				}
			}
			putchar('\n');
			unPacker->Next();
		}
		putchar('\n');
	}
}


//
//void ShowSubscribe(int subIndex, LPSUBSCRIBE_RECVDATA lpRecvData, map<int, CSubscribeParamInterface*> gAllSubscribeParam){
//
//	map<int, CSubscribeParamInterface*>::iterator it = gAllSubscribeParam.find(subIndex);
//	if (it == gAllSubscribeParam.end()){
//		printf("没有这个订阅项\n");
//		return;
//	}
//	CSubscribeParamInterface* lpSubParam = (*it).second;
//
//	printf("----------订阅项部分-------\n");
//	printf("主题名字：           %s\n",lpSubParam->GetTopicName());
//	printf("附加数据长度：       %d\n",lpRecvData->iAppDataLen);
//	if (lpRecvData->iAppDataLen > 0)
//	{
//		printf("附加数据：           %s\n",lpRecvData->lpAppData);
//	}
//	printf("过滤字段部分：\n");
//	if (lpRecvData->iFilterDataLen > 0)
//	{
//		IF2UnPacker* lpUnpack = NewUnPacker(lpRecvData->lpFilterData,lpRecvData->iFilterDataLen);
//		lpUnpack->AddRef();
//		ShowPacket(lpUnpack);
//		lpUnpack->Release();
//	}
//	printf("---------------------------\n");
//}