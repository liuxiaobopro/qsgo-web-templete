/**
这个文件有系统自动生成,请勿修改
*/
package demo

import (
	"time"

	"qsgo-web-templete/global"
	"qsgo-web-templete/models"

	replyx "github.com/liuxiaobopro/gobox/reply"
	timex "github.com/liuxiaobopro/gobox/time"
)

type demoDao struct {}

var DemoDao *demoDao

func (th *demoDao) Create(demo *models.Demo) *replyx.T {
	demo.CreateAt = timex.Time(time.Now())
	if _, err := global.DB.Insert(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Create", "err", err)
		return replyx.InternalErrT
	}
	return nil
}

func (th *demoDao) DeleteById(demo *models.Demo) *replyx.T {
	if _, err := global.DB.Where("Id = ?", demo.Id).Delete(&models.Demo{}); err != nil {
		global.ZapS.Errorw("DemoDao.Delete", "err", err)
		return replyx.InternalErrT
	}
	return nil
}

func (th *demoDao) DeleteByName(demo *models.Demo) *replyx.T {
	if _, err := global.DB.Where("Name = ?", demo.Name).Delete(&models.Demo{}); err != nil {
		global.ZapS.Errorw("DemoDao.Delete", "err", err)
		return replyx.InternalErrT
	}
	return nil
}

func (th *demoDao) DeleteByCreateAt(demo *models.Demo) *replyx.T {
	if _, err := global.DB.Where("CreateAt = ?", demo.CreateAt).Delete(&models.Demo{}); err != nil {
		global.ZapS.Errorw("DemoDao.Delete", "err", err)
		return replyx.InternalErrT
	}
	return nil
}

func (th *demoDao) DeleteByUpdateAt(demo *models.Demo) *replyx.T {
	if _, err := global.DB.Where("UpdateAt = ?", demo.UpdateAt).Delete(&models.Demo{}); err != nil {
		global.ZapS.Errorw("DemoDao.Delete", "err", err)
		return replyx.InternalErrT
	}
	return nil
}


func (th *demoDao) UpdateById(demo *models.Demo) *replyx.T {
	demo.UpdateAt = timex.Time(time.Now())
	if _, err := global.DB.Where("Id = ?", demo.Id).Update(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Update", "err", err)
		return replyx.InternalErrT
	}
	return nil
}

func (th *demoDao) UpdateByName(demo *models.Demo) *replyx.T {
	demo.UpdateAt = timex.Time(time.Now())
	if _, err := global.DB.Where("Name = ?", demo.Name).Update(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Update", "err", err)
		return replyx.InternalErrT
	}
	return nil
}

func (th *demoDao) UpdateByCreateAt(demo *models.Demo) *replyx.T {
	demo.UpdateAt = timex.Time(time.Now())
	if _, err := global.DB.Where("CreateAt = ?", demo.CreateAt).Update(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Update", "err", err)
		return replyx.InternalErrT
	}
	return nil
}

func (th *demoDao) UpdateByUpdateAt(demo *models.Demo) *replyx.T {
	demo.UpdateAt = timex.Time(time.Now())
	if _, err := global.DB.Where("UpdateAt = ?", demo.UpdateAt).Update(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Update", "err", err)
		return replyx.InternalErrT
	}
	return nil
}


func (th *demoDao) DetailById(demo *models.Demo) (*models.Demo, *replyx.T) {
	if _, err := global.DB.Where("Id = ?", demo.Id).Get(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Detail", "err", err)
		return nil, replyx.InternalErrT
	}
	return demo, nil
}

func (th *demoDao) DetailByName(demo *models.Demo) (*models.Demo, *replyx.T) {
	if _, err := global.DB.Where("Name = ?", demo.Name).Get(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Detail", "err", err)
		return nil, replyx.InternalErrT
	}
	return demo, nil
}

func (th *demoDao) DetailByCreateAt(demo *models.Demo) (*models.Demo, *replyx.T) {
	if _, err := global.DB.Where("CreateAt = ?", demo.CreateAt).Get(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Detail", "err", err)
		return nil, replyx.InternalErrT
	}
	return demo, nil
}

func (th *demoDao) DetailByUpdateAt(demo *models.Demo) (*models.Demo, *replyx.T) {
	if _, err := global.DB.Where("UpdateAt = ?", demo.UpdateAt).Get(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Detail", "err", err)
		return nil, replyx.InternalErrT
	}
	return demo, nil
}


func (th *demoDao) ExistById(demo *models.Demo) (bool, *replyx.T) {
	var (
		has bool
		err error
	)
	if has, err = global.DB.Where("Id = ?", demo.Id).Exist(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Detail", "err", err)
		return false, replyx.InternalErrT
	}
	return has, nil
}

func (th *demoDao) ExistByName(demo *models.Demo) (bool, *replyx.T) {
	var (
		has bool
		err error
	)
	if has, err = global.DB.Where("Name = ?", demo.Name).Exist(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Detail", "err", err)
		return false, replyx.InternalErrT
	}
	return has, nil
}

func (th *demoDao) ExistByCreateAt(demo *models.Demo) (bool, *replyx.T) {
	var (
		has bool
		err error
	)
	if has, err = global.DB.Where("CreateAt = ?", demo.CreateAt).Exist(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Detail", "err", err)
		return false, replyx.InternalErrT
	}
	return has, nil
}

func (th *demoDao) ExistByUpdateAt(demo *models.Demo) (bool, *replyx.T) {
	var (
		has bool
		err error
	)
	if has, err = global.DB.Where("UpdateAt = ?", demo.UpdateAt).Exist(demo); err != nil {
		global.ZapS.Errorw("DemoDao.Detail", "err", err)
		return false, replyx.InternalErrT
	}
	return has, nil
}

func (th *demoDao) List(page, size int) (*replyx.List, *replyx.T) {
	var (
		demos = make([]*models.Demo, 0)
		count int64
		err   error
	)

	if count, err = global.DB.Limit(size, (page-1)*size).FindAndCount(&demos); err != nil {
		global.ZapS.Errorw("DemoDao.List", "err", err)
		return nil, replyx.InternalErrT
	}
	return &replyx.List{Count: count, List: demos}, nil
}

