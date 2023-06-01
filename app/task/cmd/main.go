package main

import (
	"grpc_demo/app/task/internal/db"
	"grpc_demo/app/task/internal/service"
	"grpc_demo/conf"
	"grpc_demo/idl/pb/task"
	"log"
	"net"

	"google.golang.org/grpc"
)

func Init(){
     	conf.Init()
     	db.Init()
}
// task moudle grpc server start
func main() {
	Init()
	
	grpcAddress := conf.C.Services["task"].Addr[0]
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
