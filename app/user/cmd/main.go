package main

import (
	"grpc_demo/app/user/internal/db"
	"grpc_demo/app/user/internal/service"
	"grpc_demo/conf"
	"grpc_demo/idl/pb/user"
	"grpc_demo/pkg/discovery"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	srvName = "user"
)

func Init() {
	conf.Init(srvName)
	db.Init()
}

// user moudle grpc server start
func main() {
	Init()

	//注册到etcd上
	etcdAddr := []string{conf.Etcd.Addr}
	register := discovery.NewRegister(etcdAddr, logrus.New())

	node := discovery.Server{
		Name: conf.Service.Name,
		Addr: conf.Service.Addr,
	}

	if _, err := register.Register(node, 10); err != nil {
		log.Fatalf("register service %s failed,err: %v", node.Name, err)
	}

	grpcAddress := conf.Service.Addr
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
