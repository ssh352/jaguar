

#ifndef _UFX_UTIL_H
#define _UFX_UTIL_H


#include <map>
#include <vector>
#include "t2sdk_interface.h"
#include "ufxapi.h"

using namespace std;

void UFX_10001_unPacker(Result* ret, void** resps, IF2UnPacker* responseUnPacker);
void UFX_91001_unPacker(Result* ret, void** resps, IF2UnPacker* responseUnPacker);
void UFX_91101_unPacker(Result* ret, void** resps, IF2UnPacker* responseUnPacker);
void UFX_35003_unPacker(Result* ret, void** resps, IF2UnPacker* responseUnPacker);
void UFX_31001_unPacker(Result* ret, void** resps, IF2UnPacker* responseUnPacker);
void UFX_32001_unPacker(Result* ret, void** resps, IF2UnPacker* responseUnPacker);
void UFX_34001_unPacker(Result* ret, void** resps, IF2UnPacker* responseUnPacker);

void UFX_HANDLE_ENTRUST_PUSH(Result* ret, void* resp, IF2UnPacker* lpUnPack);
void UFX_HANDLE_DEAL_PUSH(Result* ret, void* resp, IF2UnPacker* lpUnPack);

void ShowPacket(IF2UnPacker* unPacker);


//void ShowSubscribe(int subIndex, LPSUBSCRIBE_RECVDATA lpRecvData, map<int, CSubscribeParamInterface*> gAllSubscribeParam);


#endif //_UFX_UTIL_H