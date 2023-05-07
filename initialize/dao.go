package initialize

import (
	"qsgo-web-templete/global"
)

func Dao() {
	gxd := NewGenXormDao(global.DB, global.GetDaoMysqlPath(), global.Conf.Mysql.Prefix, global.GetProjectName())
	if err := gxd.Gen(); err != nil {
		panic(err)
	}
}
