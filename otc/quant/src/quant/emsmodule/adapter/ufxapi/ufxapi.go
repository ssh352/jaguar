package ufxapi

//#cgo CFLAGS:  -I E:/jaguar/sandbox/C++/ufxapi/ufxapi
//#include <stdio.h>
//#include <stdlib.h>
//#include <windows.h>
//#include "ufxapi.h"
//#include "response.h"
/*

typedef int (* FuncRegisterCallBackPtr)(ufxcallback);
typedef int (* FuncConnectPtr)(const char *, int);
typedef int (* FuncSubConnectPtr)();
typedef int (* FuncLoginPtr)();
typedef int (* FuncLimitEntrustPtr)(const char *, const char *, const char *, const char *, const char *, double, int, const char *);
typedef int (* FuncWithdrawPtr)(int);
typedef int (* FuncQueryPosPtr)(const char*, const char*);
typedef int (* FuncQueryEntrustByAccPtr)(const char* account, const char* combi_bo);
typedef int (* FuncQueryEntrustByEntrustNoPtr)(const char*, const char*, int);
typedef int (* FuncQueryAccountPtr)(const char*, const char*);

HINSTANCE 						G_HANDLE 						= NULL;
FuncConnectPtr 					G_ConnectPtr 					= NULL;
FuncRegisterCallBackPtr 		G_RegisterCallBackPtr 			= NULL;
FuncLoginPtr 					G_LoginPtr 						= NULL;
FuncSubConnectPtr 				G_SubConnectPtr 				= NULL;
FuncLimitEntrustPtr 			G_LimitEntrustPtr				= NULL;
FuncWithdrawPtr					G_WithdrawPtr					= NULL;
FuncQueryPosPtr					G_QueryPosPtr					= NULL;
FuncQueryEntrustByAccPtr 		G_QueryEntrustByAccPtr 			= NULL;
FuncQueryEntrustByEntrustNoPtr 	G_QueryEntrustByEntrustNoPtr 	= NULL;
FuncQueryAccountPtr 			G_QueryAccountPtr 				= NULL;

extern void GoCallBackFunc(void*);

int RegisterUFXCallBack();
int UFXCallBack(Result* ret);


int UFXInit(){
	G_HANDLE 				= LoadLibrary("E:\\otc\\quant\\ufxapi.dll");
	G_ConnectPtr 			= (FuncConnectPtr)GetProcAddress(G_HANDLE, "Connect");
	G_RegisterCallBackPtr 	= (FuncRegisterCallBackPtr)GetProcAddress(G_HANDLE, "RegisterCallBack");
	G_LoginPtr 				= (FuncLoginPtr)GetProcAddress(G_HANDLE, "Login");
	G_SubConnectPtr 		= (FuncSubConnectPtr)GetProcAddress(G_HANDLE, "SubConnect");
	G_LimitEntrustPtr		= (FuncLimitEntrustPtr)GetProcAddress(G_HANDLE, "LimitEntrust");
	G_WithdrawPtr			= (FuncWithdrawPtr)GetProcAddress(G_HANDLE, "Withdraw");
	G_QueryPosPtr			= (FuncQueryPosPtr)GetProcAddress(G_HANDLE, "QueryPos");
	G_QueryEntrustByAccPtr 		 = (FuncQueryEntrustByAccPtr)GetProcAddress(G_HANDLE, "QueryEntrustByAcc");
	G_QueryEntrustByEntrustNoPtr = (FuncQueryEntrustByEntrustNoPtr)GetProcAddress(G_HANDLE, "QueryEntrustByEntrustNo");
	G_QueryAccountPtr 			 = (FuncQueryAccountPtr)GetProcAddress(G_HANDLE, "QueryAccount");
	RegisterUFXCallBack();
	return 0;
}

int UFXConnect(const char* serverIp, int sync){
	return G_ConnectPtr(serverIp, sync);
}

int RegisterUFXCallBack(){
	ufxcallback func = &UFXCallBack;
	if(G_RegisterCallBackPtr != NULL){
		return G_RegisterCallBackPtr(func);
	}else{
		return -2;
	}
}

int UFXLogin(){
	if(G_LoginPtr != NULL){
		return G_LoginPtr();
	}else{
		return -2;
	}
}

int UFXSubConnect(){
	if(G_SubConnectPtr != NULL){
		return G_SubConnectPtr();
	}else{
		return -2;
	}
}

int UFXLimitEntrust(const char* account_code, const char* market_no, const char* stock_code,
								   const char* combi_no, const char* BS, double price, int vol, const char* third_reff){
	if(G_LimitEntrustPtr != NULL){
		return G_LimitEntrustPtr(account_code, market_no, stock_code, combi_no, BS, price, vol, third_reff);
	}else{
		return -2;
	}
}

int UFXWithdraw(int entrustno){
	if(G_WithdrawPtr != NULL){
		return G_WithdrawPtr(entrustno);
	}else{
		return -2;
	}
}

int UFXQueryPos(const char* account, const char* combi_bo){
	if(G_QueryPosPtr != NULL){
		return G_QueryPosPtr(account, combi_bo);
	}else{
		return -2;
	}
}


int UFXQueryEntrustByAcc(const char* account, const char* combi_no){
	if(G_QueryEntrustByAccPtr != NULL){
		return G_QueryEntrustByAccPtr(account, combi_no);
	}else{
		return -2;
	}
}

int UFXQueryEntrustByEntrustNo(const char* account, const char* combi_no, int EntrustNo){
	if(G_QueryEntrustByEntrustNoPtr != NULL){
		return G_QueryEntrustByEntrustNoPtr(account, combi_no, EntrustNo);
	}else{
		return -2;
	}
}

int UFXQueryAccount(const char* account, const char* combi_no){
	if(G_QueryAccountPtr != NULL){
		return G_QueryAccountPtr(account, combi_no);
	}else{
		return -2;
	}
}

int UFXClose(){
	if(G_HANDLE != NULL){
		FreeLibrary(G_HANDLE);
	}
	return 0;
}

int UFXCallBack(Result* ret){
	GoCallBackFunc(ret);
	return 0;
}

char* xmalloc2(int len){
    static const char* s = "0123456789";
    char* p = malloc(len);
	memcpy(p, s, len);
    return p;
}

void freemem(void* p){
	free(p);
}
*/
import "C"
import (
	"fmt"
	log "github.com/thinkboy/log4go"
	"github.com/widuu/goini"
	emsbase "quant/emsmodule/base"
	"quant/helper"
	"strconv"
	"sync"
	"unsafe"
)

// ufxapi is uesed for communication with O32
type ufxapi struct {
	conf *goini.Config
	addr string
	mux  sync.Mutex
}

// ufx91001RespHandle handle limitEntrust response
func (ua *ufxapi) ufx91001RespHandle(ret *C.struct_Result) string {
	// resp := (*C.struct_EntrustResp)(ret.DataSet)
	return "0"
}

// ufx10001RespHandle handle Login response
func (ua *ufxapi) ufx10001RespHandle(ret *C.struct_Result) {
	if ret.ErrorCode < 0 {
		r := emsbase.LoginResp{RespError: &emsbase.RespError{ErrorCode: int(ret.ErrorCode), ErrorMsg: C.GoString(ret.ErrorMsg)}}
		loginRespWc <- r
	} else {
		resp := (*C.struct_LoginResp)(ret.DataSet)
		loginRespWc <- emsbase.LoginResp{UserToken: C.GoString(resp.user_token),
			VersionNo: C.GoString(resp.version_no),
			RespError: &emsbase.RespError{}}
	}
}

// ufx91101RespHandle handle Withdraw response
func (ua *ufxapi) ufx91101RespHandle(ret *C.struct_Result) string {
	resp := (*C.struct_EntrustResp)(ret.DataSet)
	fmt.Printf("%+v\n", *resp)
	return "0"
}

// ufx31001RespHandle handle QueryPos response
func (ua *ufxapi) ufx31001RespHandle(ret *C.struct_Result) string {
	resp := (*C.struct_QueryPosResp)(ret.DataSet)
	if ret.ErrorCode >= 0 {
		ndata := int(ret.DataCount)
		for i := 0; i < ndata; i++ {
			fmt.Printf("---%d---  stock_code: %s, enable_amount: %d, last_price: %f, cost_price: %f, total_profit: %f, floating_profit: %f, accumulate_profit: %f\n",
				i, C.GoString(resp.stock_code), resp.enable_amount, resp.last_price, resp.cost_price, resp.total_profit, resp.floating_profit, resp.accumulate_profit)
			resp = (*C.struct_QueryPosResp)(resp.nextdataptr)
		}
	} else {
		return C.GoString(ret.ErrorMsg)
	}
	return "0"
}

// ufx32001RespHandle handle QueryEntrustByAcc response
func (ua *ufxapi) ufx32001RespHandle(ret *C.struct_Result) string {
	resp := (*C.struct_QueryEntrustResp)(ret.DataSet)
	if ret.ErrorCode >= 0 {
		ndata := int(ret.DataCount)
		for i := 0; i < ndata; i++ {
			fmt.Printf("stock_code: %s, entrust_date: %d, entrust_time: %d, entrust_no: %d, entrust_price: %f, entrust_amount: %d\n",
				C.GoString(resp.stock_code), resp.entrust_date, resp.entrust_time, resp.entrust_no, resp.entrust_price, resp.entrust_amount)
			resp = (*C.struct_QueryEntrustResp)(resp.nextentrustptr)
		}
	} else {
		return C.GoString(ret.ErrorMsg)
	}
	return "0"
}

// ufx34001RespHandle handle QueryAccount response
func (ua *ufxapi) ufx34001RespHandle(ret *C.struct_Result) string {
	resp := (*C.struct_QueryAccountResp)(ret.DataSet)
	fmt.Printf("account: %s enable_balance_t0: %f enable_balance_t1: %f current_balance: %f\n",
		C.GoString(resp.account_code), resp.enable_balance_t0, resp.enable_balance_t1, resp.current_balance)
	return "0"
}

// Init call UFXInit\UFXConnect\UFXSubConnect\UFXLogin function
// UFXInit init function pointer with ufxapi.dll
// UFXConnect connenct to O32
// UFXSubConnect register callback function(when order is filled or part filled or order status change)
// UFXLogin the operator login, not the user.
func (ua *ufxapi) init() int {
	log.Info("EMS ufxapi init")
	ua.conf = goini.SetConfig(helper.QuantConfigFile)
	C.UFXInit()
	ua.addr = ua.conf.GetStr("emsmodule", "ufx")
	ret, _ := ua.connect(ua.addr, 1)
	if ret == 0 {
		log.Info("EMS ufxapi connect %s success.", ua.addr)
	} else {
		log.Error("EMS ufxapi connect %s fail.", ua.addr)
		return ret
	}

	ret = int(C.UFXSubConnect())
	if ret == 0 {
		log.Info("EMS ufxapi subconnect %s success.", ua.addr)
	} else {
		log.Error("EMS ufxapi subconnect %s fail.", ua.addr)
		return ret
	}

	return 0
}

func (ua *ufxapi) login() {
	C.UFXLogin()
}

func (ua *ufxapi) connect(addr string, sync int) (int, error) {
	serverIP := C.CString(addr)
	defer C.free(unsafe.Pointer(serverIP))
	return int(C.UFXConnect(serverIP, C.int(sync))), nil
}

// limitEntrust place limit price order.
// ufx91001RespHandle function will be called if success
func (ua *ufxapi) limitEntrust(e emsbase.Entrust, AccountCode, ComiNo string) {
	// log.Info("ufxapi limitEntrust")
	cAccountCode := C.CString(AccountCode)
	defer C.free(unsafe.Pointer(cAccountCode))

	cMarkerNo := C.CString(e.MarkerNo)
	defer C.free(unsafe.Pointer(cMarkerNo))

	cStockCode := C.CString(e.TradeCode)
	defer C.free(unsafe.Pointer(cStockCode))

	cComiNo := C.CString(ComiNo)
	defer C.free(unsafe.Pointer(cComiNo))

	cBS := C.CString(strconv.Itoa(e.BS))
	defer C.free(unsafe.Pointer(cBS))

	cThirdReff := C.CString(strconv.FormatInt(e.ID, 10))
	defer C.free(unsafe.Pointer(cThirdReff))

	C.UFXLimitEntrust(cAccountCode, cMarkerNo, cStockCode, cComiNo, cBS, C.double(e.Price), C.int(e.Vol), cThirdReff)
}

// Withdraw order by EntrustNo
// UFX_91101_RESP_HANDLE function will be called if success
func (ua *ufxapi) withdraw(EntrustNo int) {
	C.UFXWithdraw(C.int(EntrustNo))
}

// QueryPos query position by AccountCode, ComiNo
// UFX_31001_RESP_HANDLE function will be called if success
func (ua *ufxapi) queryPos(AccountCode string, ComiNo string) {
	cAccountCode := C.CString(AccountCode)
	defer C.free(unsafe.Pointer(cAccountCode))
	cComiNo := C.CString(ComiNo)
	defer C.free(unsafe.Pointer(cComiNo))
	C.UFXQueryPos(cAccountCode, cComiNo)
}

// QueryAccount query acount information by AccountCode, ComiNo.
// UFX_34001_RESP_HANDLE function will be called if success
func (ua *ufxapi) queryAccount(AccountCode string, ComiNo string) {
	cAccountCode := C.CString(AccountCode)
	defer C.free(unsafe.Pointer(cAccountCode))
	cComiNo := C.CString(ComiNo)
	defer C.free(unsafe.Pointer(cComiNo))
	C.UFXQueryAccount(cAccountCode, cComiNo)
}

// QueryEntrustByAcc query entrust by AccountCode, ComiNo.
// UFX_32001_RESP_HANDLE function will be called if success
func (ua *ufxapi) queryEntrustByAcc(AccountCode, ComiNo string) {
	cAccountCode := C.CString(AccountCode)
	defer C.free(unsafe.Pointer(cAccountCode))
	cComiNo := C.CString(ComiNo)
	defer C.free(unsafe.Pointer(cComiNo))
	C.UFXQueryEntrustByAcc(cAccountCode, cComiNo)
}

// QueryEntrustByEntrustNo query entrust by AccountCode, ComiNo, EntrustNo.
// UFX_32001_RESP_HANDLE function will be called if success
func (ua *ufxapi) queryEntrustByEntrustNo(AccountCode string, ComiNo string, EntrustNo int) {
	cAccountCode := C.CString(AccountCode)
	defer C.free(unsafe.Pointer(cAccountCode))
	cComiNo := C.CString(ComiNo)
	defer C.free(unsafe.Pointer(cComiNo))
	C.UFXQueryEntrustByEntrustNo(cAccountCode, cComiNo, C.int(EntrustNo))
}
