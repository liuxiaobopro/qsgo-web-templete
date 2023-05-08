package initialize

import (
	"qsgo-web-templete/global"

	xormx "github.com/liuxiaobopro/gobox/xorm"
)

func Dao() {
	gxd := xormx.NewGenXormDao(global.DB, global.GetDaoMysqlPath(), global.GetProjectName(), xormx.WithPrefix(global.Conf.Mysql.Prefix))
	if err := gxd.Gen(); err != nil {
		panic(err)
	}
}
