package actors

import (
	"braid-demo/config"
	"braid-demo/events"
	"braid-demo/models/gameproto"
	"braid-demo/models/session"
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/def"
	"github.com/pojol/braid/lib/log"
	"github.com/pojol/braid/lib/token"
	"github.com/pojol/braid/router"
)

type websocketAcceptorActor struct {
	*actor.Runtime
	echoptr *echo.Echo
	Port    string

	state *session.State
}

var (
	upgrader = websocket.Upgrader{}
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func NewWSAcceptorActor(p core.IActorBuilder) core.IActor {

	echoptr := echo.New()
	echoptr.HideBanner = true

	return &websocketAcceptorActor{
		Runtime: &actor.Runtime{Id: p.GetID(), Ty: p.GetType(), Sys: p.GetSystem()},
		echoptr: echoptr,
		Port:    p.GetOpt("port").(string),
		state: &session.State{
			SessionMap: make(map[string]*websocket.Conn),
		},
	}
}

func (a *websocketAcceptorActor) Init(ctx context.Context) {
	a.Runtime.Init(ctx)

	a.Context().WithValue(events.SessionState{}, a.state)

	recovercfg := middleware.DefaultRecoverConfig
	recovercfg.LogErrorFunc = func(c echo.Context, err error, stack []byte) error {
		log.ErrorF("recover err %v stack %v", err.Error(), string(stack))
		return nil
	}
	a.echoptr.Use(middleware.RecoverWithConfig(recovercfg))
	a.echoptr.Use(middleware.CORS())

	a.echoptr.GET("/ws", a.received)

	a.RegisterEvent(events.EvLogin, events.MakeWSLogin)
	a.RegisterEvent(events.EvWebsoketNotify, events.MakeWebsocketNotify)
}

func (a *websocketAcceptorActor) received(c echo.Context) error {

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	var userToken string
	defer func() {
		ws.Close()
		if userToken != "" {
			a.state.RemoveSession(userToken)
		}
	}()

	for {
		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("read msg err", err.Error())
			break
		}

		if len(msg) < 2 {
			fmt.Println("message too small was read", len(msg))
			continue
		}

		headerlen := binary.LittleEndian.Uint16(msg[:2])

		// 检查消息是否足够长
		if len(msg) < int(2+headerlen) {
			fmt.Printf("message too short for header: expected %d, got %d\n", 2+headerlen, len(msg))
			continue
		}

		header := &gameproto.MsgHeader{}
		err = proto.Unmarshal(msg[2:2+headerlen], header)
		if err != nil {
			fmt.Println("unmarshal proto header err")
			continue
		}

		// Create a context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		bh := &router.Header{}
		var actorid, actorty string

		switch header.Event {
		case events.EvLogin:
			actorid = def.SymbolLocalFirst
			actorty = config.ACTOR_LOGIN
		case events.EvChatSendMessage:
			actorid = def.SymbolLocalFirst
			actorty = config.ACTOR_ROUTER_CHAT
		default:
			eid, err := token.Parse(header.Token)
			if err != nil {
				fmt.Println(header.Token, "token.parse err", err.Error())
				continue
			}

			actorid = eid
			actorty = config.ACTOR_USER
			bh.Token = header.Token
		}

		sendmsg := router.NewMsgWrap(ctx).WithReqHeader(bh).WithReqBody(msg[2+headerlen:]).Build()

		// Perform the system call with the timeout context
		err = a.Call(router.Target{
			ID: actorid,
			Ty: actorty,
			Ev: header.Event,
		}, sendmsg)
		if err != nil {

			// Handle the error, such as logging or returning a response
			log.WarnF("system call actor:%v ty:%v event:%v err %v", actorid, actorty, header.Event, err)
			continue
		}

		// Get a buffer from the pool
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Reset() // Clear the buffer for reuse

		if header.Event == events.EvLogin {
			userToken = sendmsg.Res.Header.Token
			a.state.AddSession(userToken, ws)
		}

		resHeader := gameproto.MsgHeader{
			Event: header.Event,
			Token: userToken,
		}

		resHeaderByt, _ := proto.Marshal(&resHeader)
		binary.Write(buf, binary.LittleEndian, uint16(len(resHeaderByt)))
		binary.Write(buf, binary.LittleEndian, resHeaderByt)
		binary.Write(buf, binary.LittleEndian, sendmsg.Res.Body)

		err = ws.WriteMessage(websocket.BinaryMessage, buf.Bytes())

		// Put the buffer back in the pool immediately after use
		bufferPool.Put(buf)

		if err != nil {
			fmt.Println("handle write err", err.Error())
			break
		}
	}

	return nil
}

func (a *websocketAcceptorActor) Update() {
	go a.Runtime.Update()

	err := a.echoptr.Start(":" + a.Port)
	if err != nil {
		log.InfoF("echo server exit %v", err.Error())
	}
}

func (a *websocketAcceptorActor) Exit() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := a.echoptr.Shutdown(ctx); err != nil {
		log.ErrorF("failed to shutdown server: %v", err)
	}

	a.Runtime.Exit()
}
