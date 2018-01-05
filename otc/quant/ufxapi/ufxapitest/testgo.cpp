



#include <windows.h>
#include <stdio.h>
#include "ufxapi.h"
#include "response.h"

int UFXCallBack(Result* ret){
	if(ret != NULL){
		printf("************** C UFXCallBack funcno: %d **************\n", ret->FuncNo);
	}
	
	if(ret->FuncNo == 31001){
		// 持仓结果查询解析
		/*QueryPosResp* Dataptr = (QueryPosResp*) ret->DataSet;
		printf("************** C UFXCallBack funcno: %d, ndatacount:%d, next %u **************\n", ret->FuncNo, ret->DataCount, Dataptr->nextdataptr);
		int nData = ret->DataCount;
		for( int i=0; i<nData; i++){
			printf("stock_code: %s, enable_amount: %d, last_price: %f, cost_price: %f, total_profit: %f, floating_profit: %f, accumulate_profit: %f\n", 
					Dataptr->stock_code, Dataptr->enable_amount, Dataptr->last_price, Dataptr->cost_price, Dataptr->total_profit, Dataptr->floating_profit, Dataptr->accumulate_profit);
			Dataptr = (QueryPosResp*)Dataptr->nextdataptr;
		}*/
	}else if(ret->FuncNo == 32001){
		// 委托结果查询解析
	/*	QueryEntrustResp* Dataptr = (QueryEntrustResp*) ret->DataSet;
		printf("************** C UFXCallBack funcno: %d, ndatacount:%d, next %u **************\n", ret->FuncNo, ret->DataCount, Dataptr->nextentrustptr);
		int nData = ret->DataCount;
		for( int i=0; i<nData; i++){
			printf("stock_code: %s, entrust_date: %d, entrust_time: %d, entrust_amount: %d, entrust_price: %f\n", 
					Dataptr->stock_code, Dataptr->entrust_date, Dataptr->entrust_time, Dataptr->entrust_amount, Dataptr->entrust_price);
			Dataptr = (QueryEntrustResp*)Dataptr->nextentrustptr;
		}*/
	}else if(ret->FuncNo == 91001){
		// 按委托编号查询
		/*EntrustResp* Dataptr = (EntrustResp*) ret->DataSet;
		int entrust_no = Dataptr->entrust_no;
		printf("entrust_no: %d\n", entrust_no);
		QueryEntrustByEntrustNo("1007", "10072", entrust_no);*/
	}else if(ret->FuncNo == 34001){
		/*QueryAccountResp* Dataptr = (QueryAccountResp*) ret->DataSet;
		printf("account: %s enable_balance_t0: %f enable_balance_t1: %f current_balance: %f\n", 
			Dataptr->account_code, Dataptr->enable_balance_t0, Dataptr->enable_balance_t1, Dataptr->current_balance);*/
	}
	return 0;
}

typedef int (* FuncRegisterCallBackPtr)(ufxcallback);
typedef ConnectRet (* FuncConnectPtr)(const char *, int);
typedef int (* FuncSubConnectPtr)();
typedef int (* FuncLoginPtr)();
typedef int (* FuncLimitEntrustPtr)(const char *, const char *, const char *, const char *, const char *, double, int);
typedef int (* FuncWithdrawPtr)(int);
typedef int (* FuncQueryPosPtr)(const char*, const char*);
typedef int (* FuncQueryEntrustByAccPtr)(const char* account, const char* combi_bo);
typedef int (* FuncQueryEntrustByEntrustNoPtr)(const char*, const char*, int);
typedef int (* FuncQueryAccountPtr)(const char*, const char*);
typedef int (* FuncTestretmsgPtr)();

HINSTANCE 						G_HANDLE 						= NULL;
FuncConnectPtr 					G_ConnectPtr 					= NULL;
FuncRegisterCallBackPtr 		G_RegisterCallBackPtr 			= NULL;
FuncLoginPtr 					G_LoginPtr 						= NULL;
FuncSubConnectPtr 				G_SubConnectPtr 				= NULL;
FuncLimitEntrustPtr 			G_LimitEntrustPtr				= NULL;
FuncWithdrawPtr					G_WithdrawPtr					= NULL;
FuncQueryPosPtr					G_QueryPosPtr					= NULL;
FuncQueryEntrustByAccPtr 		G_QueryEntrustByAccPtr 			= NULL;
FuncQueryEntrustByEntrustNoPtr 	G_QueryEntrustByEntrustNoPtr 	= NULL;
FuncQueryAccountPtr 			G_QueryAccountPtr 				= NULL;
FuncTestretmsgPtr				G_TestretmsgPtr					= NULL;



ConnectRet UFXConnect(const char* serverIp, int sync){
	return G_ConnectPtr(serverIp, sync);
}

int RegisterUFXCallBack(){
	ufxcallback func = &UFXCallBack;
	if(G_RegisterCallBackPtr != NULL){
		return G_RegisterCallBackPtr(func);
	}else{
		return -2;
	}
}

int UFXLogin(){
	if(G_LoginPtr != NULL){
		return G_LoginPtr();
	}else{
		return -2;
	}
}

int UFXSubConnect(){
	if(G_SubConnectPtr != NULL){
		return G_SubConnectPtr();
	}else{
		return -2;
	}
}

int UFXLimitEntrust(const char* account_code, const char* market_no, const char* stock_code,
								   const char* combi_no, const char* BS, double price, int vol){
	if(G_LimitEntrustPtr != NULL){
		return G_LimitEntrustPtr(account_code, market_no, stock_code, combi_no, BS, price, vol);
	}else{
		return -2;
	}
}

int UFXWithdraw(int entrustno){
	if(G_WithdrawPtr != NULL){
		return G_WithdrawPtr(entrustno);
	}else{
		return -2;
	}
}

int UFXQueryPos(const char* account, const char* combi_bo){
	if(G_QueryPosPtr != NULL){
		return G_QueryPosPtr(account, combi_bo);
	}else{
		return -2;
	}
}


int UFXQueryEntrustByAcc(const char* account, const char* combi_no){
	if(G_QueryEntrustByAccPtr != NULL){
		return G_QueryEntrustByAccPtr(account, combi_no);
	}else{
		return -2;
	}
}

int UFXQueryEntrustByEntrustNo(const char* account, const char* combi_no, int EntrustNo){
	if(G_QueryEntrustByEntrustNoPtr != NULL){
		return G_QueryEntrustByEntrustNoPtr(account, combi_no, EntrustNo);
	}else{
		return -2;
	}
}

int UFXQueryAccount(const char* account, const char* combi_no){
	if(G_QueryAccountPtr != NULL){
		return G_QueryAccountPtr(account, combi_no);
	}else{
		return -2;
	}
}

int UFXClose(){
	if(G_HANDLE != NULL){
		FreeLibrary(G_HANDLE);
	}
	return 0;
}

int UFXInit(){
	G_HANDLE 				= LoadLibrary("E:\\otc\\quant\\ufxapi.dll");
	printf("G_HANDLE: %d\n", G_HANDLE);
	G_ConnectPtr 			= (FuncConnectPtr)GetProcAddress(G_HANDLE, "Connect");
	G_RegisterCallBackPtr 	= (FuncRegisterCallBackPtr)GetProcAddress(G_HANDLE, "RegisterCallBack");
	G_LoginPtr 				= (FuncLoginPtr)GetProcAddress(G_HANDLE, "Login");
	G_SubConnectPtr 		= (FuncSubConnectPtr)GetProcAddress(G_HANDLE, "SubConnect");
	G_LimitEntrustPtr		= (FuncLimitEntrustPtr)GetProcAddress(G_HANDLE, "LimitEntrust");
	G_WithdrawPtr			= (FuncWithdrawPtr)GetProcAddress(G_HANDLE, "Withdraw");
	G_QueryPosPtr			= (FuncQueryPosPtr)GetProcAddress(G_HANDLE, "QueryPos");
	G_QueryEntrustByAccPtr 		 = (FuncQueryEntrustByAccPtr)GetProcAddress(G_HANDLE, "QueryEntrustByAcc");
	G_QueryEntrustByEntrustNoPtr = (FuncQueryEntrustByEntrustNoPtr)GetProcAddress(G_HANDLE, "QueryEntrustByEntrustNo");
	G_QueryAccountPtr 			 = (FuncQueryAccountPtr)GetProcAddress(G_HANDLE, "QueryAccount");
	G_TestretmsgPtr				 = (FuncTestretmsgPtr)GetProcAddress(G_HANDLE, "xmalloc");

	RegisterUFXCallBack();
	return 0;
}

int main(int argc, char** argv){

	UFXInit();
	ConnectRet cr = UFXConnect("10.2.130.189:18801", 0);
	printf("%d, %s", cr.ErrorCode, cr.ErrorMsg);

	//UFXSubConnect();
	//msg = SubConnect();	
	//msg = Login();
	//
	//
	//Sleep(2000);

	////msg = LimitEntrust("1007", "1", "600000", "10072", "1", 12.880, 200);
	////printf("%s\n", msg);
	////HeartBeat();
	////QueryFundAsset("1007");
	////QueryPos("1007", "10072");
	////QueryEntrustByAcc("1007", "10072");

	//QueryAccount("1007", "10072");

	getchar();    

	return 0;
}