


// libufxapi.h
#ifndef _LIB_UFXAPI_H
#define _LIB_UFXAPI_H


#include "response.h"
typedef int (*ufxcallback)(Result*);

#ifdef __cplusplus
extern "C" __declspec(dllexport) int __stdcall Connect(const char *, int);
extern "C" __declspec(dllexport) int __stdcall SubConnect();
extern "C" __declspec(dllexport) Result* __stdcall LoginSync();
extern "C" __declspec(dllexport) int __stdcall Login();
extern "C" __declspec(dllexport) int __stdcall LimitEntrust(const char *, const char *, const char *, 
																		const char *, const char *, double, int, const char*);
extern "C" __declspec(dllexport) int __stdcall Withdraw(int);
extern "C" __declspec(dllexport) int __stdcall RegisterCallBack(ufxcallback);
extern "C" __declspec(dllexport) int __stdcall HeartBeat();
extern "C" __declspec(dllexport) int __stdcall ExitUFX();
extern "C" __declspec(dllexport) int __stdcall QueryFundAsset(const char* account);
extern "C" __declspec(dllexport) int __stdcall QueryPos(const char* account, const char* combi_bo);
extern "C" __declspec(dllexport) int __stdcall QueryEntrustByAcc(const char* account, const char* combi_bo);
extern "C" __declspec(dllexport) int __stdcall QueryEntrustByEntrustNo(const char* account, const char* combi_no, int EntrustNo);
extern "C" __declspec(dllexport) int __stdcall QueryAccount(const char * account, const char * combi_no);


#else
int __stdcall	Connect(const char *, int);
int __stdcall	SubConnect();
int __stdcall	Login();
int __stdcall	LimitEntrust(const char *, const char *, const char *, 
										  const char *, const char *, double, int, const char*);
int __stdcall	Withdraw(int);
int __stdcall	RegisterCallBack(ufxcallback);
int __stdcall	HeartBeat();
int __stdcall	ExitUFX();
int __stdcall	QueryFundAsset(const char* account);
int __stdcall	QueryPos(const char* account, const char* combi_bo);
int __stdcall	QueryEntrustByAcc(const char* account, const char* combi_bo);
int __stdcall	QueryEntrustByEntrustNo(const char* account, const char* combi_no, int EntrustNo);
int __stdcall	QueryAccount(const char * account, const char * combi_no);

msginfo __stdcall ReturnStruct2();

#endif //__cplusplus

#endif // _LIB_UFXAPI_H