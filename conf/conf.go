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
)

func Init(service string) {

	viper.SetConfigName("conf")
	viper.SetConfigType("yml")
	viper.AddConfigPath("/homework/grpc_demo/conf")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Panicln("could not find config files")
		} else {
			log.Panicln("read config error")
		}
		log.Fatal(err)
	}

	configMapping(service) // 映射配置
}

func configMapping(srv string) {
	c := new(config)
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}

	Server = &c.Server
	Server.Secret = []byte(viper.GetString("server.jwt-secret"))

	Etcd = &c.Etcd

	Mysql = &c.Mysql

	Service = GetService(srv)
}

func GetService(srvname string) *service {

	addrlist := viper.GetStringSlice("services." + srvname + ".addr")

	return &service{
		Name: viper.GetString("services." + srvname + ".name"),
		Addr: addrlist[0],
	}
}
