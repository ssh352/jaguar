#pragma once

#include ".\ThostTraderApi\ThostFtdcMdApi.h"

typedef int(*mdcallback)(CThostFtdcDepthMarketDataField*);

// subscribe market data
extern "C" __declspec(dllexport) int __stdcall SubMd(char** , int);
extern "C" __declspec(dllexport) int __stdcall RegisterMdCallBack(mdcallback);

