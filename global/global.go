package global

import (
	"qsgo-web-templete/config"

	"go.uber.org/zap"
)

var (
	ZapS *zap.SugaredLogger
	Conf config.Conf
)
