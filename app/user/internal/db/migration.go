package db

import (
	"grpc_demo/app/user/internal/model"
	"log"
)

func Migration() {
	if err := DB.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("user autoMigration failed %v", err)
		panic(err)
	}
}
