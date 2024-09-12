package actors

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/lib/log"
	"github.com/pojol/braid/router"
)

type httpHelloActor struct {
	*actor.Runtime
	echoptr *echo.Echo
	Port    string
}

func NewHttpHelloActor(p *core.CreateActorParm) core.IActor {
	return &httpHelloActor{
		Runtime: &actor.Runtime{Id: p.ID, Ty: "httpHelloActor", Sys: p.Sys},
		echoptr: echo.New(),
		Port:    "8008",
	}
}

func (a *httpHelloActor) Init() {
	a.Runtime.Init()

	recovercfg := middleware.DefaultRecoverConfig
	recovercfg.LogErrorFunc = func(c echo.Context, err error, stack []byte) error {
		log.Error("recover err %v stack %v", err.Error(), string(stack))
		return nil
	}
	a.echoptr.Use(middleware.RecoverWithConfig(recovercfg))
	a.echoptr.Use(middleware.CORS())

	a.echoptr.POST("/*", func(c echo.Context) error {

		bts, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.Blob(http.StatusBadRequest, echo.MIMEApplicationJSON, nil)
		}

		msg := router.NewMsg().WithReqBody(bts).Build()

		err = a.Call(c.Request().Context(),
			router.Target{ID: a.Id, Ty: a.Ty, Ev: strings.TrimPrefix(c.Request().URL.Path, "/")},
			msg,
		)
		if err != nil {
			log.Warn("call %v err %v", c.Request().Method, err.Error())
		}

		c.Blob(http.StatusOK, echo.MIMEApplicationJSON, msg.Res.Body)
		return nil
	})
}

func (a *httpHelloActor) Update() {
	go a.Runtime.Update()

	err := a.echoptr.Start(":" + a.Port)
	if err != nil {
		panic(fmt.Errorf("echo start err: %w", err))
	}
}

func (a *httpHelloActor) Exit() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := a.echoptr.Shutdown(ctx); err != nil {
		log.Error("failed to shutdown server: %v", err)
	}

	a.Runtime.Exit()
}
