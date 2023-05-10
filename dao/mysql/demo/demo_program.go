/**
这个文件有系统自动生成,请勿修改
*/
package demo

import (
	"qsgo-web-templete/global"
	"qsgo-web-templete/models"

	respx "github.com/liuxiaobopro/gobox/resp"
)

type demoDao struct {}

var DemoDao *demoDao

func (th *demoDao) Create(demo *models.Demo) *respx.T {
	if _, err := global.DB.Insert(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Create", "err", err)
		return respx.InternalErrT
	}
	return nil
}

func (th *demoDao) DeleteById(demo *models.Demo) *respx.T {
	if _, err := global.DB.Where("Id = ?", demo.Id).Delete(&models.Demo{}); err != nil {
		global.ZapS.Errorw("DemoDao.Delete", "err", err)
		return respx.InternalErrT
	}
	return nil
}

func (th *demoDao) DeleteByName(demo *models.Demo) *respx.T {
	if _, err := global.DB.Where("Name = ?", demo.Name).Delete(&models.Demo{}); err != nil {
		global.ZapS.Errorw("DemoDao.Delete", "err", err)
		return respx.InternalErrT
	}
	return nil
}

func (th *demoDao) DeleteByCreateAt(demo *models.Demo) *respx.T {
	if _, err := global.DB.Where("CreateAt = ?", demo.CreateAt).Delete(&models.Demo{}); err != nil {
		global.ZapS.Errorw("DemoDao.Delete", "err", err)
		return respx.InternalErrT
	}
	return nil
}

func (th *demoDao) DeleteByUpdateAt(demo *models.Demo) *respx.T {
	if _, err := global.DB.Where("UpdateAt = ?", demo.UpdateAt).Delete(&models.Demo{}); err != nil {
		global.ZapS.Errorw("DemoDao.Delete", "err", err)
		return respx.InternalErrT
	}
	return nil
}


func (th *demoDao) UpdateById(demo *models.Demo) *respx.T {
	if _, err := global.DB.Where("Id = ?", demo.Id).Update(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Update", "err", err)
		return respx.InternalErrT
	}
	return nil
}

func (th *demoDao) UpdateByName(demo *models.Demo) *respx.T {
	if _, err := global.DB.Where("Name = ?", demo.Name).Update(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Update", "err", err)
		return respx.InternalErrT
	}
	return nil
}

func (th *demoDao) UpdateByCreateAt(demo *models.Demo) *respx.T {
	if _, err := global.DB.Where("CreateAt = ?", demo.CreateAt).Update(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Update", "err", err)
		return respx.InternalErrT
	}
	return nil
}

func (th *demoDao) UpdateByUpdateAt(demo *models.Demo) *respx.T {
	if _, err := global.DB.Where("UpdateAt = ?", demo.UpdateAt).Update(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Update", "err", err)
		return respx.InternalErrT
	}
	return nil
}


func (th *demoDao) DetailById(demo *models.Demo) (*models.Demo, *respx.T) {
	if _, err := global.DB.Where("Id = ?", demo.Id).Get(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Detail", "err", err)
		return nil, respx.InternalErrT
	}
	return demo, nil
}

func (th *demoDao) DetailByName(demo *models.Demo) (*models.Demo, *respx.T) {
	if _, err := global.DB.Where("Name = ?", demo.Name).Get(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Detail", "err", err)
		return nil, respx.InternalErrT
	}
	return demo, nil
}

func (th *demoDao) DetailByCreateAt(demo *models.Demo) (*models.Demo, *respx.T) {
	if _, err := global.DB.Where("CreateAt = ?", demo.CreateAt).Get(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Detail", "err", err)
		return nil, respx.InternalErrT
	}
	return demo, nil
}

func (th *demoDao) DetailByUpdateAt(demo *models.Demo) (*models.Demo, *respx.T) {
	if _, err := global.DB.Where("UpdateAt = ?", demo.UpdateAt).Get(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Detail", "err", err)
		return nil, respx.InternalErrT
	}
	return demo, nil
}

func (th *demoDao) List() (*respx.List, *respx.T) {
	var (
		demos = make([]*models.Demo, 0)
		count int64
		err   error
	)

	if count, err = global.DB.FindAndCount(&demos); err != nil {
		global.ZapS.Errorw("DemoDao.List", "err", err)
		return nil, respx.InternalErrT
	}
	return &respx.List{Count: count, List: demos}, nil
}

