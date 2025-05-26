package usecase

import (
	"context"
	"github.com/Demonyker/personal-assistant-scheduler/internal/entity"
)

type (
	// SchedulerUseCase -.
	SchedulerUseCase interface {
		Save(ctx context.Context, task entity.Task) error
		GetMany(ctx context.Context, userId string) ([]entity.Task, error)
	}
)
