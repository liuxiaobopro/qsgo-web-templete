package error

import (
	replyx "github.com/liuxiaobopro/gobox/reply"
)

var (
	//#region demo相关
	ItemNotExist = &replyx.T{Code: 10001, Msg: "条目不存在"}
	// #endregion
)
