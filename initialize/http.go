package initialize

import (
	"context"
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
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

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
