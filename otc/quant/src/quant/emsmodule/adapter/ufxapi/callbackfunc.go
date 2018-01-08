package ufxapi

//#cgo CFLAGS:  -I D:/workspace/jaguar/otc/quant/ufxapi/ufxapi
//#include "response.h"
import "C"
import (
	log "github.com/thinkboy/log4go"
	"quant/emsmodule/adapter"
	emsbase "quant/emsmodule/base"
	"quant/helper"
	"reflect"
	"unsafe"
)

// GoCallBackFunc will be called when UFX receive push data
//export GoCallBackFunc
func GoCallBackFunc(retptr unsafe.Pointer) {
	ret := (*C.struct_Result)(retptr)
	funcNo := int(ret.FuncNo)
	inputs := make([]reflect.Value, 1)
	inputs[0] = reflect.ValueOf(ret)

	if funcNo != -999999 {
		reflectFunc, ok := handleFuncMap[funcNo]
		if ok {
			reflectFunc.Call(inputs)
		} else {
			log.Info("Please register handle function first.[FuncID: %d]\n", funcNo)
		}
	} else {
		// log.Info("PUSH MESSAGE")
		dat := ret.DataSet
		if C.GoString(ret.MsgType) == "g" {
			pushdata := (*C.struct_DealPushResp)(dat)
			tradetmpdat := formatTradeDat(pushdata)
			// log.Info("%+v", tradetmpdat)
			// log.Info("MsgType: %s, third_reff: %s", C.GoString(ret.MsgType), C.GoString(pushdata.third_reff))
			adapter.PushToOMS(emsbase.PushData{MsgType: helper.TradePushData,
				Trade: tradetmpdat,
				Port:  emsbase.Portfolio{ProductInfo: &emsbase.ProductInfo{}, StrategyInfo: &emsbase.StrategyInfo{}}})
		} else {
			pushdata := (*C.struct_EntrustPushResp)(dat)
			entrusttmpdat := formatEntrustDat(pushdata)
			// log.Info("%+v", entrusttmpdat)
			// log.Info("MsgType: %s, third_reff: %s", C.GoString(ret.MsgType), C.GoString(pushdata.third_reff))
			adapter.PushToOMS(emsbase.PushData{MsgType: helper.EntrustRespPushData,
				Entrust: entrusttmpdat,
				Port:    emsbase.Portfolio{ProductInfo: &emsbase.ProductInfo{}, StrategyInfo: &emsbase.StrategyInfo{}}})
		}
	}
}

func formatTradeDat(pushdata *C.struct_DealPushResp) emsbase.DealPushResp {
	tradetmpdat := emsbase.DealPushResp{}
	tradetmpdat.OperatorNo = C.GoString(pushdata.operator_no)
	tradetmpdat.DealDate = int(pushdata.deal_date)
	tradetmpdat.DealTime = int(pushdata.deal_time)
	tradetmpdat.DealNo = C.GoString(pushdata.deal_no)
	tradetmpdat.BatchNo = int(pushdata.batch_no)
	tradetmpdat.EntrustNo = int(pushdata.entrust_no)
	tradetmpdat.MarketNo = C.GoString(pushdata.market_no)
	tradetmpdat.StockCode = C.GoString(pushdata.stock_code)
	tradetmpdat.AccountCode = C.GoString(pushdata.account_code)
	tradetmpdat.CombiNo = C.GoString(pushdata.combi_no)
	tradetmpdat.StockholderID = C.GoString(pushdata.stockholder_id)
	tradetmpdat.ReportSeat = C.GoString(pushdata.report_seat)
	tradetmpdat.EntrustDirection = C.GoString(pushdata.entrust_direction)
	tradetmpdat.FuturesDirection = C.GoString(pushdata.futures_direction)
	tradetmpdat.EntrustAmount = int(pushdata.entrust_amount)
	tradetmpdat.EntrustStatus = C.GoString(pushdata.entrust_status)
	tradetmpdat.DealAmount = int(pushdata.deal_amount)
	tradetmpdat.DealPrice = float64(pushdata.deal_price)
	tradetmpdat.DealBalance = float64(pushdata.deal_balance)
	tradetmpdat.DealFee = float64(pushdata.deal_fee)
	tradetmpdat.TotalDealAmount = int(pushdata.total_deal_amount)
	tradetmpdat.TotalDealBalance = float64(pushdata.total_deal_balance)
	tradetmpdat.CancelAmount = int(pushdata.cancel_amount)
	tradetmpdat.ReportDirection = C.GoString(pushdata.report_direction)
	tradetmpdat.ExtsystemID = int(pushdata.extsystem_id)
	tradetmpdat.ThirdReff = C.GoString(pushdata.third_reff)
	return tradetmpdat
}

func formatEntrustDat(pushdata *C.struct_EntrustPushResp) emsbase.EntrustPushResp {
	entrusttmpdat := emsbase.EntrustPushResp{}
	entrusttmpdat.OperatorNo = C.GoString(pushdata.operator_no)
	entrusttmpdat.AccountCode = C.GoString(pushdata.account_code)
	entrusttmpdat.BatchNo = int(pushdata.batch_no)
	entrusttmpdat.BusinessDate = int(pushdata.business_date)
	entrusttmpdat.BusinessTime = int(pushdata.business_time)
	entrusttmpdat.CombiNo = C.GoString(pushdata.combi_no)
	entrusttmpdat.ConfirmNo = C.GoString(pushdata.confirm_no)
	entrusttmpdat.EntrustAmount = int(pushdata.entrust_amount)
	entrusttmpdat.CancelAmount = int(pushdata.cancel_amount)
	entrusttmpdat.EntrustDirection = C.GoString(pushdata.entrust_direction)
	entrusttmpdat.EntrustNo = C.GoString(pushdata.entrust_no)
	entrusttmpdat.EntrustPrice = float64(pushdata.entrust_price)
	entrusttmpdat.EntrustStatus = C.GoString(pushdata.entrust_status)
	entrusttmpdat.DealAmount = int(pushdata.deal_amount)
	entrusttmpdat.DealBalance = float64(pushdata.deal_balance)
	entrusttmpdat.DealPrice = float64(pushdata.deal_price)
	entrusttmpdat.FuturesDirection = C.GoString(pushdata.futures_direction)
	entrusttmpdat.InvestType = C.GoString(pushdata.invest_type)
	entrusttmpdat.MarketNo = C.GoString(pushdata.market_no)
	entrusttmpdat.PriceType = C.GoString(pushdata.price_type)
	entrusttmpdat.ReportNo = C.GoString(pushdata.report_no)
	entrusttmpdat.ReportSeat = C.GoString(pushdata.report_seat)
	entrusttmpdat.RevokeCause = C.GoString(pushdata.revoke_cause)
	entrusttmpdat.StockCode = C.GoString(pushdata.stock_code)
	entrusttmpdat.StockholderID = C.GoString(pushdata.stockholder_id)
	entrusttmpdat.ThirdReff = C.GoString(pushdata.third_reff)
	entrusttmpdat.ExtsystemID = int(pushdata.extsystem_id)
	return entrusttmpdat
}
