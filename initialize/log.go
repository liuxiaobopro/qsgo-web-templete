package initialize

import (
	"qsgo-web-templete/global"

	"go.uber.org/zap"
)

func Log() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()
	global.ZapS = sugar

	global.ZapS.Info("log init success")
}
