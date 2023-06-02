package conf

type server struct {
	Secret  []byte
	Version string
	Name    string
}

type service struct {
	Name string
	Addr string
}

type mysql struct {
	Addr     string
	Database string
	Username string
	Password string
	Charset  string
}

type etcd struct {
	Addr string
}

type config struct {
	Etcd    etcd
	Server  server
	Mysql   mysql
	Service service
}
