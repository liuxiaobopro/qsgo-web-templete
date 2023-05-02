package global

import (
	"qsgo-web-templete/config"

	"go.uber.org/zap"
)

const (
	DEV  = "dev"
	PROD = "prod"
)

var (
	ZapS *zap.SugaredLogger
	Conf config.Conf
)
