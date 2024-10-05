package middleware

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/pojol/braid/core/actor"
	"github.com/pojol/braid/lib/log"
	"github.com/pojol/braid/router"
)

type MessageUnpackCfg[T any] struct {
	MsgTy T
	Msg   interface{}
}

func MessageUnpack[T any](cfg *MessageUnpackCfg[T]) actor.EventHandler {
	return func(msg *router.MsgWrapper) error {
		var msgInstance proto.Message
		msgType := reflect.TypeOf(cfg.MsgTy)

		// 检查是否为指针类型
		if msgType.Kind() == reflect.Ptr {
			msgInstance = reflect.New(msgType.Elem()).Interface().(proto.Message)
		} else {
			msgInstance = reflect.New(msgType).Interface().(proto.Message)
		}

		// 解析消息
		err := proto.Unmarshal(msg.Req.Body, msgInstance)
		if err != nil {
			return fmt.Errorf("unpack msg err %v", err.Error())
		}

		// 打印消息类型和字段信息
		log.InfoF("[req event] actor_id : %s actor_ty : %s event : %s: params : %s",
			msg.Req.Header.TargetActorID,
			msg.Req.Header.TargetActorType,
			reflect.TypeOf(msgInstance).Elem().Name(), printMessageFields(msgInstance))

		// 将解析后的消息赋值给 cfg.Msg
		cfg.Msg = msgInstance

		return nil
	}
}

func printMessageFields(msg proto.Message) string {
	v := reflect.ValueOf(msg).Elem()
	t := v.Type()
	var fields []string

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// 跳过未导出的字段
		if !fieldType.IsExported() {
			continue
		}

		fields = append(fields, fmt.Sprintf("%s: %v", fieldType.Name, field.Interface()))
	}

	return strings.Join(fields, ", ")
}
