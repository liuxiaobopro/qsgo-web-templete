package global

import (
	"os"

	"qsgo-web-templete/config"

	otherx "github.com/liuxiaobopro/gobox/other"
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
	daoMysqlPath string
	projectName  string
)

func init() {
	// 获取项目根目录
	rootPath, _ = os.Getwd()
	daoPath = rootPath + "/dao"
	daoMysqlPath = rootPath + "/dao/mysql"
}

func GetRootPath() string {
	return rootPath
}

func GetDaoPath() string {
	return daoPath
}

func GetDaoMysqlPath() string {
	return daoMysqlPath
}

func GetProjectName() string {
	if projectName == "" {
		if p, err := otherx.GetProjectName(rootPath); err != nil {
			panic(err)
		} else {
			projectName = p
			return projectName
		}
	}
	return projectName
}
