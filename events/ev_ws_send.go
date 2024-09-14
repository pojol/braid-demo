package events

import (
	"braid-demo/models/session"
	"bytes"
	"context"
	"encoding/binary"
	"strconv"
	"sync"

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

func MakeWebsocketNotify(state *session.State) core.IChain {

	return &actor.DefaultChain{

		Handler: func(ctx context.Context, mw *router.MsgWrapper) error {

			conn, ok := state.GetSession(mw.Req.Header.Token)
			if !ok {
				log.Warn("websocket get session err, token : %v", mw.Req.Header.Token)
				return nil
			}

			// Get a buffer from the pool
			buf := bufferPool.Get().(*bytes.Buffer)
			buf.Reset() // Clear the buffer for reuse

			if _, ok := mw.Res.Header.Custom["msgid"]; ok {
				resMsgID := mw.Res.Header.Custom["msgid"]
				u16msgid, _ := strconv.Atoi(resMsgID)
				binary.Write(buf, binary.LittleEndian, uint16(u16msgid))
			}

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
