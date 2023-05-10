package logic

import (
	demoMysqlDao "qsgo-web-templete/dao/mysql/demo"
	"qsgo-web-templete/define/error"
	"qsgo-web-templete/define/types/reply"
	"qsgo-web-templete/define/types/req"
	"qsgo-web-templete/models"

	replyx "github.com/liuxiaobopro/gobox/reply"
)

type demoLogic struct{}

var Demologic = &demoLogic{}

// IndexGet get请求
func (th *demoLogic) IndexGet(in *req.DemoGetReq) (*reply.DemoGetReply, *replyx.T) {
	//TODO: write your logic here
	out := &reply.DemoGetReply{
		Id:   in.Id,
		Name: "IndexGet",
	}
	return out, nil
}

// IndexPost post请求
func (th *demoLogic) IndexPost(in *req.DemoPostReq) (*reply.DemoPostReply, *replyx.T) {
	//TODO: write your logic here
	out := &reply.DemoPostReply{
		Id:   in.Id,
		Name: "IndexPost",
	}
	return out, nil
}

func (th *demoLogic) Create(in *req.DemoCreateReq) (*reply.DemoCreateReply, *replyx.T) {
	var (
		i = &models.Demo{
			Name: in.Name,
		}
		out = &reply.DemoCreateReply{}
	)
	if has, err := demoMysqlDao.DemoDao.ExistByName(i); err != nil {
		return nil, err
	} else if has {
		return nil, replyx.ExistErrT
	}
	if err := demoMysqlDao.DemoDao.Create(i); err != nil {
		return nil, err
	}
	return out, nil
}

func (th *demoLogic) Delete(in *req.DemoDeleteReq) (*reply.DemoDeleteReply, *replyx.T) {
	var (
		out = &reply.DemoDeleteReply{}
	)
	if has, err := demoMysqlDao.DemoDao.ExistById(&models.Demo{Id: in.Id}); err != nil {
		return nil, err
	} else if !has {
		return nil, replyx.NotFoundErrT
	}

	if err := demoMysqlDao.DemoDao.DeleteById(&models.Demo{Id: in.Id}); err != nil {
		return nil, err
	}

	return out, nil
}

func (th *demoLogic) Update(in *req.DemoUpdateReq) (*reply.DemoUpdateReply, *replyx.T) {
	var (
		out = &reply.DemoUpdateReply{}
	)
	if has, err := demoMysqlDao.DemoDao.ExistById(&models.Demo{Id: in.Id}); err != nil {
		return nil, err
	} else if !has {
		return nil, replyx.NotFoundErrT
	}

	if err := demoMysqlDao.DemoDao.UpdateById(&models.Demo{Id: in.Id, Name: in.Name}); err != nil {
		return nil, err
	}
	return out, nil
}

func (th *demoLogic) Detail(in *req.DemoDetailReq) (*reply.DemoDetailReply, *replyx.T) {
	var (
		d   *models.Demo
		err *replyx.T
	)
	if has, err := demoMysqlDao.DemoDao.ExistById(&models.Demo{Id: in.Id}); err != nil {
		return nil, err
	} else if !has {
		return nil, error.ItemNotExist
		// return nil, replyx.NotFoundErrT
	}
	if d, err = demoMysqlDao.DemoDao.DetailById(&models.Demo{Id: in.Id}); err != nil {
		return nil, err
	}
	return &reply.DemoDetailReply{
		Id:   d.Id,
		Name: d.Name,
	}, nil
}

func (th *demoLogic) List(in *req.DemoListReq) (*reply.DemoListReply, *replyx.T) {
	var (
		list *replyx.List
		err  *replyx.T
	)
	if list, err = demoMysqlDao.DemoDao.List(in.Page, in.Size); err != nil {
		return nil, err
	}
	return &reply.DemoListReply{List: list}, nil
}
