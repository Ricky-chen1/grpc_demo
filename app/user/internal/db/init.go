package db

import (
	"grpc_demo/pkg/util"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Init() {

	db, err := gorm.Open(mysql.Open(util.GetMysqlDSN()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Println("database err")
	}

	DB = db
	Migration()
}
