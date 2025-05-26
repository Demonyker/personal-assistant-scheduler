package repo

import (
	"context"
	"github.com/Demonyker/personal-assistant-scheduler/internal/entity"
)

type (
	// TasksRepo -.
	TasksRepo interface {
		Save(ctx context.Context, task entity.Task) error
		GetMany(ctx context.Context, userId string) ([]entity.Task, error)
	}
)
