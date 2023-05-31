package db

import (
	"grpc_demo/app/task/internal/model"
	"log"
)

func Migration() {
	if err := DB.AutoMigrate(&model.Task{}); err != nil {
		log.Fatalf("task autoMigration failed %v", err)
		panic(err)
	}
}
