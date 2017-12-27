package helper

//Request 模块间调用请求
type Request struct {
	From   string   // 调用发起模块
	To     string   // 目标模块
	Cmd    string   // 调用请求
	Params []string // 参数【从左往右填入Cmd】
}

//Response 模块间调用返回应答
type Response struct {
	From string   // 应答发送模块
	To   string   // 应答路由目标
	Cmd  string   // 调用请求
	Ret  int      // 错误标识
	Msg  string   // 应答信息
	Dat  []string // 返回数据
}
