package rpc

import (
	"grpc_demo/idl/pb/task"
	"grpc_demo/idl/pb/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	userClient user.UserServiceClient
	taskClient task.TaskServiceClient
)

func Init() {
	conn := ConnetToServer()
	SaveStub(conn)
}

// 创建存根文件,以便进行rpc方法的调用
func SaveStub(conn *grpc.ClientConn) {
	userClient = user.NewUserServiceClient(conn)
	taskClient = task.NewTaskServiceClient(conn)
}

// grpc客户端与grpc服务端建立通信
func ConnetToServer() *grpc.ClientConn {

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial("localhost:8080", opts...)

	if err != nil {
		panic(err)
	}

	defer conn.Close()
	return conn
}
