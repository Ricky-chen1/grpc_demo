package handler

import (
	"grpc_demo/app/api/rpc"
	"grpc_demo/idl/pb/user"
	"grpc_demo/pkg/errno"
	"grpc_demo/pkg/util"

	"github.com/gin-gonic/gin"
)

// 网关进入后客户端调用rpc方法
func UserRegister(c *gin.Context) {
	req := new(user.RegisterReq)

	if err := c.ShouldBind(req); err != nil {
		BuildFailResponse(c, errno.ParamsInvalid, errno.CodeTag[errno.ParamsInvalid])
		return
	}

	// 通过封装的rpc包调用pb中rpc方法
	err := rpc.UserRegister(c, req)
	if err != nil {
		BuildFailResponse(c, errno.CallRPCFailed, errno.CodeTag[errno.CallRPCFailed])
		return
	}

	BuildSuccessResponse(c, errno.Success, errno.CodeTag[errno.Success], nil)
}

func UserLogin(c *gin.Context) {
	req := new(user.LoginReq)

	if err := c.ShouldBind(req); err != nil {
		BuildFailResponse(c, errno.ParamsInvalid, errno.CodeTag[errno.ParamsInvalid])
		return
	}

	userData, err := rpc.UserLogin(c, req)
	if err != nil {
		BuildFailResponse(c, errno.CallRPCFailed, errno.CodeTag[errno.CallRPCFailed])
		return
	}

	token, err := util.SignToken(userData.Id)
	if err != nil {
		BuildFailResponse(c, errno.TokenGenerateFail, errno.CodeTag[errno.TokenGenerateFail])
		return
	}

	BuildSuccessResponse(c, errno.Success, errno.CodeTag[errno.Success], gin.H{
		"user":  userData,
		"token": token,
	})
}
