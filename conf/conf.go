package conf

import (
	"log"

	"github.com/spf13/viper"
)

var (
	Server  *server
	Mysql   *mysql
	Service *service
	Etcd    *etcd
	C       *config
)

func Init() {
	viper.SetConfigName("conf")
	viper.AddConfigPath("/homework/grpc_demo/conf")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("config not found %v", err)
		} else {
			log.Fatalf("other error %v", err)
		}
	}

	if err := viper.Unmarshal(&C); err != nil {
		log.Fatalf("mapping config error %v", err)
		panic(err)
	}

	Server = &C.Server
	Server.Secret = []byte(viper.GetString("server.jwt-secret"))

	Mysql = &C.Mysql

}
