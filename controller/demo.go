package controller

import (
	"qsgo-web-templete/define/types/req"
	"qsgo-web-templete/logic"

	"github.com/gin-gonic/gin"
	httpx "github.com/liuxiaobopro/gobox/http"
	replyx "github.com/liuxiaobopro/gobox/reply"
)

type demoHandle struct {
	httpx.GinHandle
}

var DemoController = &demoHandle{}

func (th *demoHandle) Create(c *gin.Context) {
	var r req.DemoCreateReq
	if err := th.ShouldBind(c, &r); err != nil {
		th.ReturnErr(c, replyx.ParamErrT)
		return
	}
	data, err := logic.Demologic.Create(&r)
	if err != nil {
		th.ReturnErr(c, err)
		return
	}
	th.RetuenOk(c, data)
}

func (th *demoHandle) Delete(c *gin.Context) {
	var r req.DemoDeleteReq
	if err := th.ShouldBindUri(c, &r); err != nil {
		th.ReturnErr(c, replyx.ParamErrT)
		return
	}
	data, err := logic.Demologic.Delete(&r)
	if err != nil {
		th.ReturnErr(c, err)
		return
	}
	th.RetuenOk(c, data)
}

func (th *demoHandle) Update(c *gin.Context) {
	var r req.DemoUpdateReq
	if err := th.ShouldBindUri(c, &r); err != nil {
		th.ReturnErr(c, replyx.ParamErrT)
		return
	}
	if err := th.ShouldBindJSON(c, &r); err != nil {
		th.ReturnErr(c, replyx.ParamErrT)
		return
	}
	data, err := logic.Demologic.Update(&r)
	if err != nil {
		th.ReturnErr(c, err)
		return
	}
	th.RetuenOk(c, data)
}

func (th *demoHandle) Detail(c *gin.Context) {
	var r req.DemoDetailReq
	if err := th.ShouldBindUri(c, &r); err != nil {
		th.ReturnErr(c, replyx.ParamErrT)
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
	if err := th.ShouldBind(c, &r); err != nil {
		th.ReturnErr(c, replyx.ParamErrT)
		return
	}
	data, err := logic.Demologic.List(&r)
	if err != nil {
		th.ReturnErr(c, err)
		return
	}
	th.RetuenOk(c, data)
}

func (th *demoHandle) Redis(c *gin.Context) {
	var r req.DemoRedisReq
	if err := th.ShouldBind(c, &r); err != nil {
		th.ReturnErr(c, replyx.ParamErrT)
		return
	}
	data, err := logic.Demologic.Redis(&r)
	if err != nil {
		th.ReturnErr(c, err)
		return
	}
	th.RetuenOk(c, data)
}
