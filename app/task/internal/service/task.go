package service

import (
	"context"
	"grpc_demo/app/task/internal/db"
	"grpc_demo/app/task/internal/pack"
	"grpc_demo/idl/pb/task"
	"grpc_demo/pkg/errno"
)

type TaskService struct {
	task.UnimplementedTaskServiceServer
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (ts *TaskService) TaskCreate(ctx context.Context, req *task.CreateReq) (*task.CreateRes, error) {
	res := new(task.CreateRes)
	newTask, err := db.CreateTask(req)

	if err != nil {
		res.Base = pack.BuildBaseResp(errno.Errno{
			Code: errno.TaskCreateFail,
			Msg:  errno.CodeTag[errno.TaskCreateFail],
		})
		return res, err
	}

	res.Base = pack.BuildBaseResp(errno.Errno{
		Code: errno.Success,
		Msg:  errno.CodeTag[errno.Success],
	})

	res.Task = pack.BuildTask(newTask)
	return res, nil
}

func (ts *TaskService) TaskUpdate(ctx context.Context, req *task.UpdateReq) (*task.UpdateRes, error) {
	res := new(task.UpdateRes)

	newTask, err := db.UpdateTask(req)
	if err != nil {
		res.Base = pack.BuildBaseResp(errno.Errno{
			Code: errno.TaskUpdateFail,
			Msg:  errno.CodeTag[errno.TaskUpdateFail],
		})
		return res, nil
	}

	res.Base = pack.BuildBaseResp(errno.Errno{
		Code: errno.Success,
		Msg:  errno.CodeTag[errno.Success],
	})
	res.Task = pack.BuildTask(newTask)
	return res, nil
}

func (ts *TaskService) TaskListGet(ctx context.Context, req *task.GetListReq) (*task.GetListRes, error) {
	res := new(task.GetListRes)

	newTasks, err := db.GetTaskList(req)
	if err != nil {
		res.Base = pack.BuildBaseResp(errno.Errno{
			Code: errno.GetTaskListFail,
			Msg:  errno.CodeTag[errno.GetTaskListFail],
		})
		return res, nil
	}

	res.Base = pack.BuildBaseResp(errno.Errno{
		Code: errno.Success,
		Msg:  errno.CodeTag[errno.Success],
	})
	res.Task = pack.BuildTaskList(newTasks)

	return res, nil
}

func (ts *TaskService) TaskSearch(ctx context.Context, req *task.SearchReq) (*task.SearchRes, error) {
	res := new(task.SearchRes)

	newTasks, err := db.SearchTask(req)
	if err != nil {
		res.Base = pack.BuildBaseResp(errno.Errno{
			Code: errno.TaskSearchFail,
			Msg:  errno.CodeTag[errno.TaskSearchFail],
		})
	}

	res.Base = pack.BuildBaseResp(errno.Errno{
		Code: errno.Success,
		Msg:  errno.CodeTag[errno.Success],
	})
	res.Task = pack.BuildTaskList(newTasks)

	return res, nil
}

func (ts *TaskService) TaskDelete(ctx context.Context, req *task.DeleteReq) (*task.DeleteRes, error) {
	res := new(task.DeleteRes)

	if err := db.DeleteTask(req); err != nil {
		res.Base = pack.BuildBaseResp(errno.Errno{
			Code: errno.TaskDeleteFail,
			Msg:  errno.CodeTag[errno.TaskDeleteFail],
		})
		return res, err
	}

	res.Base = pack.BuildBaseResp(errno.Errno{
		Code: errno.Success,
		Msg:  errno.CodeTag[errno.Success],
	})

	return res, nil
}
