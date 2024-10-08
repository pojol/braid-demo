package actors

import (
	"braid-demo/events"
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

type httpAcceptorActor struct {
	*actor.Runtime
	echoptr *echo.Echo
	Port    string
}

func NewHttpAcceptorActor(p core.IActorBuilder) core.IActor {
	return &httpAcceptorActor{
		Runtime: &actor.Runtime{Id: p.GetID(), Ty: p.GetType(), Sys: p.GetSystem()},
		echoptr: echo.New(),
		Port:    p.GetOpt("port").(string),
	}
}

func (a *httpAcceptorActor) Init(ctx context.Context) {
	a.Runtime.Init(ctx)

	a.RegisterEvent(events.EvHttpHello, events.HttpHello)

	recovercfg := middleware.DefaultRecoverConfig
	recovercfg.LogErrorFunc = func(c echo.Context, err error, stack []byte) error {
		log.ErrorF("recover err %v stack %v", err.Error(), string(stack))
		return nil
	}
	a.echoptr.Use(middleware.RecoverWithConfig(recovercfg))
	a.echoptr.Use(middleware.CORS())

	a.echoptr.POST("/*", func(c echo.Context) error {

		bts, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.Blob(http.StatusBadRequest, echo.MIMEApplicationJSON, nil)
		}

		msg := router.NewMsgWrap(context.TODO()).WithReqBody(bts).Build()

		err = a.Call(router.Target{ID: a.Id, Ty: a.Ty, Ev: strings.TrimPrefix(c.Request().URL.Path, "/")},
			msg,
		)
		if err != nil {
			log.WarnF("call %v err %v", c.Request().Method, err.Error())
		}

		c.Blob(http.StatusOK, echo.MIMEApplicationJSON, msg.Res.Body)
		return nil
	})
}

func (a *httpAcceptorActor) Update() {
	go a.Runtime.Update()

	err := a.echoptr.Start(":" + a.Port)
	if err != nil {
		fmt.Println(fmt.Errorf("echo start err: %w", err))
	}
}

func (a *httpAcceptorActor) Exit() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := a.echoptr.Shutdown(ctx); err != nil {
		log.ErrorF("failed to shutdown server: %v", err)
	}

	a.Runtime.Exit()
}
