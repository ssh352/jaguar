// testTraderApi.cpp : �������̨Ӧ�ó������ڵ㡣
//

#include ".\ThostTraderApi\ThostFtdcMdApi.h"
#include "MdSpi.h"
#include <iostream>
using namespace std;

// UserApi����
CThostFtdcMdApi* pUserApi;

// ���ò���
char FRONT_ADDR[] = "tcp://58.62.112.19:42213";		// ǰ�õ�ַ
TThostFtdcBrokerIDType	BROKER_ID = "2358";				// ���͹�˾����
TThostFtdcInvestorIDType INVESTOR_ID = "00092";			// Ͷ���ߴ���
TThostFtdcPasswordType  PASSWORD = "888888";			// �û�����
char *ppInstrumentID[] = {"au1803"};			// ���鶩���б�
int iInstrumentID = 2;									// ���鶩������

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
	
	
	pUserSpi->OnRspUserLogin();

	cout << "��ʼ���ɹ�" << endl;
	pUserApi->Join();
	cout << "Join ����" << endl;
//	pUserApi->Release();
}