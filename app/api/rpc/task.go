package rpc

import (
	"context"
	"errors"
	"grpc_demo/idl/pb/task"
	"log"
)

func TaskCreate(ctx context.Context, req *task.CreateReq) (*task.Task, error) {
	res, err := taskClient.TaskCreate(ctx, req)
	if err != nil {
		log.Fatalf("task create failed %v", err)
		return nil, errors.New("task rpc call failed")
	}

	return res.Task, nil
}

func TaskListGet(ctx context.Context, req *task.GetListReq) ([]*task.Task, error) {
	res, err := taskClient.TaskListGet(ctx, req)
	if err != nil {
		return nil, errors.New("task rpc call faild")
	}

	return res.Task, nil
}

func TaskSearch(ctx context.Context, req *task.SearchReq) ([]*task.Task, error) {
	res, err := taskClient.TaskSearch(ctx, req)
	if err != nil {
		return nil, errors.New("task rpc call faild")
	}

	return res.Task, nil
}

func TaskUpdate(ctx context.Context, req *task.UpdateReq) (*task.Task, error) {
	res, err := taskClient.TaskUpdate(ctx, req)
	if err != nil {
		return nil, errors.New("task rpc call failed")
	}

	return res.Task, nil
}

func TaskDelete(ctx context.Context, req *task.DeleteReq) error {
	_, err := taskClient.TaskDelete(ctx, req)
	if err != nil {
		return errors.New("task rpc call failed")
	}

	return nil
}
