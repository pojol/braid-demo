package events

import (
	"braid-demo/models/gameproto"
	"braid-demo/models/session"
	"bytes"
	"context"
	"encoding/binary"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/lib/log"
	"github.com/pojol/braid/router"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func MakeWebsocketNotify(actorCtx context.Context) core.IChain {

	return &actor.DefaultChain{

		Handler: func(mw *router.MsgWrapper) error {
			state := core.GetState(actorCtx).(*session.State)
			conn, ok := state.GetSession(mw.Res.Header.Token)
			if !ok {
				log.Warn("websocket get session err, token : %v", mw.Res.Header.Token)
				return nil
			}

			// Get a buffer from the pool
			buf := bufferPool.Get().(*bytes.Buffer)
			buf.Reset() // Clear the buffer for reuse

			resHeader := gameproto.MsgHeader{
				Event: mw.Res.Header.Event,
				Token: mw.Req.Header.Token,
			}

			resHeaderByt, _ := proto.Marshal(&resHeader)
			binary.Write(buf, binary.LittleEndian, uint16(len(resHeaderByt)))
			binary.Write(buf, binary.LittleEndian, resHeaderByt)
			binary.Write(buf, binary.LittleEndian, mw.Res.Body)

			err := conn.WriteMessage(websocket.BinaryMessage, buf.Bytes())
			if err != nil {
				return err
			}

			// Put the buffer back in the pool immediately after use
			bufferPool.Put(buf)

			return nil
		},
	}

}
