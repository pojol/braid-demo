// Code generated by go generate; DO NOT EDIT.

package template

// Actor 类型模版定义

// 配置字段说明：
// name: Actor 的类型名称，用于在代码中识别和创建 Actor
// id: Actor 的唯一标识符，可以包含 {nodeid} 占位符，将被替换为实际的节点 ID
// unique: 布尔值，表示该 Actor类型 在节点中是否唯一
// weight: 整数，表示 Actor 的权重，用于负载均衡
// limit: 整数，表示 Actor 的全局数量限制，0 表示无限制
// options: 可选的 Actor 特定配置项

// weight 我们可以默认设计一个权重计算方法
// 比如 2c4g 的pod， 可以等于 2x4x1000 = 8000
// 在这个基础上，我们可以通过压测数据，去规划调整 actor 的权重值

const (
  // WebSocket 接收器
  // 用于接受 WebSocket 连接的 Actor
  // 选项:
  //   - port: WebSocket 服务器端口
    ACTOR_WEBSOCKET_ACCEPTOR = "WEBSOCKET_ACCEPTOR"

    ACTOR_HTTP_ACCEPTOR = "HTTP_ACCEPTOR"

  // 登录处理
  // 处理用户登录请求的 Actor
    ACTOR_LOGIN = "LOGIN"

    ACTOR_USER = "USER"

  // 动态选择器
  // 用于动态选择其他 Actor 的 Actor
    ACTOR_DYNAMIC_PICKER = "DYNAMIC_PICKER"

  // 动态注册器
  // 用于动态注册其他 Actor 的 Actor
    ACTOR_DYNAMIC_REGISTER = "DYNAMIC_REGISTER"

// actor 控制
    ACTOR_CONTROL = "CONTROL"

  // 聊天
  // 选项:
  //   - channel: 聊天频道名称
    ACTOR_CHAT = "CHAT"

  // 聊天路由
  // 用于路由聊天消息的 Actor
    ACTOR_ROUTER_CHAT = "ROUTER_CHAT"

)