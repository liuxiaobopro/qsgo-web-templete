package initialize

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"qsgo-web-templete/global"
	"qsgo-web-templete/router"

	"github.com/gin-gonic/gin"
)

func Http() {
	r := gin.Default()

	router.AddRouter(r)

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
	global.ZapS.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.ZapS.Error("Server Shutdown:", err)
	}
	global.ZapS.Info("Server exiting")
}
