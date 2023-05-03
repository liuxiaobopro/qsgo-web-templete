package controller

import (
	"qsgo-web-templete/logic"
	"qsgo-web-templete/types/req"

	"github.com/gin-gonic/gin"
	httpx "github.com/liuxiaobopro/gobox/http"
	respx "github.com/liuxiaobopro/gobox/resp"
)

type demoHandle struct {
	httpx.GinHandle
}

var DemoController = &demoHandle{}

// IndexGet get请求
func (th *demoHandle) IndexGet(c *gin.Context) {
	var r req.DemoGetReq
	if err := th.ShouldBind(c, &r); err != nil {
		th.ReturnErr(c, respx.ParamErrT.ToPt())
		return
	}
	data, err := logic.Demologic.IndexGet(&r)
	if err != nil {
		th.ReturnErr(c, err)
		return
	}
	th.RetuenOk(c, data)
}

// IndexPost post请求
func (th *demoHandle) IndexPost(c *gin.Context) {
	var r req.DemoPostReq
	if err := th.ShouldBindJSON(c, &r); err != nil {
		th.ReturnErr(c, respx.ParamErrT.ToPt())
		return
	}
	data, err := logic.Demologic.IndexPost(&r)
	if err != nil {
		th.ReturnErr(c, err)
		return
	}
	th.RetuenOk(c, data)
}
