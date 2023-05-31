package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id             string
	Username       string `gorm:"unique"`
	PasswordDigest string
}

const cost = 12

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
