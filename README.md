# grpc-todolist

## 项目结构

### 整体分层
```
.
├── Makefile
├── README.md
├── app                 # 应用程序
│   ├── api             # 网关
│   ├── task            # 任务模块
│   └── user            # 用户模块
├── config
│   ├── config.go
│   ├── config.aml      # 配置项
├── go.mod
├── go.sum
├── postman             # 接口测试配置
├── idl
│   ├── pb
│   ├── task.proto
│   └── user.proto
└── pkg
    ├── errno           # 自定义错误
    └── util            # 工具项
```
### 微服务
```
.
├── cmd                     # 启动入口
├── internal                # (不对外暴露)
    ├── db                  # 数据库初始化及操作  
    ├── model               # 数据库映射
    │   └── model.go
    ├── pack                # 打包
    │   └── pack.go
    └── service             # 业务逻辑
        └── service.go
```
### 网关
```
.
├── cmd                 # 网关启动入口    
├── handler             # 请求处理
│   ├── handler.go
│   ├── task.go
│   └── user.go
├── middleware          # gin中间件
├── router              # 路由模块
└── rpc                 # rpc调用  

```
### 手动启动  

1.保证mysql,etcd等服务活跃,在每个模块下的cmd执行以下命令  

```go
go run main.go
```

