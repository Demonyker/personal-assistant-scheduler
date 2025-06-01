package repo

import (
	"context"
	"github.com/Demonyker/personal-assistant-scheduler/internal/entity"
)

type (
	// TasksRepo - repository to work with user tasks.
	TasksRepo interface {
		Save(ctx context.Context, tasks []entity.Task) error
		GetMany(ctx context.Context, userID string) ([]entity.Task, error)
	}
)
