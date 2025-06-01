package usecase

import (
	"context"
	v1 "github.com/Demonyker/personal-assistant-contracts/contracts/scheduler/v1"
)

type (
	// SchedulerUseCase - use-case to work with user tasks.
	SchedulerUseCase interface {
		Save(ctx context.Context, request *v1.AddTasksRequest) (*v1.AddTasksResponse, error)
		GetMany(ctx context.Context, request *v1.GetUserTasksRequest) (*v1.GetUserTasksResponse, error)
	}
)
