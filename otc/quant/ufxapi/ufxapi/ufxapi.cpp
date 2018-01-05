/*
 *   UFX�ӿ��첽ʹ��ʾ��(C++��)
*
*    UFX��ʹ�ù�����ʵ����ʹ��t2sdk��������UFX�������������ӡ����Ͳ�����ҵ����Ϣ�Ĺ��̡�
*    ���У�T2SDK��������ҵ���޹صģ���ʹ�ÿ��Բο���T2SDK �ⲿ�濪��ָ��.docx��
*               ҵ����Ϣ�Ķ�������ҵ����صģ�ÿ���ӿڶ����Լ��Ķ��壬���Բο�������Ͷ�ʹ���ϵͳO3.2_�ܱ߽ӿڹ淶_x.x.x.x.xls��
*/
#include <string>
#include <iostream>
#include <iomanip>
#include <sstream>
#include <cassert>
#include <msgpack.hpp>

#include <list>
#include <map>
#include <vector>
#include <sstream>
#include <string>
#include "ufxapi.h"
#include "csubcallback.h"
#include "ccallback.h"
#include "ufxutil.h"


typedef void(*HandlePushMsg)(Result*, void**, IF2UnPacker*);


int InitFuncMap();
int ReadConfig();

string gOperatorNo  = "";			// ����Ա
string gPassword    = "";           // ����Ա����
string gUserToken;

string G_MAC				= "";
string G_OP_STATION			= "";
string G_IP_ADDRESS			= "";
string G_AUTHORIZATION_ID	= "";
int G_TIMEOUT				= 5000;

map<int, CSubscribeParamInterface*> gAllSubscribeParam;
map<int, HandlePushMsg> G_HANDLEMSGFUNC;


CConnectionInterface*	G_CONNECTION1 = NULL;
CConnectionInterface*	G_CONNECTION2 = NULL;
CCallback				G_CALLBACK1;
CCallback				G_CALLBACK2;
CSubCallback			G_SUBCALLBACK;
CSubscribeInterface*	G_SUBSCRIBE = NULL;
ufxcallback				G_GOCALLBACK = NULL;

int _retcode1 = InitFuncMap();
int _retcode2 = ReadConfig();



int ReadConfig(){
	CConfigInterface * lpConfig = NewConfig();
	lpConfig->AddRef();
	lpConfig->Load("./conf/ufx/ufxapi.ini");
	
	gOperatorNo			= lpConfig->GetString("ufxpai", "operate_no", "");
	gPassword			= lpConfig->GetString("ufxpai", "operate_pwd", "");
	G_MAC				= lpConfig->GetString("ufxpai", "mac", "");
	G_OP_STATION		= lpConfig->GetString("ufxpai", "op_station", "");
	G_IP_ADDRESS		= lpConfig->GetString("ufxpai", "ip_address", "");
	G_AUTHORIZATION_ID  = lpConfig->GetString("ufxpai", "authorization_id", "");
	G_TIMEOUT			= lpConfig->GetInt("ufxpai", "timeout", 5000);
	lpConfig->Release();
	return 0;
}

int InitFuncMap(){
	 G_HANDLEMSGFUNC[10001] = (HandlePushMsg)(&UFX_10001_unPacker);	
	 G_HANDLEMSGFUNC[91001] = (HandlePushMsg)(&UFX_91001_unPacker);	
	 G_HANDLEMSGFUNC[91101] = (HandlePushMsg)(&UFX_91101_unPacker);
	 G_HANDLEMSGFUNC[35003] = (HandlePushMsg)(&UFX_35003_unPacker);	
	 G_HANDLEMSGFUNC[31001] = (HandlePushMsg)(&UFX_31001_unPacker);	
	 G_HANDLEMSGFUNC[32001] = (HandlePushMsg)(&UFX_32001_unPacker);	
	 G_HANDLEMSGFUNC[34001] = (HandlePushMsg)(&UFX_34001_unPacker);
	 return 0;
}

void HandleUfxPushMsg(int FuncNO, Result* ret, void* resp, IF2UnPacker* responseUnPacker){

}

int CallService(CConnectionInterface* connection, IBizMessage* requestBizMessage){
	return connection->SendBizMsg(requestBizMessage,1);
}

int __stdcall RegisterCallBack(ufxcallback funcptr){
	G_GOCALLBACK = funcptr;
	return 0;
}


/****************************** ������������ ******************************/
// success: 0
// failed: -1
// sync : 0 ��ʾͬ�����ӣ� 1 ��ʾ�첽����
int __stdcall Connect(const char* serverAddr, int sync = 0){

	//��������	
	//CConfigInterface��CConnectionInterface��ʹ�ÿ��Բο���T2SDK �ⲿ�濪��ָ��.docx��
	//����T2SDK���ö��󣬲�����UFX��������ַ�Լ�License�ļ�
	
	CConfigInterface * lpConfig = NewConfig();
	lpConfig->AddRef();

	lpConfig->SetString("t2sdk", "servers", serverAddr);
	lpConfig->SetString("t2sdk", "license_file", "license.dat");

	G_CONNECTION1 = NewConnection(lpConfig);
	G_CONNECTION1->AddRef();
	int ret = -1;
	if(sync == 0){
		ret = G_CONNECTION1->Create2BizMsg(NULL);
	}else{
		//�������Ӷ��󣬲�ע��ص�
		ret = G_CONNECTION1->Create2BizMsg(&G_CALLBACK1);
	}

	if (ret != 0) {
		lpConfig->Release();
		string msg = "[Connect] Create2BizMsg Error, errormsg�� " + string(G_CONNECTION1->GetErrorMsg(ret));
		cout << msg << endl;
		G_CONNECTION1->Release();
		G_CONNECTION1 = NULL;
		return ret;
	}

	//����UFX������������G_TIMEOUTΪ��ʱ��������λ����
	ret = G_CONNECTION1->Connect(G_TIMEOUT); 
	if (ret != 0){
		lpConfig->Release();
		cout << G_CONNECTION1->GetErrorMsg(ret) << endl;
		string msg = "[Connect] Connect Error, errormsg�� " + string(G_CONNECTION1->GetErrorMsg(ret));
		cout << msg << endl;
		G_CONNECTION1->Release();
		G_CONNECTION1 = NULL;
		return ret;
	}
	return 0;
}

///****************************** ������������ ******************************/
int __stdcall SubConnect(){
	////��������
	
	//ͨ��T2SDK����������������ȡһ���µ�CConfig����ָ��
	//�˶����ڴ������Ӷ���ʱ�����ݣ��������������������Ӷ���ĸ������ԣ����������IP��ַ����ȫģʽ�ȣ�
	//ֵ��ע����ǣ��������ö�������������Ϣʱ��������Ϣ�ȿ��Դ�ini�ļ������룬
	//Ҳ�����ڳ���������趨��������2�ߵĻ�ϣ������ͬһ���������費ͬ��ֵ���������һ������Ϊ׼
	CConfigInterface * lpConfig = NewConfig();
	lpConfig->AddRef();
	lpConfig->Load("./conf/ufx/subscriber.ini");

	//�������Ӷ��󣬲�ע��ص�(���ӵ�ַ�Ͷ˿���subscriber.ini�ļ�������)
	G_CONNECTION2 = NewConnection(lpConfig);
	G_CONNECTION2->AddRef();
	int ret = G_CONNECTION2->Create2BizMsg(&G_CALLBACK2);
	if (ret != 0){
		string msg = "[SubConnect] Create2BizMsg Error, errormsg: " + string(G_CONNECTION2->GetErrorMsg(ret));
		cout << msg << endl;
		lpConfig->Release();
		G_CONNECTION2->Release();
		G_CONNECTION2 = NULL;
		return -1;
	}

	//����UFX������������3000Ϊ��ʱ��������λ����
	ret = G_CONNECTION2->Connect(3000); 
	if (ret != 0){
		string msg = "[SubConnect] Connect Error, errormsg: " + string(G_CONNECTION2->GetErrorMsg(ret));
		cout << msg << endl;
		lpConfig->Release();
		G_CONNECTION2->Release();
		G_CONNECTION2 = NULL;
		return -1;
	}

	//�������Ķ���
	char* bizName = (char*)lpConfig->GetString("subcribe","biz_name","");
	G_SUBSCRIBE = G_CONNECTION2->NewSubscriber(&G_SUBCALLBACK ,bizName, 5000);
	if (G_SUBSCRIBE == NULL){
		string msg = "[SubConnect] NewSubscribe Error, errormsg: " + string(G_CONNECTION2->GetMCLastError());
		cout << msg << endl;
		lpConfig->Release();
		G_CONNECTION2->Release();
		G_CONNECTION2 = NULL;
		return  -1;
	}
	G_SUBSCRIBE->AddRef();

	/****************************** ��ȡ���Ĳ��� ******************************/
	CSubscribeParamInterface* lpSubscribeParam = NewSubscribeParam();
	lpSubscribeParam->AddRef();

	char* topicName = (char*)lpConfig->GetString("subcribe","topic_name","");//��������
	lpSubscribeParam->SetTopicName(topicName);

	char* isFromNow = (char*)lpConfig->GetString("subcribe","is_rebulid","");//�Ƿ�ȱ
	if (strcmp(isFromNow,"true")==0){
		lpSubscribeParam->SetFromNow(true);
	}else{
		lpSubscribeParam->SetFromNow(false);
	}

	char* isReplace = (char*)lpConfig->GetString("subcribe","is_replace","");//�Ƿ񸲸�
	if (strcmp(isReplace,"true")==0){
		lpSubscribeParam->SetReplace(true);
	}else{
		lpSubscribeParam->SetReplace(false);
	}

	char* lpApp = "lixuebin";
	lpSubscribeParam->SetAppData(lpApp,8);//��Ӹ�������

	//��ӹ����ֶ�
	int nCount = lpConfig->GetInt("subcribe", "filter_count", 0);
	for (int i=1;i<=nCount;i++){
		char lName[128]={0};
		sprintf(lName,"filter_name%d",i);
		char* filterName = (char*)lpConfig->GetString("subcribe",lName,"");
		char lValue[128]={0};
		sprintf(lValue,"filter_value%d",i);
		char* filterValue = (char*)lpConfig->GetString("subcribe",lValue,"");
		lpSubscribeParam->SetFilter(filterName,filterValue);
	}

	//��ӷ���Ƶ��
	lpSubscribeParam->SetSendInterval(lpConfig->GetInt("subcribe","send_interval",0));

	//��ӷ����ֶ�
	nCount = lpConfig->GetInt("subcribe","return_count",0);
	for (int k=1;k<=nCount;k++){
		char lName[128]={0};
		sprintf(lName,"return_filed%d",k);
		char* filedName = (char*)lpConfig->GetString("subcribe",lName,"");
		lpSubscribeParam->SetReturnFiled(filedName);
	}

	//����һ��ҵ���
	IF2Packer* pack = NewPacker(2);
	pack->AddRef();
	pack->BeginPack();
	pack->AddField("login_operator_no");
	pack->AddField("password");
	pack->AddStr(gOperatorNo.c_str());//��������Ĳ���Ա
	pack->AddStr(gPassword.c_str());//��������Ĳ���Ա����
	pack->EndPack();
	IF2UnPacker* lpBack = NULL;

	int iRet = G_SUBSCRIBE->SubscribeTopic(lpSubscribeParam, 5000, &lpBack, pack);
	if (iRet <= 0){
		if (lpBack != NULL) ShowPacket(lpBack);
		string msg = "[SubConnect] SubscribeTopic Error, errormsg: " + string(G_CONNECTION2->GetMCLastError());
		cout << msg << endl;
		pack->FreeMem(pack->GetPackBuf());
		pack->Release();
		lpSubscribeParam->Release();
		lpConfig->Release();
		G_CONNECTION2->Release();
		return -1;
	}

	int subscribeIndex = iRet;
	gAllSubscribeParam[subscribeIndex] = lpSubscribeParam;//���浽map�У������Ժ��ȡ������

	pack->FreeMem(pack->GetPackBuf());
	pack->Release();
	lpConfig->Release();

	return 0;
}

Result* __stdcall LoginSync(){
	Result *ret = new Result;
	if(G_CONNECTION1 == NULL){
		string msg = "[LoginSync] G_CONNECTION1 is NULL, errormsg�� please connect first!";
		ret->ErrorCode = -1;
		ret->ErrorMsg = msg.c_str();
		return ret;
	}else if(G_CONNECTION2 == NULL){
		string msg = "[LoginSync] G_CONNECTION2 is NULL, errormsg�� please subconnect first!";
		cout << msg << endl;
		return ret;
	}
	IBizMessage* lpBizMessage = NewBizMessage();
	lpBizMessage->AddRef();
	//��¼���ܺţ�10001
	lpBizMessage->SetFunction(10001);
	lpBizMessage->SetPacketType(REQUEST_PACKET);

	IF2Packer* requestPacker = NewPacker(2);
	requestPacker->AddRef();
	requestPacker->BeginPack();
	requestPacker->AddField("operator_no",     'S',16, 0);
	requestPacker->AddField("password",        'S',32, 0);
	requestPacker->AddField("mac_address",     'S',255,0);
	requestPacker->AddField("op_station",      'S',255,0);
	requestPacker->AddField("ip_address",      'S',32, 0);
	requestPacker->AddField("authorization_id",'S',64, 0);
	requestPacker->AddStr(gOperatorNo.c_str());
	requestPacker->AddStr(gPassword.c_str());
	requestPacker->AddStr(G_MAC.c_str());
	requestPacker->AddStr(G_OP_STATION.c_str());
	requestPacker->AddStr(G_IP_ADDRESS.c_str());
	requestPacker->AddStr(G_AUTHORIZATION_ID.c_str());
	requestPacker->EndPack();

	IF2Packer* lpPacker = requestPacker;
	lpBizMessage->SetContent(lpPacker->GetPackBuf(),lpPacker->GetPackLen());

	int iret = G_CONNECTION1->SendBizMsg(lpBizMessage);
	if(iret<0){
		ret->errorNo = iret;
		ret->errorInfo = G_CONNECTION1->GetErrorMsg(iret);
		return ret;
	}else{
		IBizMessage* lpBizMessageRecv = NULL;
		iret = G_CONNECTION1->RecvBizMsg(iret, &lpBizMessageRecv);
		if(iret == 0){
			int iReturnCode = lpBizMessageRecv->GetReturnCode();
			if(iReturnCode==1 || iReturnCode==-1) {
				//����
				ret->errorNo = lpBizMessageRecv->GetErrorNo();
				ret->errorInfo = lpBizMessageRecv->GetErrorInfo();
				return ret;
			}else if(iReturnCode==0) {
				// ��ȷ
				int iLen = 0;
				const void * lpBuffer = lpBizMessageRecv->GetContent(iLen);
				IF2UnPacker * lpUnPacker = NewUnPacker((void *)lpBuffer,iLen);
				//G_HANDLEMSGFUNC[10001](ret, resp, lpUnPacker);
			}
		}	
	}
	return ret;
}

//int __stdcall ReturnStruct(Result2* ret){
//	ret->errorInfo = "success";
//	ret->FuncNo = 999;
//	return 0;
//}

void print(std::string const& buf) {
    for (std::string::const_iterator it = buf.begin(), end = buf.end();
         it != end;
         ++it) {
        std::cout
            << std::setw(2)
            << std::hex
            << std::setfill('0')
            << (static_cast<int>(*it) & 0xff)
            << ' ';
    }
    std::cout << std::dec << std::endl;
}

int __stdcall Login(){
	if(G_CONNECTION1 == NULL){
		string msg = "[Login] G_CONNECTION1 is NULL, errormsg�� please connect first!";
		cout << msg << endl;
		return -1;
	}else if(G_CONNECTION2 == NULL){
		string msg = "[Login] G_CONNECTION2 is NULL, errormsg�� please subconnect first!";
		cout << msg << endl;
		return -1;
	}

	IBizMessage* lpBizMessage = NewBizMessage();
	lpBizMessage->AddRef();
	//��¼���ܺţ�10001
	lpBizMessage->SetFunction(10001);
	lpBizMessage->SetPacketType(REQUEST_PACKET);

	IF2Packer* requestPacker = NewPacker(2);
	requestPacker->AddRef();
	requestPacker->BeginPack();
	requestPacker->AddField("operator_no",     'S',16, 0);
	requestPacker->AddField("password",        'S',32, 0);
	requestPacker->AddField("mac_address",     'S',255,0);
	requestPacker->AddField("op_station",      'S',255,0);
	requestPacker->AddField("ip_address",      'S',32, 0);
	requestPacker->AddField("authorization_id",'S',64, 0);
	requestPacker->AddStr(gOperatorNo.c_str());
	requestPacker->AddStr(gPassword.c_str());
	requestPacker->AddStr(G_MAC.c_str());
	requestPacker->AddStr(G_OP_STATION.c_str());
	requestPacker->AddStr(G_IP_ADDRESS.c_str());
	requestPacker->AddStr(G_AUTHORIZATION_ID.c_str());
	requestPacker->EndPack();

	IF2Packer* lpPacker = requestPacker;
	lpBizMessage->SetContent(lpPacker->GetPackBuf(),lpPacker->GetPackLen());
	int ret = CallService(G_CONNECTION1, lpBizMessage);
	lpPacker->FreeMem(lpPacker->GetPackBuf());
	lpPacker->Release();
	lpBizMessage->Release();				
	
	if(ret<=0){
		string msg = "[Login] Login Error, errormsg�� " + string(G_CONNECTION1->GetErrorMsg(ret));
		cout << msg << endl;
		return -1;
	}
	return 0;
}


int __stdcall LimitEntrust(const char* account_code, const char* market_no, const char* stock_code, 
								   const char* combi_no, const char* BS, double price, int vol, const char* third_reff){
	if(G_CONNECTION1 == NULL){
		string msg = "[LimitEntrust] G_CONNECTION1 judge Error, errormsg�� please connect first!";
		cout << msg << endl;
		return -1;
	}else if(G_CONNECTION2 == NULL){
		string msg = "[LimitEntrust] G_CONNECTION2 is NULL, errormsg�� please subconnect first!";
		cout << msg << endl;
		return -1;
	}

	IBizMessage* lpBizMessage = NewBizMessage();
	lpBizMessage->AddRef();
	//ί�й��ܺţ�91001���ӿڹ��ܺż��������������������Բο�������Ͷ�ʹ���ϵͳO3.2_�ܱ߽ӿڹ淶_x.x.x.x.xls��
	lpBizMessage->SetFunction(91001);
	lpBizMessage->SetPacketType(REQUEST_PACKET);
	
	IF2Packer* requestPacker = NewPacker(2);
	requestPacker->AddRef();
	requestPacker->BeginPack();
	requestPacker->AddField("user_token",        'S',512,0);
	requestPacker->AddField("batch_no",          'I',8,0);
	requestPacker->AddField("account_code",      'S',32,0);
	requestPacker->AddField("combi_no"  ,        'S',8,0); 
	requestPacker->AddField("market_no" ,        'S',3,0);
	requestPacker->AddField("stock_code",        'S',16,0);
	requestPacker->AddField("entrust_direction", 'S',1,0);
	requestPacker->AddField("price_type",        'S',1,0);
	requestPacker->AddField("entrust_price",     'F',9,3);
	requestPacker->AddField("entrust_amount",    'F',16,2);
	requestPacker->AddField("third_reff",		 'S',128,0);
	requestPacker->AddStr(gUserToken.c_str());
	requestPacker->AddInt(0);
	requestPacker->AddStr(account_code);
	requestPacker->AddStr(combi_no);
	requestPacker->AddStr(market_no);
	requestPacker->AddStr(stock_code);
	requestPacker->AddStr(BS);
	requestPacker->AddStr("0");                  //�޼�
	requestPacker->AddDouble(price);
	requestPacker->AddDouble(vol);
	requestPacker->AddStr(third_reff);
	
	requestPacker->EndPack();
	IF2Packer* lpPacker = requestPacker;

	lpBizMessage->SetContent(lpPacker->GetPackBuf(),lpPacker->GetPackLen());
	int ret = CallService(G_CONNECTION1, lpBizMessage);
	lpPacker->FreeMem(lpPacker->GetPackBuf());
	lpPacker->Release();
	lpBizMessage->Release();

	if(ret<=0){
		string msg = "[LimitEntrust] LimitEntrust Error, errormsg�� " + string(G_CONNECTION1->GetErrorMsg(ret));
		cout << msg << endl;
		return -1;
	}
	
	return 0;
}

int __stdcall Withdraw(int entrustno){
	IBizMessage* lpBizMessage = NewBizMessage();
	lpBizMessage->AddRef();
	//�������ܺţ�91101
	lpBizMessage->SetFunction(91101);
	lpBizMessage->SetPacketType(REQUEST_PACKET);

	IF2Packer* requestPacker = NewPacker(2);
	requestPacker->AddRef();
	requestPacker->BeginPack();
	requestPacker->AddField("user_token",'S',512,0);
	requestPacker->AddField("entrust_no",'I');
	requestPacker->AddStr(gUserToken.c_str());
	requestPacker->AddInt(entrustno);
	requestPacker->EndPack();

	IF2Packer* lpPacker = requestPacker;
	lpBizMessage->SetContent(lpPacker->GetPackBuf(),lpPacker->GetPackLen());
	int ret = CallService(G_CONNECTION1, lpBizMessage);
	lpPacker->FreeMem(lpPacker->GetPackBuf());
	lpPacker->Release();
	lpBizMessage->Release();

	if (ret <= 0){
		string msg = "[Withdraw] Withdraw Error, errormsg: " + string(G_CONNECTION1->GetErrorMsg(ret));
		cout << msg << endl;
		return -1;
	}
	return 0;
}

int __stdcall HeartBeat(){
	IBizMessage* lpBizMessage = NewBizMessage();
	lpBizMessage->AddRef();
	//������10000
	lpBizMessage->SetFunction(10000);
	lpBizMessage->SetPacketType(REQUEST_PACKET);

	IF2Packer* requestPacker = NewPacker(2);
	requestPacker->AddRef();
	requestPacker->BeginPack();
	requestPacker->AddField("user_token",'S',512,0);
	requestPacker->AddStr(gUserToken.c_str());
	requestPacker->EndPack();

	IF2Packer* lpPacker = requestPacker;
	lpBizMessage->SetContent(lpPacker->GetPackBuf(),lpPacker->GetPackLen());
	int ret = CallService(G_CONNECTION1, lpBizMessage);
	lpPacker->FreeMem(lpPacker->GetPackBuf());
	lpPacker->Release();
	lpBizMessage->Release();

	if (ret <= 0){
		string msg = "[HeartBeat] HeartBeat Error, errormsg: " + string(G_CONNECTION1->GetErrorMsg(ret));
		cout << msg << endl;
		return -1;
	}
	return 0;
}


int __stdcall ExitUFX(){
	IBizMessage* lpBizMessage = NewBizMessage();
	lpBizMessage->AddRef();
	//�˳���½��10002
	lpBizMessage->SetFunction(10002);
	lpBizMessage->SetPacketType(REQUEST_PACKET);

	IF2Packer* requestPacker = NewPacker(2);
	requestPacker->AddRef();
	requestPacker->BeginPack();
	requestPacker->AddField("user_token",'S',512,0);
	requestPacker->AddStr(gUserToken.c_str());
	requestPacker->EndPack();

	IF2Packer* lpPacker = requestPacker;
	lpBizMessage->SetContent(lpPacker->GetPackBuf(),lpPacker->GetPackLen());
	int ret = CallService(G_CONNECTION1, lpBizMessage);
	lpPacker->FreeMem(lpPacker->GetPackBuf());
	lpPacker->Release();
	lpBizMessage->Release();

	if (ret <= 0){
		string msg = "[ExitUFX] ExitUFX Error, errormsg: " + string(G_CONNECTION1->GetErrorMsg(ret));
		cout << msg << endl;
		return -1;
	}
	return 0;
}


int __stdcall QueryFundAsset(const char* account){
	IBizMessage* lpBizMessage = NewBizMessage();
	lpBizMessage->AddRef();
	//�˻��ʲ���ѯ��35003
	lpBizMessage->SetFunction(35003);
	lpBizMessage->SetPacketType(REQUEST_PACKET);

	IF2Packer* requestPacker = NewPacker(2);
	requestPacker->AddRef();
	requestPacker->BeginPack();
	requestPacker->AddField("user_token",	'S',512,0);
	requestPacker->AddField("account_code", 'S',32,0);
	requestPacker->AddStr(gUserToken.c_str());
	requestPacker->AddStr(account);
	requestPacker->EndPack();

	IF2Packer* lpPacker = requestPacker;
	lpBizMessage->SetContent(lpPacker->GetPackBuf(),lpPacker->GetPackLen());
	int ret = CallService(G_CONNECTION1, lpBizMessage);
	lpPacker->FreeMem(lpPacker->GetPackBuf());
	lpPacker->Release();
	lpBizMessage->Release();

	if (ret <= 0){
		string msg = "[QueryAccount] QueryAccount Error, errormsg: " + string(G_CONNECTION1->GetErrorMsg(ret));
		cout << msg << endl;
		return -1;
	}
	return 0;
}


int __stdcall QueryPos(const char* account, const char* combi_bo){
	IBizMessage* lpBizMessage = NewBizMessage();
	lpBizMessage->AddRef();
	//�˻��ʲ���ѯ��31001
	lpBizMessage->SetFunction(31001);
	lpBizMessage->SetPacketType(REQUEST_PACKET);

	IF2Packer* requestPacker = NewPacker(2);
	requestPacker->AddRef();
	requestPacker->BeginPack();
	requestPacker->AddField("user_token",	'S',512,0);
	requestPacker->AddField("account_code", 'S',32,0);
	requestPacker->AddField("combi_no"  ,   'S',8,0); 
	requestPacker->AddStr(gUserToken.c_str());
	requestPacker->AddStr(account);
	requestPacker->AddStr(combi_bo);
	requestPacker->EndPack();

	IF2Packer* lpPacker = requestPacker;
	lpBizMessage->SetContent(lpPacker->GetPackBuf(),lpPacker->GetPackLen());
	int ret = CallService(G_CONNECTION1, lpBizMessage);
	lpPacker->FreeMem(lpPacker->GetPackBuf());
	lpPacker->Release();
	lpBizMessage->Release();

	if (ret <= 0){
		string msg = "[QueryPos] QueryPos Error, errormsg: " + string(G_CONNECTION1->GetErrorMsg(ret));
		cout << msg << endl;
		return -1;
	}
	return 0;
}

int __stdcall QueryEntrustByAcc(const char* account, const char* combi_no){
	IBizMessage* lpBizMessage = NewBizMessage();
	lpBizMessage->AddRef();
	//֤ȯί�в�ѯ��32001
	lpBizMessage->SetFunction(32001);
	lpBizMessage->SetPacketType(REQUEST_PACKET);

	IF2Packer* requestPacker = NewPacker(2);
	requestPacker->AddRef();
	requestPacker->BeginPack();
	requestPacker->AddField("user_token",	'S',512,0);
	requestPacker->AddField("account_code", 'S',32,0);
	requestPacker->AddField("combi_no"  ,   'S',8,0); 
	requestPacker->AddStr(gUserToken.c_str());
	requestPacker->AddStr(account);
	requestPacker->AddStr(combi_no);
	requestPacker->EndPack();

	IF2Packer* lpPacker = requestPacker;
	lpBizMessage->SetContent(lpPacker->GetPackBuf(),lpPacker->GetPackLen());
	int ret = CallService(G_CONNECTION1, lpBizMessage);
	lpPacker->FreeMem(lpPacker->GetPackBuf());
	lpPacker->Release();
	lpBizMessage->Release();

	if (ret <= 0){
		string msg = "[QueryPos] QueryPos Error, errormsg: " + string(G_CONNECTION1->GetErrorMsg(ret));
		cout << msg << endl;
		return -1;
	}
	return 0;
}


int __stdcall QueryEntrustByEntrustNo(const char* account, const char* combi_no, int EntrustNo ){
	IBizMessage* lpBizMessage = NewBizMessage();
	lpBizMessage->AddRef();
	//�˻��ʲ���ѯ��32001
	lpBizMessage->SetFunction(32001);
	lpBizMessage->SetPacketType(REQUEST_PACKET);

	IF2Packer* requestPacker = NewPacker(2);
	requestPacker->AddRef();
	requestPacker->BeginPack();
	requestPacker->AddField("user_token",	'S',512,0);
	requestPacker->AddField("account_code", 'S',32,0);
	requestPacker->AddField("combi_no"  ,   'S',8,0); 
	requestPacker->AddField("entrust_no",'I');
	requestPacker->AddStr(gUserToken.c_str());
	requestPacker->AddStr(account);
	requestPacker->AddStr(combi_no);
	requestPacker->AddInt(EntrustNo);
	requestPacker->EndPack();
	IF2Packer* lpPacker = requestPacker;
	lpBizMessage->SetContent(lpPacker->GetPackBuf(),lpPacker->GetPackLen());
	int ret = CallService(G_CONNECTION1, lpBizMessage);
	lpPacker->FreeMem(lpPacker->GetPackBuf());
	lpPacker->Release();
	lpBizMessage->Release();

	if (ret <= 0){
		string msg = "[QueryPos] QueryPos Error, errormsg: " + string(G_CONNECTION1->GetErrorMsg(ret));
		cout << msg << endl;
		return -1;
	}
	return 0;
}


int __stdcall QueryAccount(const char * account, const char * combi_no){
	IBizMessage* lpBizMessage = NewBizMessage();
	lpBizMessage->AddRef();
	//�˻��ʲ���ѯ��34001
	lpBizMessage->SetFunction(34001);
	lpBizMessage->SetPacketType(REQUEST_PACKET);

	IF2Packer* requestPacker = NewPacker(2);
	requestPacker->AddRef();
	requestPacker->BeginPack();
	requestPacker->AddField("user_token",'S',512,0);
	requestPacker->AddField("account_code",      'S',32,0);
	requestPacker->AddField("combi_no"  ,   'S',8,0); 
	requestPacker->AddStr(gUserToken.c_str());
	requestPacker->AddStr(account);
	requestPacker->AddStr(combi_no);

	requestPacker->EndPack();

	IF2Packer* lpPacker = requestPacker;
	lpBizMessage->SetContent(lpPacker->GetPackBuf(),lpPacker->GetPackLen());
	int ret = CallService(G_CONNECTION1, lpBizMessage);
	lpPacker->FreeMem(lpPacker->GetPackBuf());
	lpPacker->Release();
	lpBizMessage->Release();

	if (ret <= 0){
		string msg = "[QueryPos] QueryPos Error, errormsg: " + string(G_CONNECTION1->GetErrorMsg(ret));
		cout << msg << endl;
		return -1;
	}
	return 0;
}



//
//int main(int argc, char** argv){
//
//	Connect("10.2.130.189:18801");
//	SubConnect();
//
//	int iOrderID = 0;
//	while (1){
//		cout << endl;
//		cout << "1����¼ 2��֤ȯ����ί�� 3������ί�� 4������ 0���˳�" << endl;
//		cout << "������ָ��ţ�";
//		scanf("%d",&iOrderID);
//		switch(iOrderID){
//			case 0:
//				{
//					G_SUBSCRIBE->Release();
//					G_CONNECTION2 ->Release();
//					G_CONNECTION1->Release();
//					return 0;
//				}
//			case 1:
//				{
//					//��¼
//					cout << Login() << endl;
//					break;
//				}
//			case 2:
//				{
//					//֤ȯ����ί��
//					cout << LimitEntrust("1007", "1", "600000", "10072", "2", 12.880, 200) << endl;
//					break;
//				}
//			case 4:
//				{
//					//����
//					cout << Withdraw(111) << endl;
//					break;
//				}
//			default:
//				{
//					cout << "�����ָ��Ų���ȷ��" << endl;
//					//continue;
//				}
//		}
//	}
//
//	map<int,CSubscribeParamInterface*>::iterator it = gAllSubscribeParam.begin();
//	for (; it != gAllSubscribeParam.end(); it++)
//	{
//		if ((*it).second != NULL)
//		{
//			(*it).second->Release();
//			(*it).second = NULL;
//		}
//	}
//	gAllSubscribeParam.clear();
//	G_SUBSCRIBE -> Release();
//	G_CONNECTION2 -> Release();
//	G_CONNECTION1 -> Release();
//    getchar();    
//    return 0;
//}

