package global

import (
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
)
