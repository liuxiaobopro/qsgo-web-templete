package middleware

import (
	"net/http"
	"qsgo-web-templete/global"

	"github.com/gin-gonic/gin"
	replyx "github.com/liuxiaobopro/gobox/reply"
)

func Recover(debug bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.Errorf(c, "middleware recover err: %v", err)

				if debug {
					c.JSON(http.StatusInternalServerError, &replyx.T{
						Code: replyx.InternalErrT.Code,
						Msg:  err.(error).Error(),
					})
				} else {
					c.JSON(http.StatusInternalServerError, replyx.InternalErrT)
				}
				c.Abort()
			}
		}()
		c.Next()
	}
}
