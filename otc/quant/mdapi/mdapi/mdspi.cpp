#include "mdspi.h"
#include "mdapi.h"
#include <iostream>
//using namespace std;

#pragma warning(disable : 4996)

// USER_API参数
extern CThostFtdcMdApi* pUserApi;
extern mdcallback G_MDCALLBACK;
// 配置参数
//extern char FRONT_ADDR[];		
//extern TThostFtdcBrokerIDType	BROKER_ID;
//extern TThostFtdcInvestorIDType INVESTOR_ID;
//extern TThostFtdcPasswordType	PASSWORD;
//extern char* ppInstrumentID[];	
//extern int iInstrumentID;

// 请求编号
//extern int iRequestID;

void CMdSpi::OnRspError(CThostFtdcRspInfoField *pRspInfo,
		int nRequestID, bool bIsLast){
	printf("%s%s\n", "--->>> ", "OnRspError" );
	IsErrorRspInfo(pRspInfo);
}

void CMdSpi::OnFrontDisconnected(int nReason){
	printf("%s%s\n", "--->>> ", "OnFrontDisconnected");
	printf("--->>> Reason = %d\n", nReason );
}
		
void CMdSpi::OnHeartBeatWarning(int nTimeLapse){
	printf("%s%s\n", "--->>> ", "OnHeartBeatWarning");
	printf("--->>> nTimerLapse = %d\n", nTimeLapse);
}

void CMdSpi::OnFrontConnected(){
	printf("--->>> OnFrontConnected \n");
	///用户登录请求
	//ReqUserLogin();
}

void CMdSpi::ReqUserLogin(TThostFtdcBrokerIDType broker_id, 
						  TThostFtdcInvestorIDType investor_id, 
						  TThostFtdcPasswordType password,
						   int iRequestID){
	CThostFtdcReqUserLoginField req;
	memset(&req, 0, sizeof(req));
	strcpy(req.BrokerID, broker_id);
	strcpy(req.UserID, investor_id);
	strcpy(req.Password, password);
	int iResult = pUserApi->ReqUserLogin(&req, iRequestID);
	printf("--->>> 发送用户登录请求: %s\n", ((iResult == 0) ? "成功" : "失败"));
}

void CMdSpi::OnRspUserLogin(CThostFtdcRspUserLoginField *pRspUserLogin,
		CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
	printf("--->>> OnRspUserLogin\n");
	if (bIsLast && !IsErrorRspInfo(pRspInfo)){
		///获取当前交易日
		printf("--->>> 获取当前交易日 = %s\n", pUserApi->GetTradingDay());
		// 请求订阅行情
		//SubscribeMarketData();	
	}
}

void CMdSpi::SubscribeMarketData(char** instrumentID, int num){
	int iResult = pUserApi->SubscribeMarketData(instrumentID, num);
	printf("--->>> 发送行情订阅请求: %s\n", ((iResult == 0) ? "成功" : "失败"));
}

void CMdSpi::OnRspSubMarketData(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
	printf("OnRspSubMarketData\n");
}

void CMdSpi::OnRspUnSubMarketData(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
	printf("OnRspUnSubMarketData\n");
}

void CMdSpi::OnRtnDepthMarketData(CThostFtdcDepthMarketDataField *pDepthMarketData){
	printf("OnRtnDepthMarketData lastprice: %f\n", pDepthMarketData->LastPrice);
	G_MDCALLBACK(pDepthMarketData);
}

bool CMdSpi::IsErrorRspInfo(CThostFtdcRspInfoField *pRspInfo){
	// 如果ErrorID != 0, 说明收到了错误的响应
	bool bResult = ((pRspInfo) && (pRspInfo->ErrorID != 0));
	if (bResult)
		printf("--->>> ErrorID=%d, ErrorMsg=%s", pRspInfo->ErrorID, pRspInfo->ErrorMsg);
	return bResult;
}


