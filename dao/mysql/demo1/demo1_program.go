/**
这个文件有系统自动生成,请勿修改
*/
package demo1

import (
	"qsgo-web-templete/global"
	"qsgo-web-templete/models"

	respx "github.com/liuxiaobopro/gobox/resp"
	"go.uber.org/zap"
	"xorm.io/xorm"
)

type demo1Dao struct {
	mysql *xorm.Engine
	log   *zap.SugaredLogger
}

var Demo1Dao *demo1Dao

func init() {
	Demo1Dao = &demo1Dao{
		mysql: global.DB,
		log:   global.ZapS,
	}
}

func (th *demo1Dao) Create(demo1 *models.Demo1) *respx.T {
	if _, err := th.mysql.Insert(demo1); err != nil {
		th.log.Errorw("Demo1Dao.Create", "err", err)
		return respx.InternalErrT
	}
	return nil
}

func (th *demo1Dao) DeleteById(demo1 *models.Demo1) *respx.T {
	if _, err := th.mysql.Where("id = ?", demo1.Id).Delete(&models.Demo1{}); err != nil {
		th.log.Errorw("Demo1Dao.Delete", "err", err)
		return respx.InternalErrT
	}
	return nil
}

func (th *demo1Dao) UpdateById(demo1 *models.Demo1) *respx.T {
	if _, err := th.mysql.Where("id = ?", demo1.Id).Update(demo1); err != nil {
		th.log.Errorw("Demo1Dao.Update", "err", err)
		return respx.InternalErrT
	}
	return nil
}

func (th *demo1Dao) DetailById(demo1 *models.Demo1) (*models.Demo1, *respx.T) {
	if _, err := th.mysql.Where("id = ?", demo1.Id).Get(demo1); err != nil {
		th.log.Errorw("Demo1Dao.Detail", "err", err)
		return nil, respx.InternalErrT
	}
	return demo1, nil
}

func (th *demo1Dao) List() (*respx.List, *respx.T) {
	var (
		demo1s = make([]*models.Demo1, 0)
		count int64
		err   error
	)

	if count, err = th.mysql.FindAndCount(&demo1s); err != nil {
		th.log.Errorw("Demo1Dao.List", "err", err)
		return nil, respx.InternalErrT
	}
	return &respx.List{Count: count, List: demo1s}, nil
}

