package middleware

import (
	"net/http"
	"qsgo-web-templete/global"

	"github.com/gin-gonic/gin"
	respx "github.com/liuxiaobopro/gobox/resp"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.ZapS.Errorf("[middlewares recover] panic:%v", err)
				c.JSON(http.StatusInternalServerError, respx.InternalErrT)
				c.Abort()
			}
		}()
		c.Next()
	}
}
