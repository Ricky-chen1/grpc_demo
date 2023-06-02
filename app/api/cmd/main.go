package main

import (
	"grpc_demo/app/api/router"
	"grpc_demo/app/api/rpc"
	"grpc_demo/conf"
	"log"
)

const (
	srvName = "api"
)

func Init() {
	conf.Init(srvName)
	rpc.Init()
}

// start gateway server
func main() {
	Init()

	r := router.NewRouter()
	grpcAddress := conf.Service.Addr

	if err := r.Run(grpcAddress); err != nil {
		log.Fatal("gateway server start error")
		panic(err)
	}

}
