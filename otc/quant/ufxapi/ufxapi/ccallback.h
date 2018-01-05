

#ifndef _CCALL_BACK_H
#define _CCALL_BACK_H

#include "t2sdk_interface.h"
#include <iostream>
using namespace std;


class CCallback : public CCallbackInterface{

	public:
		// ��ΪCCallbackInterface�����մ��������IKnown��������Ҫʵ��һ����3������
		unsigned long  FUNCTION_CALL_MODE QueryInterface(const char *iid, IKnown **ppv);
		unsigned long  FUNCTION_CALL_MODE AddRef();
		unsigned long  FUNCTION_CALL_MODE Release();

		// �����¼�����ʱ�Ļص�������ʵ��ʹ��ʱ���Ը�����Ҫ��ѡ��ʵ�֣����ڲ���Ҫ���¼��ص���������ֱ��return
		// Reserved?Ϊ����������Ϊ�Ժ���չ��׼����ʵ��ʱ��ֱ��return��return 0��
		void FUNCTION_CALL_MODE OnConnect(CConnectionInterface *lpConnection);
		void FUNCTION_CALL_MODE OnSafeConnect(CConnectionInterface *lpConnection);
		void FUNCTION_CALL_MODE OnRegister(CConnectionInterface *lpConnection);
		void FUNCTION_CALL_MODE OnClose(CConnectionInterface *lpConnection);
		void FUNCTION_CALL_MODE OnSent(CConnectionInterface *lpConnection, int hSend, void *reserved1, void *reserved2, int nQueuingData);
		void FUNCTION_CALL_MODE Reserved1(void *a, void *b, void *c, void *d);
		void FUNCTION_CALL_MODE Reserved2(void *a, void *b, void *c, void *d);
		int  FUNCTION_CALL_MODE Reserved3();
		void FUNCTION_CALL_MODE Reserved4();
		void FUNCTION_CALL_MODE Reserved5();
		void FUNCTION_CALL_MODE Reserved6();
		void FUNCTION_CALL_MODE Reserved7();
		void FUNCTION_CALL_MODE OnReceivedBiz(CConnectionInterface *lpConnection, int hSend, const void *lpUnPackerOrStr, int nResult);
		void FUNCTION_CALL_MODE OnReceivedBizEx(CConnectionInterface *lpConnection, int hSend, LPRET_DATA lpRetData, const void *lpUnpackerOrStr, int nResult);
		void FUNCTION_CALL_MODE OnReceivedBizMsg(CConnectionInterface *lpConnection, int hSend, IBizMessage* lpMsg);
};



#endif //_CCALL_BACK_H