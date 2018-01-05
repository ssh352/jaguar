// testTraderApi.cpp : �������̨Ӧ�ó������ڵ㡣
//

#include "mdapi.h"
#include "mdspi.h"
//#include "stdio.h"
#include <iostream>

// UserApi����
CThostFtdcMdApi* pUserApi;

// ���ò���
char FRONT_ADDR[]						= "tcp://61.140.230.188:41213";			// ǰ�õ�ַ
TThostFtdcBrokerIDType	BROKER_ID		= "2358";				// ���͹�˾����
TThostFtdcInvestorIDType INVESTOR_ID	= "999814006";			// Ͷ���ߴ���
TThostFtdcPasswordType  PASSWORD		= "*****";			// �û�����
mdcallback			G_MDCALLBACK = NULL;

char* flowfile = "ctpmd.con";

int iRequestID = 0;

int InitMdApi() {
	// ��ʼ��UserApi
	pUserApi = CThostFtdcMdApi::CreateFtdcMdApi(flowfile, true);			// ����UserApi
	CThostFtdcMdSpi* pUserSpi = new CMdSpi();
	pUserApi->RegisterSpi(pUserSpi);										// ע���¼���
	pUserApi->RegisterFront(FRONT_ADDR);									// connect
	pUserApi->Init();
	printf("��ʼ���ɹ�\n");
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
	char *ppInstrumentID[] = { "au1806" };					// ���鶩���б�
	int iInstrumentID = 1;									// ���鶩������

	int ret = SubMd(ppInstrumentID, iInstrumentID);
	printf("%d\n",ret);
	pUserApi->Join();
	printf("Join ����\n");
//	pUserApi->Release();
}


