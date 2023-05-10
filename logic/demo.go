package logic

import (
	demoMysqlDao "qsgo-web-templete/dao/mysql/demo"
	redisDao "qsgo-web-templete/dao/redis"
	"qsgo-web-templete/define/error"
	"qsgo-web-templete/define/types/reply"
	"qsgo-web-templete/define/types/req"
	"qsgo-web-templete/models"

	replyx "github.com/liuxiaobopro/gobox/reply"
)

type demoLogic struct{}

var Demologic = &demoLogic{}

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

func (th *demoLogic) Redis(in *req.DemoRedisReq) (*reply.DemoRedisReply, *replyx.T) {
	// if err := redisDao.SetString(); err != nil {
	// 	return nil, replyx.InternalErrT
	// }

	// if s, err := redisDao.GetString(); err != nil {
	// 	return nil, replyx.InternalErrT
	// } else {
	// 	return &reply.DemoRedisReply{Value: s}, nil
	// }

	// if s, err := redisDao.GetStringNotFound(); err != nil {
	// 	if err == redis.Nil {
	// 		return &reply.DemoRedisReply{Value: ""}, nil
	// 	} else {
	// 		return nil, replyx.InternalErrT
	// 	}
	// } else {
	// 	return &reply.DemoRedisReply{Value: s}, nil
	// }

	if err := redisDao.SetHash(); err != nil {
		return nil, replyx.InternalErrT
	}

	return nil, nil
}
