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
    ├── discovery       # etcd服务注册与发现中心的实现
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

1.保证mysql,etcd等服务活跃,在每个模块的cmd下执行  

```go
go run main.go
```
### 遇到的困难或犯的错误  
1.grpc官方英文文档有点难懂  
2.只是大概能明白grpc服务和etcd注册中心的交互过程，最后还是copy了demo的实现  
3.很蠢的错误:对空指针进行了操作，导致了一些error  

### 项目后续
1.为主要的逻辑添加单元测试  
2.docker环境部署，Makefile精简化指令  
...  
