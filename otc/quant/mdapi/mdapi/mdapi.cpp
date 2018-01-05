// testTraderApi.cpp : 定义控制台应用程序的入口点。
//

#include "mdapi.h"
#include "mdspi.h"
//#include "stdio.h"
#include <iostream>

// UserApi对象
CThostFtdcMdApi* pUserApi;

// 配置参数
char FRONT_ADDR[]						= "tcp://61.140.230.188:41213";			// 前置地址
TThostFtdcBrokerIDType	BROKER_ID		= "2358";				// 经纪公司代码
TThostFtdcInvestorIDType INVESTOR_ID	= "999814006";			// 投资者代码
TThostFtdcPasswordType  PASSWORD		= "*****";			// 用户密码
mdcallback			G_MDCALLBACK = NULL;

char* flowfile = "ctpmd.con";

int iRequestID = 0;

int InitMdApi() {
	// 初始化UserApi
	pUserApi = CThostFtdcMdApi::CreateFtdcMdApi(flowfile, true);			// 创建UserApi
	CThostFtdcMdSpi* pUserSpi = new CMdSpi();
	pUserApi->RegisterSpi(pUserSpi);										// 注册事件类
	pUserApi->RegisterFront(FRONT_ADDR);									// connect
	pUserApi->Init();
	printf("初始化成功\n");
	return 0;
}

int __stdcall SubMd(char** ppInstrumentID, int num) {
	return pUserApi->SubscribeMarketData(ppInstrumentID, 1);
}

int __stdcall RegisterMdCallBack(mdcallback funcptr) {
	G_MDCALLBACK = funcptr;
	return 0;
}

int ret = InitMdApi();

void main(void){
	// login
	//CThostFtdcReqUserLoginField req;
	//memset(&req, 0, sizeof(req));
	//strcpy(req.BrokerID, BROKER_ID);
	//strcpy(req.UserID, INVESTOR_ID);
	//strcpy(req.Password, PASSWORD);
	//int ret = pUserApi->ReqUserLogin(&req, 101);
	//cout << ret << endl;

	// sub
	char *ppInstrumentID[] = { "au1806" };					// 行情订阅列表
	int iInstrumentID = 1;									// 行情订阅数量

	int ret = SubMd(ppInstrumentID, iInstrumentID);
	printf("%d\n",ret);
	pUserApi->Join();
	printf("Join 结束\n");
//	pUserApi->Release();
}


