



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
		QueryPosResp* Dataptr = (QueryPosResp*) ret->DataSet;
		printf("************** C UFXCallBack funcno: %d, ndatacount:%d, next %u **************\n", ret->FuncNo, ret->DataCount, Dataptr->nextdataptr);
		int nData = ret->DataCount;
		for( int i=0; i<nData; i++){
			printf("stock_code: %s, enable_amount: %d, last_price: %f, cost_price: %f, total_profit: %f, floating_profit: %f, accumulate_profit: %f\n", 
					Dataptr->stock_code, Dataptr->enable_amount, Dataptr->last_price, Dataptr->cost_price, Dataptr->total_profit, Dataptr->floating_profit, Dataptr->accumulate_profit);
			Dataptr = (QueryPosResp*)Dataptr->nextdataptr;
		}
	}else if(ret->FuncNo == 32001){
		// 委托结果查询解析
		QueryEntrustResp* Dataptr = (QueryEntrustResp*) ret->DataSet;
		printf("************** C UFXCallBack funcno: %d, ndatacount:%d, next %u **************\n", ret->FuncNo, ret->DataCount, Dataptr->nextentrustptr);
		int nData = ret->DataCount;
		for( int i=0; i<nData; i++){
			printf("stock_code: %s, entrust_date: %d, entrust_time: %d, entrust_amount: %d, entrust_price: %f\n", 
					Dataptr->stock_code, Dataptr->entrust_date, Dataptr->entrust_time, Dataptr->entrust_amount, Dataptr->entrust_price);
			Dataptr = (QueryEntrustResp*)Dataptr->nextentrustptr;
		}
	}else if(ret->FuncNo == 91001){
		// 按委托编号查询
		EntrustResp* Dataptr = (EntrustResp*) ret->DataSet;
		int entrust_no = Dataptr->entrust_no;
		//printf("entrust_no: %d\n", entrust_no);
		//QueryEntrustByEntrustNo("1007", "10072", entrust_no);
	}else if(ret->FuncNo == 34001){
		QueryAccountResp* Dataptr = (QueryAccountResp*) ret->DataSet;
		printf("account: %s enable_balance_t0: %f enable_balance_t1: %f current_balance: %f\n", 
			Dataptr->account_code, Dataptr->enable_balance_t0, Dataptr->enable_balance_t1, Dataptr->current_balance);
	}
	return 0;
}


int main(int argc, char** argv){

	int msg;

	//HINSTANCE handle = LoadLibrary("E:\\jaguar\\sandbox\\C++\\ufxapi\\x64\\Debug\\ufxapi.dll");
	//connptr = (FunConnectPtr)GetProcAddress(handle, "Connect");
	//connptr("10.2.130.189:18801");
	//FreeLibrary(handle);

	ufxcallback func = &UFXCallBack;
	RegisterCallBack(func);

	msg = Connect("10.2.130.189:18801", 1);
	if(msg == -1){
		printf("press enter exit app.");
		getchar();
		exit(-1);
	}
	
	msg = SubConnect();	
	msg = Login();
	
	
	Sleep(2000);

	msg = LimitEntrust("1007", "1", "600000", "10072", "1", 12.880, 200, "xxx");
	//printf("%s\n", msg);
	//HeartBeat();
	//QueryFundAsset("1007");
	//QueryPos("1007", "10072");
	//QueryEntrustByAcc("1007", "10072");

	//QueryAccount("1007", "10072");

	getchar();    

	return 0;
}