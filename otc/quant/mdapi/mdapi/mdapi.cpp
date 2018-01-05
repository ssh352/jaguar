// testTraderApi.cpp : �������̨Ӧ�ó������ڵ㡣
//

#include ".\ThostTraderApi\ThostFtdcMdApi.h"
#include "mdspi.h"
#include <iostream>
using namespace std;

// UserApi����
CThostFtdcMdApi* pUserApi;

// ���ò���
char FRONT_ADDR[] = "tcp://61.140.230.188:41213";			// ǰ�õ�ַ
TThostFtdcBrokerIDType	BROKER_ID = "2358";				// ���͹�˾����
TThostFtdcInvestorIDType INVESTOR_ID = "999814006";			// Ͷ���ߴ���
TThostFtdcPasswordType  PASSWORD = "*****";			// �û�����

char *ppInstrumentID[] = {"au1806"};					// ���鶩���б�
int iInstrumentID = 1;									// ���鶩������
char* flowfile = "ctpmd.con";
// ������
int iRequestID = 0;

void main(void)
{
	// ��ʼ��UserApi
	pUserApi = CThostFtdcMdApi::CreateFtdcMdApi(flowfile, true);			// ����UserApi
	CThostFtdcMdSpi* pUserSpi = new CMdSpi();
	pUserApi->RegisterSpi(pUserSpi);						// ע���¼���
	pUserApi->RegisterFront(FRONT_ADDR);					// connect
	pUserApi->Init();
	cout << "��ʼ���ɹ�" << endl;

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
	cout << "Join ����" << endl;
//	pUserApi->Release();
}