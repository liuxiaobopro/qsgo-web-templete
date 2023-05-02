package initialize

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"qsgo-web-templete/global"

	"github.com/gin-gonic/gin"
)

func Http() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		var conf = global.Conf
		c.JSON(http.StatusOK, conf)
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", global.Conf.Http.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	fmt.Println("Server Run:", fmt.Sprintf("http://localhost:%d", global.Conf.Http.Port))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
