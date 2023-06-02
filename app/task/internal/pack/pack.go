package pack

import (
	"grpc_demo/app/task/internal/model"
	"grpc_demo/idl/pb/task"
	"grpc_demo/pkg/errno"
)

func BuildTask(t *model.Task) *task.Task {
	return &task.Task{
		Id:      t.Id,
		Content: t.Content,
		Title:   t.Title,
		Status:  int64(t.Status),
	}
}

func BuildTaskList(items []model.Task) []*task.Task {
	var list []*task.Task
	for _, item := range items {
		task := BuildTask(&item)
		list = append(list, task)
	}
	return list
}

func BuildBaseResp(err errno.Errno) *task.Base {
	return &task.Base{
		Code:    uint64(err.Code),
		Message: err.Msg,
	}
}
