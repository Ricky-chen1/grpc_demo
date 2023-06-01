package db

import (
	"grpc_demo/pkg/util"

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
		panic(err)
	}

	DB = db
	Migration()
}
