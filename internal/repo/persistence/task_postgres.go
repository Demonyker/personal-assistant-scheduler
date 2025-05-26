package persistent

import (
	"context"
	"fmt"
	"github.com/Demonyker/personal-assistant-scheduler/internal/entity"
	"github.com/Demonyker/personal-assistant-scheduler/pkg/postgres"
)

// TasksRepo -.
type TasksRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *TasksRepo {
	return &TasksRepo{pg}
}

func (tr *TasksRepo) Save(ctx context.Context, task entity.Task) error {
	sql, args, err := tr.Builder.
		Insert("scheduler.tasks").
		Columns("id", "description", "date", "user_id").
		Values(task.ID, task.Description, task.Date, task.UserID).
		ToSql()

	if err != nil {
		return fmt.Errorf("TasksRepo - Save - ur.Builder: %w", err)
	}

	_, err = ur.Pool.Exec(ctx, sql, args...)

	if err != nil {
		return fmt.Errorf("TasksRepo - Save - ur.Pool.Exec: %w", err)
	}

	return nil
}

func (tr *TasksRepo) GetMany(ctx context.Context, userId string) ([]entity.Task, error) {
	sql, args, err := tr.Builder.
		Select("id", "description", "date", "user_id").
		From("scheduler.tasks").
		Where("user_id = ?", userId).
		ToSql()

	tasks := make([]entity.Task, 0, 15)

	if err != nil {
		return tasks, fmt.Errorf("TaskRepo - GetMany - ur.Builder: %w", err)
	}

	rows, err := tr.Pool.Query(ctx, sql, args...)

	if err != nil {
		return tasks, fmt.Errorf("TaskRepo - GetMany - ur.Pool.Exec: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		entityTask := entity.Task{}
		err = rows.Scan(&entityTask.ID, &entityTask.Description, &entityTask.Date, &entityTask.UserID)

		if err != nil {
			return tasks, fmt.Errorf("TaskRepo - GetMany - rows.Scan: %w", err)
		}

		tasks = append(tasks, entityTask)
	}

	return tasks, nil
}
