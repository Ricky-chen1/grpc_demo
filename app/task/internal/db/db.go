package db

import (
	"grpc_demo/app/task/internal/model"
	"grpc_demo/idl/pb/task"
	"grpc_demo/pkg/util"
)

func CreateTask(req *task.CreateReq) (*model.Task, error) {
	task := &model.Task{
		User_id: req.UserId,
		Id:      util.NewUuid(),
		Title:   req.Title,
		Content: req.Content,
		Status:  0,
	}

	err := DB.Model(&model.Task{}).Create(task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func GetTaskList(req *task.GetListReq) ([]model.Task, error) {
	var tasks []model.Task
	if err := DB.Model(&model.Task{}).Where("status = ?", req.Status).
		Find(&tasks).Limit(int(req.PageSize)).Offset((int(req.PageNum) - 1) * int(req.PageSize)).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func SearchTask(req *task.SearchReq) ([]model.Task, error) {
	var tasks []model.Task
	if err := DB.Model(&model.Task{}).Where("content LIKE ? OR title LIKE ?", req.Key, req.Key).Find(&tasks).
		Limit(int(req.PageSize)).Offset((int(req.PageNum) - 1) * int(req.PageSize)).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// 更新状态
func UpdateTask(req *task.UpdateReq) (*model.Task, error) {
	var task model.Task
	if err := DB.Model(&model.Task{}).Where("id = ?", req.Id).Find(&task).Error; err != nil {
		return nil, err
	}
	task.Status = int(req.Status)
	if err := DB.Model(&model.Task{}).Save(task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func DeleteTask(req *task.DeleteReq) error {
	var task model.Task
	if err := DB.Model(&model.Task{}).Where("id = ?", req.Id).Delete(task).Error; err != nil {
		return err
	}
	return nil
}
