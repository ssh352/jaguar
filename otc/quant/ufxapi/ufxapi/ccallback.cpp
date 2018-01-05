
#include <vector>
#include <list>
#include "ccallback.h"
#include "ufxutil.h"

using namespace std;

extern string		gUserToken;
extern ufxcallback	G_GOCALLBACK;


typedef void(*HandlePushMsg)(Result*, void**, IF2UnPacker*);
extern map<int, HandlePushMsg> G_HANDLEMSGFUNC;


unsigned long CCallback::QueryInterface(const char *iid, IKnown **ppv){return 0;}

unsigned long CCallback::AddRef(){return 0;}

unsigned long CCallback::Release(){return 0;}

void CCallback::OnConnect(CConnectionInterface *lpConnection){
    //cout << "OnConnect: successfully connected." << endl;
}

void CCallback::OnSafeConnect(CConnectionInterface *lpConnection){
    //cout << "OnSafeConnect: successfully connected." << endl;
}

void CCallback::OnRegister(CConnectionInterface *lpConnection){
    //cout << "OnRegister: successfully registered." << endl;
}

void CCallback::OnClose(CConnectionInterface *lpConnection){
    cout << "OnClose: sdk connection closed." << endl;
}

void CCallback::OnSent(CConnectionInterface *lpConnection, int hSend, void *reserved1, void *reserved2, int nQueuingData){
    cout << "OnSend: hSend(" << hSend << ") send successed, queuingData(" << nQueuingData << ")." << endl;
}

void CCallback::Reserved1(void *a, void *b, void *c, void *d){}

void CCallback::Reserved2(void *a, void *b, void *c, void *d){}

int  CCallback::Reserved3(){ return 0;}

void CCallback::Reserved4(){}

void CCallback::Reserved5(){}

void CCallback::Reserved6(){}

void CCallback::Reserved7(){}

void CCallback::OnReceivedBiz(CConnectionInterface *lpConnection, int hSend, const void *lpUnPackerOrStr, int nResult){}

void CCallback::OnReceivedBizEx(CConnectionInterface *lpConnection, int hSend, LPRET_DATA lpRetData, const void *lpUnpackerOrStr, int nResult){}

void CCallback::OnReceivedBizMsg(CConnectionInterface *lpConnection, int hSend, IBizMessage* lpMsg){
	//printf("CCallback::OnReceivedBizMsg\n");
	Result *ret			= (Result *)malloc(sizeof(Result));
	void **resp			= NULL;
	int nData			= 0;
	if (lpMsg == NULL){
		ret->ErrorCode		= -1;
		ret->ErrorMsg		= "CCallback::OnReceivedBizMsg, Message package's ptr is null";
	}else{
		int iFuncNo			= lpMsg->GetFunction();
		int iPacketType		= lpMsg->GetPacketType();

		ret->FuncNo			= iFuncNo;
		ret->PacketType		= iPacketType;

		int iReturnCode		= lpMsg->GetReturnCode();
		ret->ReturnCode		= iReturnCode;
	
		if (iReturnCode == 1 || iReturnCode == -1){
			ret->errorNo	= lpMsg->GetErrorNo();
			ret->errorInfo	= lpMsg->GetErrorInfo();
		}else{
			int iLen = 0;
			const void* responseBuffer		= lpMsg->GetContent(iLen);
			IF2UnPacker* responseUnPacker	= NewUnPacker((void *)responseBuffer,iLen);
			AddRef();
			//ShowPacket(responseUnPacker);

			responseUnPacker->SetCurrentDatasetByIndex(0);				//第1个结果集是消息头,肯定存在
			responseUnPacker->First();
			ret->ErrorCode		= responseUnPacker->GetInt("ErrorCode");		//接口调用结果,0表示成功,非0表示失败
			ret->ErrorMsg		= responseUnPacker->GetStr("ErrorMsg");			//失败时的错误信息,成功时为空
			ret->MsgDetail		= responseUnPacker->GetStr("MsgDetail");		//详细错误信息,建议取这个字段,方便排查问题
			ret->DataCount		= responseUnPacker->GetInt("DataCount");		//表示第2个结果集的记录数
			nData				= ret->DataCount;
			resp				= new void*[nData];
			if (responseUnPacker->GetDatasetCount() > 1){
				responseUnPacker->SetCurrentDatasetByIndex(1);
				responseUnPacker->First();

				if (iFuncNo == 10001){			//登录
					gUserToken = responseUnPacker->GetStr("user_token");
				}
				// Deal funcno response
				if(G_HANDLEMSGFUNC.find(iFuncNo) != G_HANDLEMSGFUNC.end()){
					G_HANDLEMSGFUNC[iFuncNo](ret, resp, responseUnPacker);
				}else{
					nData = 0;
					printf("Error: Please resgister funcno: '%d' response handle func\n", iFuncNo);	
				}	
			}
			responseUnPacker->Release();
		}
	}

	if(G_GOCALLBACK != NULL){
		G_GOCALLBACK(ret);
	}
	for(int i=0; i<nData; i++){
		void* ptr = resp[i];
		//printf("iFuncNo: %d, free ptr---%d: %u\n", ret->FuncNo, i, ptr);
		free(ptr);
	}
	if(resp != NULL){
		delete[] resp;
	}
	free(ret);
}
