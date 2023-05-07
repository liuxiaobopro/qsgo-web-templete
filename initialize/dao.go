package initialize

import (
	"qsgo-web-templete/global"
)

func Dao() {
	gxd := NewGenXormDao(global.DB, global.GetDaoPath(), global.Conf.Mysql.Prefix)
	if err := gxd.Gen(); err != nil {
		panic(err)
	}
}
