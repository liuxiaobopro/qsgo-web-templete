package initialize

import (
	"qsgo-web-templete/global"

	"go.uber.org/zap"
)

func Log() {
	var (
		logger *zap.Logger
		err    error
	)
	if global.Conf.Runmode == global.PROD {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()
	global.ZapS = sugar
}
