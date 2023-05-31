package db

import (
	"grpc_demo/pkg/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(mysql.Open(util.GetMysqlDSN()), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Migration()
	DB = db
}
