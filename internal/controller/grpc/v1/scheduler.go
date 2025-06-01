package v1

import (
	"context"
	v1 "github.com/Demonyker/personal-assistant-contracts/contracts/scheduler/v1"
)

func (r *V1) Save(ctx context.Context, request *v1.AddTasksRequest) (*v1.AddTasksResponse, error) {
	return r.suc.Save(ctx, request)
}

func (r *V1) GetMany(ctx context.Context, request *v1.GetUserTasksRequest) (*v1.GetUserTasksResponse, error) {
	return r.suc.GetMany(ctx, request)
}
