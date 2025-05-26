package scheduler

import (
	"context"
	"github.com/Demonyker/personal-assistant-scheduler/internal/entity"
	"github.com/Demonyker/personal-assistant-scheduler/internal/repo"
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

func (u *UseCase) Save(ctx context.Context, task entity.Task) error {
	return u.repo.Save(ctx, task)
}

func (u *UseCase) GetMany(ctx context.Context, userId string) ([]entity.Task, error) {
	return u.repo.GetMany(ctx, userId)
}
