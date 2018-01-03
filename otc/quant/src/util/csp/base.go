package csp

//Request 模块间调用请求
type Request struct {
	TO     string   // 目标模块
	FROM   string   // 调用发起模块
	CMD    string   // 调用请求
	PARAMS []string // 参数【从左往右填入CMD】
}

//Response 模块间调用返回应答
type Response struct {
	TO   string // 应答路由目标
	FROM string // 应答发送模块
	CMD  string // 调用请求
	RET  int    // 错误标识
	MSG  string // 应答信息
	DAT  []byte // 返回数据
}

// IHandleMsg the interface used to deal with request, service should implement this interface
type IHandleMsg interface {
	HandleReq(Request) Response // deal with request
	HandleBReq([]byte) []byte
}

// IRequestMsg the interface used to send request to service, client should implement this interface
type IRequestMsg interface {
	Request(Request) Response // send request to service
	RequestB([]byte) []byte
}

// SetRepV set response field value by req
func SetRepV(req *Request, rep *Response) {
	rep.FROM = req.TO
	rep.TO = req.FROM
	rep.CMD = req.CMD
}
