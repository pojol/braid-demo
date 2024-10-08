package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var types = `
# Actor 类型定义

actor_types:
  # WebSocket 接收器
  # 用于接受 WebSocket 连接的 Actor
  # 选项:
  #   - port: WebSocket 服务器端口
- name: "WEBSOCKET_ACCEPTOR"
  id: "{NODE_ID}_websocket_acceptor"
  unique: true
  weight: 800
  limit: 1
  # options 部分被移除，将在 node conf.yml 中定义
  # - port 对外端口号
  # like:
  #options:
  #  port: "8008"

- name: "HTTP_ACCEPTOR"
  id: "{NODE_ID}_http_acceptor"
  unique: true
  weight: 800
  limit: 1
  # options 部分被移除，将在 node conf.yml 中定义
  # - port 对外端口号
  # like:
  #options:
  #  port: "8008"

  # 登录处理
  # 处理用户登录请求的 Actor
- name: "LOGIN"
  id: "{NODE_ID}_login"
  unique: false
  weight: 800
  limit: 2

- name: "USER"
  id: "{NODE_ID}_user"
  unique: false
  weight: 80
  limit: 10000

  # 动态选择器
  # 用于动态选择其他 Actor 的 Actor
- name: "DYNAMIC_PICKER"
  id: "{NODE_ID}_dynamic_picker"
  unique: true
  weight: 80
  limit: 10

  # 动态注册器
  # 用于动态注册其他 Actor 的 Actor
- name: "DYNAMIC_REGISTER"
  id: "{NODE_ID}_dynamic_register"
  unique: true
  weight: 80
  limit: 0

# actor 控制
- name: "CONTROL"
  id : "{NODE_ID}_control"
  unique : true
  weight : 40
  limit : 0

  # 全局聊天
  # 处理全局聊天消息的 Actor
  # 选项:
  #   - channel: 聊天频道名称
- name: "GLOBAL_CHAT"
  id: "{NODE_ID}_global_chat"
  unique: false
  weight: 3000
  limit: 1

# 私聊
- name: "PRIVATE_CHAT"
  id: "{NODE_ID}_private_chat"
  unique: false
  weight: 40
  limit: 10000

  # 聊天路由
  # 用于路由聊天消息的 Actor
- name: "ROUTER_CHAT"
  id: "{NODE_ID}_router_chat"
  unique: true
  weight: 80
  limit: 1
`

var conf = `
node:
  # 节点唯一标识符，可以通过环境变量 NODE_ID 传入
  id: "{NODE_ID}"
  weight: "{NODE_WEIGHT}"
  
  # Actor 配置列表
  actors:
    # WebSocket 接收器 Actor
    - name: "WEBSOCKET_ACCEPTOR"
      options:
        port: "8008"

    # 登录 Actor
    - name: "LOGIN"

    - name: "CONTROL"

    # 动态选择器 Actor
    - name: "DYNAMIC_PICKER"

    # 动态注册器 Actor
    - name: "DYNAMIC_REGISTER"

`

func TestConfig(t *testing.T) {
	t.Setenv("NODE_ID", "test_node_1")
	t.Setenv("NODE_WEIGHT", "5000")

	nodCfg, actorTypes, err := ParseConfigFromString(conf, types)
	assert.Equal(t, err, nil)

	assert.Equal(t, nodCfg.ID, "test_node_1")

	flag := false
	for _, v := range nodCfg.Actors {
		if v.Name == "WEBSOCKET_ACCEPTOR" {
			flag = true
		}
	}
	assert.Equal(t, flag, true)

	for _, v := range nodCfg.Actors {
		if v.Name == "WEBSOCKET_ACCEPTOR" {
			assert.Equal(t, v.Options, map[string]string{"port": "8008"})
		}
	}

	for _, v := range actorTypes {
		if v.Name == "WEBSOCKET_ACCEPTOR" {
			assert.Equal(t, v.Limit, 1)
			assert.Equal(t, v.Weight, 800)
			assert.Equal(t, v.Unique, true)
		}
	}
}
