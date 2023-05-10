package req

import (
	reqx "github.com/liuxiaobopro/gobox/req"
)

type DemoDetailReq struct {
	Id int `form:"id" json:"id" uri:"id" binding:"required"`
}

type DemoListReq struct{
	reqx.List
}

type DemoUpdateReq struct {
	Id   int    `form:"id" json:"id" uri:"id"`
	Name string `form:"name" json:"name" uri:"name"`
}

type DemoAddReq struct{}

type DemoDeleteReq struct {
	Id int `form:"id" json:"id" uri:"id" binding:"required"`
}

type DemoCreateReq struct {
	Name string `form:"name" json:"name" uri:"name" binding:"required"`
}
