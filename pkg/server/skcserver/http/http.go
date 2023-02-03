package http

import (
	"fmt"
	"html/template"
	htmltemplate "html/template"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/brpaz/echozap"
	"github.com/kendallgoto/skc/pkg/execute"
	"github.com/kendallgoto/skc/pkg/lang"
	"github.com/kendallgoto/skc/pkg/tour"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

const (
	routeNewCodePost = "/"
	routeHomeGet     = "/"
	routeTourHomeGet = "/tour"
	routeTourPageGet = "/tour/:page"
	routeTourGet     = "/tour/:page/:short"
)

type newCodeRequest struct {
	Code string `json:"code"`
}

type newCodeResponse struct {
	Input     string            `json:"input"`
	Output    string            `json:"output"`
	Execution string            `json:"execution"`
	Errors    []lang.ParseError `json:"errors"`
}

type TemplateRenderer struct {
	template *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}

func CreateServer(logger *zap.SugaredLogger) (*echo.Echo, error) {
	e := echo.New()
	e.HideBanner = true // echo labstack banner
	// use central logger
	e.Use(echozap.ZapLogger(logger.Desugar()))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	// graceful recovery
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			_ = c.NoContent(http.StatusInternalServerError)
			logger.With(zap.Stack("stack")).Fatalln(err)
			return err
		},
	}))
	e.Pre(middleware.RemoveTrailingSlash())
	// debug dump
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		logger.Infof(string(resBody))
	}))
	addRoutes(e, logger)

	t, err := htmltemplate.New("template").Parse(tourTemplateHTML)
	if err != nil {
		return nil, err
	}
	renderer := &TemplateRenderer{
		template: t,
	}
	e.Renderer = renderer
	return e, nil
}
func addRoutes(e *echo.Echo, logger *zap.SugaredLogger) {
	e.POST(routeNewCodePost, func(c echo.Context) error {
		// retrieve code from body
		req := new(newCodeRequest)
		if err := c.Bind(req); err != nil {
			return err
		}
		result, parseErrors, err := lang.Parse(req.Code)
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
			Errors:    parseErrors,
		}
		c.JSON(http.StatusOK, responseObj)
		return nil
	})
	e.GET(routeHomeGet, func(c echo.Context) error {
		c.Redirect(301, "/tour")
		return nil
	})
	e.GET(routeTourHomeGet, func(c echo.Context) error {
		c.Redirect(307, "/tour/1")
		return nil
	})
	e.GET(routeTourPageGet, func(c echo.Context) error {
		pageNum, err := strconv.Atoi(c.Param("page"))
		if err != nil {
			return err
		}
		return renderPage(c, pageNum, "")
	})
	//e.GET(routeTourGet, echo.WrapHandler(fserv))

	e.GET(routeTourGet, func(c echo.Context) error {
		pageNum, err := strconv.Atoi(c.Param("page"))
		if err != nil {
			return err
		}
		short := c.Param("short")
		return renderPage(c, pageNum, short)
	})

	e.GET("/static/*", echo.WrapHandler(http.FileServer(http.FS(staticFiles))))
}

func renderPage(c echo.Context, pageNum int, short string) error {
	if pageNum > len(tour.TourItems) {
		return echo.NewHTTPError(404, "Page doesn't exist")
	}
	toRender := tour.TourItems[pageNum-1]
	var prevItem, nextItem string
	if pageNum > 1 {
		item := tour.TourItems[pageNum-2]
		prevItem = fmt.Sprintf("/tour/%d/%s", pageNum-1, item.ShortName)
	}
	if pageNum < len(tour.TourItems) {
		item := tour.TourItems[pageNum]
		nextItem = fmt.Sprintf("/tour/%d/%s", pageNum+1, item.ShortName)
	}
	if !strings.EqualFold(short, toRender.ShortName) {
		return c.Redirect(301, fmt.Sprintf("/tour/%d/%s", pageNum, toRender.ShortName))
	}
	return c.Render(http.StatusOK, "template", map[string]interface{}{
		"num":   pageNum,
		"title": toRender.Title,
		"short": toRender.ShortName,
		"body":  htmltemplate.HTML(toRender.Body),
		"code":  toRender.Code,
		"prev":  prevItem,
		"next":  nextItem,
	})
}
