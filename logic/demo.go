package logic

import (
	mysqlDao "qsgo-web-templete/dao/mysql"
	"qsgo-web-templete/models"
	"qsgo-web-templete/types/reply"
	"qsgo-web-templete/types/req"

	respx "github.com/liuxiaobopro/gobox/resp"
)

type demoLogic struct{}

var Demologic = &demoLogic{}

// IndexGet get请求
func (th *demoLogic) IndexGet(in *req.DemoGetReq) (*reply.DemoGetReply, *respx.T) {
	//TODO: write your logic here
	out := &reply.DemoGetReply{
		Id:   in.Id,
		Name: "IndexGet",
	}
	return out, nil
}

// IndexPost post请求
func (th *demoLogic) IndexPost(in *req.DemoPostReq) (*reply.DemoPostReply, *respx.T) {
	//TODO: write your logic here
	out := &reply.DemoPostReply{
		Id:   in.Id,
		Name: "IndexPost",
	}
	return out, nil
}

func (th *demoLogic) Detail(in *req.DemoDetailReq) (*reply.DemoDetailReply, *respx.T) {
	var (
		demo = &models.Demo{}
		err  *respx.T
	)
	if demo, err = mysqlDao.DemoDao.Detail(in.Id); err != nil {
		return nil, err
	}
	return &reply.DemoDetailReply{
		Id:   demo.Id,
		Name: demo.Name,
	}, nil
}

func (th *demoLogic) List(in *req.DemoListReq) (*reply.DemoListReply, *respx.T) {
	//TODO: write your logic here
	out := &reply.DemoListReply{}
	return out, nil
}

func (th *demoLogic) Update(in *req.DemoUpdateReq) (*reply.DemoUpdateReply, *respx.T) {
	//TODO: write your logic here
	out := &reply.DemoUpdateReply{}
	return out, nil
}

func (th *demoLogic) Add(in *req.DemoAddReq) (*reply.DemoAddReply, *respx.T) {
	//TODO: write your logic here
	out := &reply.DemoAddReply{}
	return out, nil
}

func (th *demoLogic) Delete(in *req.DemoDeleteReq) (*reply.DemoDeleteReply, *respx.T) {
	//TODO: write your logic here
	out := &reply.DemoDeleteReply{}
	return out, nil
}

func (th *demoLogic) Create(in *req.DemoCreateReq) (*reply.DemoCreateReply, *respx.T) {
	//TODO: write your logic here
	out := &reply.DemoCreateReply{}
	return out, nil
}
