



#include <iostream>
#include <string>
#include <array>
#include "csubcallback.h"
#include "ufxutil.h"
using namespace std;

extern ufxcallback	G_GOCALLBACK;

unsigned long CSubCallback::QueryInterface(const char *iid, IKnown **ppv){return 0;}

void CSubCallback::OnRecvTickMsg(CSubscribeInterface *lpSub,int subscribeIndex,const char* TickMsgInfo){}

unsigned long CSubCallback::AddRef(){return 0;}

unsigned long CSubCallback::Release(){return 0;}


void CSubCallback::OnReceived(CSubscribeInterface *lpSub, int subscribeIndex, const void *lpData, int nLength,LPSUBSCRIBE_RECVDATA lpRecvData){
	/*printf("CSubCallback::OnReceived\n");*/
	Result *ret				= (Result *)malloc(sizeof(Result));
	void *resp				= NULL;
	ret->FuncNo				= -999999;
	string sMsgType;
	IF2UnPacker* lpUnPack	= NULL;

	if (lpData == NULL){
		ret->ErrorCode		= -1;
		ret->ErrorMsg		= "CSubCallback::OnReceived, Message push's ptr is null";
	}else{

		lpUnPack = NewUnPacker((void*)lpData,nLength);
		AddRef();

		lpUnPack->SetCurrentDatasetByIndex(0);
		lpUnPack->First();
		
		
		sMsgType			= lpUnPack->GetStr("msgtype");  //主推消息类型
		ret->MsgType		= sMsgType.c_str();
		string sMarketNo	= lpUnPack->GetStr("market_no");
		string sOperatorNo  = lpUnPack->GetStr("operator_no");
		
		array<string, 6> entruststatus = { "a", "b", "c", "d", "e", "f" };
		if (find(begin(entruststatus ), end(entruststatus), sMsgType) != std::end(entruststatus)){
			//ShowPacket(lpUnPack);
			UFX_HANDLE_ENTRUST_PUSH(ret, resp, lpUnPack);
		}else if (sMsgType == "g"){
			//ShowPacket(lpUnPack);
			UFX_HANDLE_DEAL_PUSH(ret, resp, lpUnPack);
		}
	}
	
	if(G_GOCALLBACK != NULL){
		G_GOCALLBACK(ret);
	}

	if(resp != NULL){
		free(resp);
	}
	free(ret);

	if (lpUnPack != NULL){
		lpUnPack->Release();
	}
}

