package scheduler

import (
	"context"
	v1 "github.com/Demonyker/personal-assistant-contracts/contracts/scheduler/v1"
	"github.com/Demonyker/personal-assistant-scheduler/internal/entity"
	"github.com/Demonyker/personal-assistant-scheduler/internal/repo"
	"github.com/google/uuid"
	"time"
)

// UseCase -.
type UseCase struct {
	repo repo.TasksRepo
}

// New -.
func New(r repo.TasksRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

func (u *UseCase) Save(ctx context.Context, request *v1.AddTasksRequest) (*v1.AddTasksResponse, error) {
	id := uuid.New()
	date, err := time.Parse("2006-01-02", request.Date)
	if err != nil {
		return nil, err
	}
	taskEntities := make([]entity.Task, 0, len(request.Tasks))
	for _, task := range request.Tasks {
		taskEntity := entity.Task{
			ID:          id.String(),
			Description: task,
			Date:        date,
			UserID:      request.UserId,
		}

		taskEntities = append(taskEntities, taskEntity)
	}

	err = u.repo.Save(ctx, taskEntities)
	if err != nil {
		return nil, err
	}

	return &v1.AddTasksResponse{IsSuccess: true}, nil
}

func (u *UseCase) GetMany(ctx context.Context, request *v1.GetUserTasksRequest) (*v1.GetUserTasksResponse, error) {
	tasks, err := u.repo.GetMany(ctx, request.UserId)

	if err != nil {
		return nil, err
	}

	tasksResp := make([]*v1.Task, 0, len(tasks))

	for _, task := range tasks {
		tasksResp = append(tasksResp, &v1.Task{Id: task.ID, Description: task.Description, Date: task.Date.String()})
	}

	return &v1.GetUserTasksResponse{Tasks: tasksResp}, nil
}
