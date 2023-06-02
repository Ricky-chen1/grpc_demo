package handler

import (
	"grpc_demo/app/api/rpc"
	"grpc_demo/idl/pb/task"
	"grpc_demo/pkg/errno"

	"github.com/gin-gonic/gin"
)

func TaskCreate(c *gin.Context) {
	req := new(task.CreateReq)

	if err := c.ShouldBind(req); err != nil {
		BuildFailResponse(c, errno.ParamsInvalid, errno.CodeTag[errno.ParamsInvalid])
		return
	}

	// 通过封装的rpc包调用pb中rpc方法
	res, err := rpc.TaskCreate(c, req)
	if err != nil {
		BuildFailResponse(c, errno.CallRPCFailed, errno.CodeTag[errno.CallRPCFailed])
		return
	}

	BuildSuccessResponse(c, errno.Success, errno.CodeTag[errno.Success], res)
}

func TaskUpdate(c *gin.Context) {
	req := new(task.UpdateReq)

	if err := c.ShouldBind(req); err != nil {
		BuildFailResponse(c, errno.ParamsInvalid, errno.CodeTag[errno.ParamsInvalid])
		return
	}

	// 通过封装的rpc包调用pb中rpc方法
	res, err := rpc.TaskUpdate(c, req)
	if err != nil {
		BuildFailResponse(c, errno.CallRPCFailed, errno.CodeTag[errno.CallRPCFailed])
		return
	}

	BuildSuccessResponse(c, errno.Success, errno.CodeTag[errno.Success], res)

}

func TaskListGet(c *gin.Context) {
	req := new(task.GetListReq)

	if err := c.ShouldBind(req); err != nil {
		BuildFailResponse(c, errno.ParamsInvalid, errno.CodeTag[errno.ParamsInvalid])
		return
	}

	// 通过封装的rpc包调用pb中rpc方法
	res, err := rpc.TaskListGet(c, req)
	if err != nil {
		BuildFailResponse(c, errno.CallRPCFailed, errno.CodeTag[errno.CallRPCFailed])
		return
	}

	BuildSuccessResponse(c, errno.Success, errno.CodeTag[errno.Success], res)
}

func TaskSearch(c *gin.Context) {
	req := new(task.SearchReq)

	if err := c.ShouldBind(req); err != nil {
		BuildFailResponse(c, errno.ParamsInvalid, errno.CodeTag[errno.ParamsInvalid])
		return
	}

	// 通过封装的rpc包调用pb中rpc方法
	res, err := rpc.TaskSearch(c, req)
	if err != nil {
		BuildFailResponse(c, errno.CallRPCFailed, errno.CodeTag[errno.CallRPCFailed])
		return
	}

	BuildSuccessResponse(c, errno.Success, errno.CodeTag[errno.Success], res)
}

func TaskDelete(c *gin.Context) {
	req := new(task.DeleteReq)

	if err := c.ShouldBind(req); err != nil {
		BuildFailResponse(c, errno.ParamsInvalid, errno.CodeTag[errno.ParamsInvalid])
		return
	}

	// 通过封装的rpc包调用pb中rpc方法
	err := rpc.TaskDelete(c, req)
	if err != nil {
		BuildFailResponse(c, errno.CallRPCFailed, errno.CodeTag[errno.CallRPCFailed])
		return
	}

	BuildSuccessResponse(c, errno.Success, errno.CodeTag[errno.Success], nil)
}
