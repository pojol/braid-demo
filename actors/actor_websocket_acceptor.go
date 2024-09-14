package actors

import (
	"braid-demo/constant"
	"braid-demo/events"
	"braid-demo/models/gameproto"
	"braid-demo/models/session"
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"strconv"
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

func NewWSAcceptorActor(p *core.CreateActorParm) core.IActor {

	return &websocketAcceptorActor{
		Runtime: &actor.Runtime{Id: p.ID, Ty: constant.ActorWebsoketAcceptor, Sys: p.Sys},
		echoptr: echo.New(),
		Port:    p.Options["port"].(string),
		state: &session.State{
			SessionMap: make(map[string]*websocket.Conn),
		},
	}
}

func (a *websocketAcceptorActor) Init() {
	a.Runtime.Init()

	recovercfg := middleware.DefaultRecoverConfig
	recovercfg.LogErrorFunc = func(c echo.Context, err error, stack []byte) error {
		log.Error("recover err %v stack %v", err.Error(), string(stack))
		return nil
	}
	a.echoptr.Use(middleware.RecoverWithConfig(recovercfg))
	a.echoptr.Use(middleware.CORS())

	a.echoptr.GET("/ws", a.received)
	a.RegisterEvent(events.EvWebsoketNotify, events.MakeWebsocketNotify(a.state))
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

		if header.Event == events.EvLogin {
			actorid = def.SymbolLocalFirst
			actorty = constant.ActorLogin
		} else if header.Event == events.EvChatSendMessage {
			actorid = def.SymbolLocalFirst
			//actorty = constant.ActorChat
		} else {
			eid, err := token.Parse(header.Token)
			if err != nil {
				fmt.Println(header.Token, "token.parse err", err.Error())
				continue
			}

			actorid = eid
			actorty = constant.ActorUser
			bh.Token = header.Token
		}

		sendmsg := router.NewMsg().WithReqHeader(bh).WithReqBody(msg[2+headerlen:]).Build()

		// Perform the system call with the timeout context
		err = a.Call(ctx, router.Target{
			ID: actorid,
			Ty: actorty,
			Ev: header.Event,
		}, sendmsg)
		if err != nil {
			// Handle the error, such as logging or returning a response
			log.Warn("system call actor:%v ty:%v event:%v err %v", actorid, actorty, header.Event, err)
			continue
		}

		// Get a buffer from the pool
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Reset() // Clear the buffer for reuse

		if header.Event == events.EvLogin {
			userToken = "token"
			a.state.AddSession("token", ws)
		}

		if _, ok := sendmsg.Res.Header.Custom["msgid"]; ok {
			resMsgID := sendmsg.Res.Header.Custom["msgid"]
			u16msgid, _ := strconv.Atoi(resMsgID)
			binary.Write(buf, binary.LittleEndian, uint16(u16msgid))
		}

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
		panic(fmt.Errorf("echo start err: %w", err))
	}
}

func (a *websocketAcceptorActor) Exit() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := a.echoptr.Shutdown(ctx); err != nil {
		log.Error("failed to shutdown server: %v", err)
	}

	a.Runtime.Exit()
}
