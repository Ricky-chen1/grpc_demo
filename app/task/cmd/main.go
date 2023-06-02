package main

import (
	"grpc_demo/app/task/internal/db"
	"grpc_demo/app/task/internal/service"
	"grpc_demo/conf"
	"grpc_demo/idl/pb/task"
	"grpc_demo/pkg/discovery"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	srvName = "task"
)

func Init() {
	conf.Init(srvName)
	db.Init()
}

// task moudle grpc server start
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
	//注册备忘录模块服务
	task.RegisterTaskServiceServer(grpcServer, service.NewTaskService())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start %v", err)
		panic(err)
	}
}
