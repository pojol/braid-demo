package errcode

import "github.com/pojol/braid/lib/errcode"

var (
	Unknow          = func(args ...interface{}) errcode.Code { return errcode.Add(-1, " 未知错误", args...) }
	ErrIOReadALL    = func(args ...interface{}) errcode.Code { return errcode.Add(30001, " 消息包读取错误", args...) }
	ErrProtoUnmarsh = func(args ...interface{}) errcode.Code { return errcode.Add(30002, " 协议解析失败", args...) }
	ErrMongoCmd     = func(args ...interface{}) errcode.Code { return errcode.Add(30003, " Mongo指令执行失败", args...) }
)
