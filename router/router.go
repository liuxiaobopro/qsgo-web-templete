package router

import (
	"net/http"
	"qsgo-web-templete/global"

	"github.com/gin-gonic/gin"
)

func AddRouter(r *gin.Engine) {
	if global.Conf.Runmode != global.PROD {
		r.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello, qsgo-web-templete",
			})
		})
		r.GET("/config", func(c *gin.Context) {
			var conf = global.Conf
			c.JSON(http.StatusOK, conf)
		})
	}
}
