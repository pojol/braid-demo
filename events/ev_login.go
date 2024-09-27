package events

import (
	"braid-demo/constant"
	"braid-demo/errcode"
	"braid-demo/middleware"
	"braid-demo/models/gameproto"
	"braid-demo/models/user"
	"context"
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"github.com/pojol/braid/3rd/mgo"
	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/lib/log"
	"github.com/pojol/braid/lib/token"
	"github.com/pojol/braid/router"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeWSLogin(actorCtx context.Context) core.IChain {
	unpackCfg := &middleware.MessageUnpackCfg[*gameproto.LoginReq]{}

	return &actor.DefaultChain{
		Before: []actor.EventHandler{middleware.MessageUnpack(unpackCfg)},
		Handler: func(mw *router.MsgWrapper) error {
			req := unpackCfg.Msg.(*gameproto.LoginReq)
			resp := &gameproto.LoginResp{}
			sys := core.GetSystem(actorCtx)

			// 检查 db 是否存在， 创建 / 登录
			e := &user.EntityWrapper{
				User: &user.EntityUserModule{},
			}
			create := false

			err := mgo.Collection(constant.MongoDatabase, constant.MongoCollection).FindOne(mw.Ctx, bson.M{"openid": req.Uid}).Decode(e)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					create = true
				} else {
					return errcode.ErrMongoCmd(req.Uid, err)
				}
			}

			if create { // 创建账号

				e.ID = uuid.NewString()
				e.User.Token, _ = token.Create(e.ID)
				e.TimeInfo = &user.EntityTimeInfoModule{
					ID:         e.ID,
					CreateTime: time.Now().Unix(),
				}
				e.Bag = &user.EntityBagModule{ID: e.ID}
				e.Airship = &user.EntityAirshipModule{ID: e.ID}

				_, err = mgo.Collection(constant.MongoDatabase, constant.MongoCollection).InsertOne(mw.Ctx, e)
				if err != nil {
					return errcode.ErrMongoCmd(err)
				}

			} else { // 刷新token
				newToken, err := token.Create(e.ID)
				if err != nil {
					return err
				}

				e.User.Token = newToken
				log.Info("user %v refresh token %v", e.ID, newToken)
			}

			mw.Req.Header.PrevActorType = constant.ActorLogin
			err = sys.Loader().Builder(constant.ActorUser).
				WithID(e.ID).
				WithOpt("gateActor", mw.Req.Header.OrgActorID).RegisterDynamically()
			if err != nil {
				fmt.Println("login ->", "regist actor err", err.Error())
				return err
			}

			resp.Uid = e.ID
			resp.Token = e.User.Token
			mw.Res.Header.Token = e.User.Token
			mw.Res.Body, _ = proto.Marshal(resp)

			return nil
		},
	}

}
