package main

import (
	"grpc_demo/app/user/internal/db"
	"grpc_demo/app/user/internal/service"
	"grpc_demo/conf"
	"grpc_demo/idl/pb/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

// user moudle grpc server start
func main() {
	//数据库初始化
	db.Init()
	conf.Init()

	grpcAddress := conf.C.Services["user"].Addr[0]
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
		panic(err)
	}

	grpcServer := grpc.NewServer()

	//注册用户模块服务
	user.RegisterUserServiceServer(grpcServer, service.NewUserService())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start %v", err)
		panic(err)
	}
}
