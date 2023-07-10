package router

import (
	"github.com/gin-gonic/gin"
)

func AddRouter(r *gin.Engine) {
	r.Group("/api")
}
