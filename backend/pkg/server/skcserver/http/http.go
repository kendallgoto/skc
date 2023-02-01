package http

import (
	"net/http"

	"github.com/brpaz/echozap"
	"github.com/kendallgoto/skc/pkg/execute"
	"github.com/kendallgoto/skc/pkg/lang"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

const (
	routeNewCodePost = "/"
)

type newCodeRequest struct {
	Code string `json:"code"`
}

type newCodeResponse struct {
	Input     string `json:"input"`
	Output    string `json:"output"`
	Execution string `json:"execution"`
}

func CreateServer(logger *zap.SugaredLogger) (*echo.Echo, error) {
	e := echo.New()
	e.HideBanner = true // echo labstack banner
	// use central logger
	e.Use(echozap.ZapLogger(logger.Desugar()))
	// graceful recovery
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			_ = c.NoContent(http.StatusInternalServerError)
			logger.With(zap.Stack("stack")).Fatalln(err)
			return err
		},
	}))
	// debug dump
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		logger.Infof(string(resBody))
	}))
	addRoutes(e, logger)
	return e, nil
}

func addRoutes(e *echo.Echo, logger *zap.SugaredLogger) {
	e.POST(routeNewCodePost, func(c echo.Context) error {
		// retrieve code from body
		req := new(newCodeRequest)
		if err := c.Bind(req); err != nil {
			return err
		}
		result, err := lang.Parse(req.Code)
		if err != nil {
			return err
		}
		execution, err := execute.Run(result)
		if err != nil {
			return err
		}
		responseObj := newCodeResponse{
			Input:     req.Code,
			Output:    result,
			Execution: execution,
		}
		c.JSON(http.StatusOK, responseObj)
		return nil
	})
}
