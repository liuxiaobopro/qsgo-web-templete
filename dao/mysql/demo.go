package dao

import (
	"qsgo-web-templete/global"
	"qsgo-web-templete/models"

	respx "github.com/liuxiaobopro/gobox/resp"
	"go.uber.org/zap"
	"xorm.io/xorm"
)

type demoDao struct {
	mysql *xorm.Engine
	log   *zap.SugaredLogger
}

var DemoDao *demoDao

func init() {
	DemoDao = &demoDao{
		mysql: global.DB,
		log:   global.ZapS,
	}
}

func (th *demoDao) Detail(id int) (*models.Demo, *respx.T) {
	demo := &models.Demo{}
	if _, err := th.mysql.Where("id = ?", id).Get(demo); err != nil {
		th.log.Errorw("DemoDao.Detail", "err", err)
		return nil, respx.InternalErrT
	}
	return demo, nil
}

func (th *demoDao) List() ([]*models.Demo, *respx.T) {
	demos := make([]*models.Demo, 0)
	if err := th.mysql.Find(&demos); err != nil {
		th.log.Errorw("DemoDao.List", "err", err)
		return nil, respx.InternalErrT
	}
	return demos, nil
}

func (th *demoDao) Update(demo *models.Demo) *respx.T {
	if _, err := th.mysql.Where("id = ?", demo.Id).Update(demo); err != nil {
		th.log.Errorw("DemoDao.Update", "err", err)
		return respx.InternalErrT
	}
	return nil
}

func (th *demoDao) Create(demo *models.Demo) *respx.T {
	if _, err := th.mysql.Insert(demo); err != nil {
		th.log.Errorw("DemoDao.Create", "err", err)
		return respx.InternalErrT
	}
	return nil
}

func (th *demoDao) Delete(id int) *respx.T {
	if _, err := th.mysql.Where("id = ?", id).Delete(&models.Demo{}); err != nil {
		th.log.Errorw("DemoDao.Delete", "err", err)
		return respx.InternalErrT
	}
	return nil
}
