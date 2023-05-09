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
		th.ReturnErr(c, respx.ParamErrT)
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
		th.ReturnErr(c, respx.ParamErrT)
		return
	}
	data, err := logic.Demologic.IndexPost(&r)
	if err != nil {
		th.ReturnErr(c, err)
		return
	}
	th.RetuenOk(c, data)
}

func (th *demoHandle) Detail(c *gin.Context) {
	var r req.DemoDetailReq
	if err := th.ShouldBindUri(c, &r); err != nil { // get=>ShouldBind post=>ShouldBindJSON
		th.ReturnErr(c, respx.ParamErrT)
		return
	}
	data, err := logic.Demologic.Detail(&r)
	if err != nil {
		th.ReturnErr(c, err)
		return
	}
	th.RetuenOk(c, data)
}

func (th *demoHandle) List(c *gin.Context) {
	var r req.DemoListReq
	if err := th.ShouldBind(c, &r); err != nil { // get=>ShouldBind post=>ShouldBindJSON
		th.ReturnErr(c, respx.ParamErrT)
		return
	}
	data, err := logic.Demologic.List(&r)
	if err != nil {
		th.ReturnErr(c, err)
		return
	}
	th.RetuenOk(c, data)
}

func (th *demoHandle) Update(c *gin.Context) {
	var r req.DemoUpdateReq
	if err := th.ShouldBind(c, &r); err != nil { // get=>ShouldBind post=>ShouldBindJSON
		th.ReturnErr(c, respx.ParamErrT)
		return
	}
	data, err := logic.Demologic.Update(&r)
	if err != nil {
		th.ReturnErr(c, err)
		return
	}
	th.RetuenOk(c, data)
}

func (th *demoHandle) Delete(c *gin.Context) {
	var r req.DemoDeleteReq
	if err := th.ShouldBind(c, &r); err != nil { // get=>ShouldBind post=>ShouldBindJSON
		th.ReturnErr(c, respx.ParamErrT)
		return
	}
	data, err := logic.Demologic.Delete(&r)
	if err != nil {
		th.ReturnErr(c, err)
		return
	}
	th.RetuenOk(c, data)
}

func (th *demoHandle) Create(c *gin.Context) {
	var r req.DemoCreateReq
	if err := th.ShouldBind(c, &r); err != nil { // get=>ShouldBind post=>ShouldBindJSON
		th.ReturnErr(c, respx.ParamErrT)
		return
	}
	data, err := logic.Demologic.Create(&r)
	if err != nil {
		th.ReturnErr(c, err)
		return
	}
	th.RetuenOk(c, data)
}
