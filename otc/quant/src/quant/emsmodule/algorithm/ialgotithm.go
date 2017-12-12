package algorithm

import (
	emsbase "quant/emsmodule/base"
)

// IAlgotithm is the interface which should be implemented by each algorithm.
// Trade will be called by algorithmadmin.go
type iAlgotithm interface {
	init()                         // 初始化
	trade(emsbase.Portfolio) error // 交易
	checkPortStatus(string)        // 检查组合状态
	append()                       // 补单
}
