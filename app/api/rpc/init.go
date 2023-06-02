package rpc

import (
	"context"
	"fmt"
	"grpc_demo/conf"
	"grpc_demo/idl/pb/task"
	"grpc_demo/idl/pb/user"
	"grpc_demo/pkg/discovery"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

var (
	Register   *discovery.Resolver
	ctx        context.Context
	CancelFunc context.CancelFunc

	userClient user.UserServiceClient
	taskClient task.TaskServiceClient
)

const (
	userSrvName = "user"
	taskSrvName = "task"
)

func Init() {
	Register = discovery.NewResolver([]string{conf.Etcd.Addr}, logrus.New())
	resolver.Register(Register)
	ctx, CancelFunc = context.WithTimeout(context.Background(), 3*time.Second)

	defer Register.Close()

	initClient(conf.GetService(userSrvName).Name, &userClient)
	initClient(conf.GetService(taskSrvName).Name, &taskClient)
}

func initClient(serviceName string, client interface{}) {
	conn, err := connectServer(serviceName)

	if err != nil {
		panic(err)
	}

	switch c := client.(type) {
	case *user.UserServiceClient:
		*c = user.NewUserServiceClient(conn)
	case *task.TaskServiceClient:
		*c = task.NewTaskServiceClient(conn)
	default:
		panic("unsupported client type")
	}
}

func connectServer(serviceName string) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	addr := fmt.Sprintf("%s:///%s", Register.Scheme(), serviceName)

	return grpc.DialContext(ctx, addr, opts...)
}
