package skcserver

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"time"

	gohttp "net/http"

	"github.com/kendallgoto/skc/pkg/server/skcserver/http"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Start(logger *zap.SugaredLogger) error {
	// Open DB connection
	e, err := http.CreateServer(logger)
	if err != nil {
		return err
	}
	go func() {
		if err := e.Start(":8080"); err != nil && !errors.Is(err, gohttp.ErrServerClosed) {
			logger.Errorln("http server died", err)
			if err != nil {
				panic(err)
			}
		}
	}()
	err = listenForShutdown(e)
	if err != nil {
		return err
	}
	return nil
}
func listenForShutdown(echoServer *echo.Echo) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	echoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := echoServer.Shutdown(echoCtx); err != nil {
		return err
	}
	return nil
}
