package db

import (
	"grpc_demo/app/user/internal/model"
	"grpc_demo/idl/pb/user"
)

func CreateUser(req *user.RegisterReq) error {
	user := model.User{
		Username: req.Username,
	}

	//加密在哪一层进行?
	if err := user.SetPassword(req.Password); err != nil {
		return err
	}

	if err := DB.Model(&model.User{}).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByName(name string) (*model.User, error) {
	var user model.User
	if err := DB.Model(&model.User{}).Where("username = ?", name).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
