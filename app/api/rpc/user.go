package rpc

import (
	"context"
	"errors"
	"grpc_demo/idl/pb/user"
)

// user moudle grpc client call
func UserRegister(ctx context.Context, req *user.RegisterReq) error {
	_, err := userClient.UserRegister(ctx, req)

	if err != nil {
		return errors.New("user rpc call failed")
	}
	return nil
}

func UserLogin(ctx context.Context, req *user.LoginReq) (*user.User, error) {
	res, err := userClient.UserLogin(ctx, req)
	if err != nil {
		return nil, errors.New("user rpc call failed")
	}

	return res.User, nil
}
