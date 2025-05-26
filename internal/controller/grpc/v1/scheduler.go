package v1

import (
	"context"
	v1 "github.com/Demonyker/personal-assistant-contracts/contracts/scheduler/v1"
	"github.com/Demonyker/personal-assistant-scheduler/internal/entity"
	"github.com/google/uuid"
	"time"
)

func (r *V1) Save(ctx context.Context, request *v1.AddTasksRequest) (*v1.AddTasksResponse, error) {
	id := uuid.New()
	date, err := time.Parse("2006-01-02", request.Date)
	if err != nil {
		return nil, err
	}
	for _, task := range request.Tasks {
		err = r.suc.Save(ctx, entity.Task{
			ID:          id.String(),
			Description: task,
			Date:        date,
			UserID:      request.UserId,
		})

		if err != nil {
			return nil, err
		}
	}

	tasks, err := r.suc.GetMany(ctx, request.UserId)

	if err != nil {
		return nil, err
	}

	tasksResp := make([]*v1.Task, 0, len(tasks))

	for _, task := range tasks {
		tasksResp = append(tasksResp, &v1.Task{Id: task.ID, Description: task.Description, Date: task.Date.String()})
	}
	return &v1.AddTasksResponse{Tasks: tasksResp}, nil
}

func (r *V1) GetMany(ctx context.Context, request *v1.GetUserTasksRequest) (*v1.GetUserTasksResponse, error) {
	tasks, err := r.suc.GetMany(ctx, request.UserId)

	if err != nil {
		return nil, err
	}

	tasksResp := make([]*v1.Task, 0, len(tasks))

	for _, task := range tasks {
		tasksResp = append(tasksResp, &v1.Task{Id: task.ID, Description: task.Description, Date: task.Date.String()})
	}

	return &v1.GetUserTasksResponse{Tasks: tasksResp}, nil
}
