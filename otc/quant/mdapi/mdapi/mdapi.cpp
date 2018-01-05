// testTraderApi.cpp : 定义控制台应用程序的入口点。
//

#include ".\ThostTraderApi\ThostFtdcMdApi.h"
#include "mdspi.h"
#include <iostream>
using namespace std;

// UserApi对象
CThostFtdcMdApi* pUserApi;

// 配置参数
char FRONT_ADDR[] = "tcp://61.140.230.188:41213";			// 前置地址
TThostFtdcBrokerIDType	BROKER_ID = "2358";				// 经纪公司代码
TThostFtdcInvestorIDType INVESTOR_ID = "999814006";			// 投资者代码
TThostFtdcPasswordType  PASSWORD = "*****";			// 用户密码

char *ppInstrumentID[] = {"au1806"};					// 行情订阅列表
int iInstrumentID = 1;									// 行情订阅数量
char* flowfile = "ctpmd.con";
// 请求编号
int iRequestID = 0;

void main(void)
{
	// 初始化UserApi
	pUserApi = CThostFtdcMdApi::CreateFtdcMdApi(flowfile, true);			// 创建UserApi
	CThostFtdcMdSpi* pUserSpi = new CMdSpi();
	pUserApi->RegisterSpi(pUserSpi);						// 注册事件类
	pUserApi->RegisterFront(FRONT_ADDR);					// connect
	pUserApi->Init();
	cout << "初始化成功" << endl;

	// login
	//CThostFtdcReqUserLoginField req;
	//memset(&req, 0, sizeof(req));
	//strcpy(req.BrokerID, BROKER_ID);
	//strcpy(req.UserID, INVESTOR_ID);
	//strcpy(req.Password, PASSWORD);
	//int ret = pUserApi->ReqUserLogin(&req, 101);
	//cout << ret << endl;

	// sub
	int ret = pUserApi->SubscribeMarketData(ppInstrumentID, 1);
	pUserApi->Join();
	cout << "Join 结束" << endl;
//	pUserApi->Release();
}