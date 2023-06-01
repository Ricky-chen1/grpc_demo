package service

import (
	"context"
	"grpc_demo/app/user/internal/db"
	"grpc_demo/app/user/internal/pack"
	"grpc_demo/idl/pb/user"
	"grpc_demo/pkg/errno"
)

type UserService struct {
	user.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) UserRegister(ctx context.Context, req *user.RegisterReq) (*user.RegisterRes, error) {
	res := new(user.RegisterRes)

	//数据库交互
	if err := db.CreateUser(req); err != nil {
		res.Base.Code = errno.UserCreateFail
		res.Base.Message = errno.CodeTag[errno.UserCreateFail]
		return res, err
	}

	res.Base.Code = errno.Success
	res.Base.Message = errno.CodeTag[errno.Success]
	return res, nil

}

func (us *UserService) UserLogin(ctx context.Context, req *user.LoginReq) (*user.LoginRes, error) {
	res := new(user.LoginRes)
	//与数据库交互
	newUser, err := db.GetUserByName(req.Username)
	if err != nil {
		res.Base.Code = errno.UserLoginFail
		res.Base.Message = errno.CodeTag[errno.UserLoginFail]
		return res, err
	}

	if err := newUser.CheckPassword(req.Password); err != nil {
		res.Base.Code = errno.CheckPasswordFail
		res.Base.Message = errno.CodeTag[errno.CheckPasswordFail]
		return res, err
	}

	res.Base.Code = errno.Success
	res.Base.Message = errno.CodeTag[errno.Success]
	res.User = pack.BuildUser(*newUser)

	return res, nil
}
