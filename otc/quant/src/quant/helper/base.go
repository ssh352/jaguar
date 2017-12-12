package helper

//Request 模块间调用请求
type Request struct {
	From   string   // 调用发起模块
	To     string   // 目标模块
	Cmd    string   // 命令
	Params []string // 参数【从左往右填入Cmd】
}
