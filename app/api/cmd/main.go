package main

import (
	"grpc_demo/app/api/router"
	"grpc_demo/app/api/rpc"
	"grpc_demo/conf"
	"log"
)

func Init() {
	conf.Init()
	rpc.Init()
}

// start gateway server
func main() {
	Init()

	r := router.NewRouter()
	grpcAddress := conf.C.Services["api"].Addr[0]
	if err := r.Run(grpcAddress); err != nil {
		log.Fatal("gateway server start error")
		panic(err)
	}

}
