package logic

import (
	"qsgo-web-templete/types/reply"
	"qsgo-web-templete/types/req"

	respx "github.com/liuxiaobopro/gobox/resp"
)

type demoLogic struct{}

var Demologic = &demoLogic{}

// IndexGet get请求
func (th *demoLogic) IndexGet(in *req.DemoGetReq) (out *reply.DemoGetReply, err *respx.Pt) {
	//TODO: write your logic here
	out = &reply.DemoGetReply{
		Id:   in.Id,
		Name: "IndexGet",
	}
	return
}

// IndexPost post请求
func (th *demoLogic) IndexPost(in *req.DemoPostReq) (out *reply.DemoPostReply, err *respx.Pt) {
	//TODO: write your logic here
	out = &reply.DemoPostReply{
		Id:   in.Id,
		Name: "IndexPost",
	}
	return
}
