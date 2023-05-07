package global

import (
	"os"
	"qsgo-web-templete/config"

	"go.uber.org/zap"
	"xorm.io/xorm"
)

const (
	DEV  = "dev"
	PROD = "prod"
)

var (
	ZapS *zap.SugaredLogger
	Conf config.Conf
	DB   *xorm.Engine

	rootPath     string
	daoPath      string
	mysqlDaoPath string
)

func init() {
	// 获取项目根目录
	rootPath, _ = os.Getwd()
	daoPath = rootPath + "/dao"
	mysqlDaoPath = rootPath + "/dao/mysql"
}

func GetRootPath() string {
	return rootPath
}

func GetDaoPath() string {
	return daoPath
}

func GetMysqlDaoPath() string {
	return mysqlDaoPath
}
