package pack

import (
	"grpc_demo/app/user/internal/model"
	"grpc_demo/idl/pb/user"
	"grpc_demo/pkg/errno"
)

func BuildUser(u model.User) *user.User {
	return &user.User{
		Id:       u.Id,
		Username: u.Username,
	}
}

func BuildBaseResp(err errno.Errno) *user.Base {
	return &user.Base{
		Code:    uint64(err.Code),
		Message: err.Msg,
	}
}
