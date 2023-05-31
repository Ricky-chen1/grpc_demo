package pack

import (
	"grpc_demo/app/user/internal/model"
	"grpc_demo/idl/pb/user"
)

func BuildUser(u model.User) *user.User {
	return &user.User{
		Id:       u.Id,
		Username: u.Username,
	}
}
