package util

import (
	"fmt"
	"grpc_demo/conf"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID                 string `json:"id"`
	jwt.StandardClaims        // 实现了valid方法
}

var expireTime = time.Hour * 24

func SignToken(id string) (string, error) {
	claims := Claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireTime).Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString(conf.Server.Secret)
	fmt.Println(conf.Server.Secret, "-------------------------------")

	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string) (*Claims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return conf.Server.Secret, nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := jwtToken.Claims.(*Claims)
	if !ok || !jwtToken.Valid {
		return nil, err
	}

	return claims, nil
}
