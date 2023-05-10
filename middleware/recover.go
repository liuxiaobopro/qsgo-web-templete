package middleware

import (
	"net/http"
	"qsgo-web-templete/global"

	"github.com/gin-gonic/gin"
	replyx "github.com/liuxiaobopro/gobox/reply"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.ZapS.Errorf("[middlewares recover] panic:%v", err)
				c.JSON(http.StatusInternalServerError, replyx.InternalErrT)
				c.Abort()
			}
		}()
		c.Next()
	}
}
