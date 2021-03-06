package ufxapi

import (
	// log "github.com/thinkboy/log4go"
	"github.com/widuu/goini"
	emsbase "quant/emsmodule/base"
	"quant/helper"
	"reflect"
	"time"
)

// UFX wrap ufxapi. It is used by emsmodule
type UFX struct {
	api     ufxapi
	timeout int
}

var (
	// HandleFuncMap is used by callbackfunc.go when GoCallBackFunc callde by ufx
	handleFuncMap map[int]reflect.Value
	loginRespWc   = make(chan emsbase.LoginResp, 1)
)

// Init filled handleFuncMap with handle function. handleFuncMap will be used when ufx callback in callbackfunc.go file
func (u *UFX) Init() int {
	u.api = ufxapi{}
	u.api.init()
	handleFuncMap = make(map[int]reflect.Value)
	handleFuncMap[91001] = reflect.ValueOf(u.api.ufx91001RespHandle)
	handleFuncMap[10001] = reflect.ValueOf(u.api.ufx10001RespHandle)
	handleFuncMap[91101] = reflect.ValueOf(u.api.ufx91101RespHandle)
	handleFuncMap[31001] = reflect.ValueOf(u.api.ufx31001RespHandle)
	handleFuncMap[32001] = reflect.ValueOf(u.api.ufx32001RespHandle)
	handleFuncMap[34001] = reflect.ValueOf(u.api.ufx34001RespHandle)

	conf := goini.SetConfig(helper.QuantConfigFile)
	u.timeout = conf.GetInt(helper.ConfigEMSSessionName, helper.ConfigEMSTimeout)
	u.login()
	return 0
}

func (u *UFX) login() (emsbase.LoginResp, error) {
	u.api.login()
	to := time.NewTimer(time.Millisecond * time.Duration(u.timeout))
	for {
		to.Reset(time.Millisecond * time.Duration(u.timeout))
		select {
		case resp := <-loginRespWc:
			// log.Info("token:%s, version:%s", resp.UserToken, resp.VersionNo)
			return resp, nil
		case <-to.C:
			return emsbase.LoginResp{}, &emsbase.RespError{-1, "UFX login timeout."}
		}
	}
}

func (u *UFX) LimitEntrust(e emsbase.Entrust, AccountCode, ComiNo string) {
	u.api.limitEntrust(e, AccountCode, ComiNo)
}
func (u *UFX) QueryPos(AccountCode string, ComiNo string)                               {}
func (u *UFX) QueryAccount(AccountCode string, ComiNo string)                           {}
func (u *UFX) QueryEntrustByAcc(AccountCode string, ComiNo string)                      {}
func (u *UFX) QueryEntrustByEntrustNo(AccountCode string, ComiNo string, EntrustNo int) {}
