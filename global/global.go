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

	rootPath string
	daoPath  string
)

func init() {
	// 获取项目根目录
	rootPath, _ = os.Getwd()
	daoPath = rootPath + "/dao"
}

func GetRootPath() string {
	return rootPath
}

func GetDaoPath() string {
	return daoPath
}
