/**
这个文件有系统自动生成,请勿修改
*/
package demo

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

func (th *demoDao) Create(demo *models.Demo) *respx.T {
	if _, err := th.mysql.Insert(demo); err != nil {
		th.log.Errorw("DemoDao.Create", "err", err)
		return respx.InternalErrT
	}
	return nil
}

func (th *demoDao) DeleteById(demo *models.Demo) *respx.T {
	if _, err := th.mysql.Where("id = ?", demo.Id).Delete(&models.Demo{}); err != nil {
		th.log.Errorw("DemoDao.Delete", "err", err)
		return respx.InternalErrT
	}
	return nil
}

func (th *demoDao) UpdateById(demo *models.Demo) *respx.T {
	if _, err := th.mysql.Where("id = ?", demo.Id).Update(demo); err != nil {
		th.log.Errorw("DemoDao.Update", "err", err)
		return respx.InternalErrT
	}
	return nil
}

func (th *demoDao) DetailById(demo *models.Demo) (*models.Demo, *respx.T) {
	if _, err := th.mysql.Where("id = ?", demo.Id).Get(demo); err != nil {
		th.log.Errorw("DemoDao.Detail", "err", err)
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

	if count, err = th.mysql.FindAndCount(&demos); err != nil {
		th.log.Errorw("DemoDao.List", "err", err)
		return nil, respx.InternalErrT
	}
	return &respx.List{Count: count, List: demos}, nil
}

