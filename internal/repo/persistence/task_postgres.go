package persistent

import (
	"context"
	"fmt"
	"github.com/Demonyker/personal-assistant-scheduler/internal/entity"
	"github.com/Demonyker/personal-assistant-scheduler/pkg/postgres"
)

// TasksRepo - repository to work with user tasks.
type TasksRepo struct {
	*postgres.Postgres
}

// New - func to create tasks repository.
func New(pg *postgres.Postgres) *TasksRepo {
	return &TasksRepo{pg}
}

func (tr *TasksRepo) Save(ctx context.Context, tasks []entity.Task) error {
	queryBuilder := tr.Builder.
		Insert("scheduler.tasks").
		Columns("id", "description", "date", "user_id")

	for _, v := range tasks {
		queryBuilder.Values(v.ID, v.Description, v.Date, v.UserID)
	}

	sql, args, err := queryBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("TasksRepo - Save - ur.Builder: %w", err)
	}

	_, err = tr.Pool.Exec(ctx, sql, args...)

	if err != nil {
		return fmt.Errorf("TasksRepo - Save - ur.Pool.Exec: %w", err)
	}

	return nil
}

func (tr *TasksRepo) GetMany(ctx context.Context, userID string) ([]entity.Task, error) {
	sql, args, err := tr.Builder.
		Select("id", "description", "date", "user_id").
		From("scheduler.tasks").
		Where("user_id = ?", userID).
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
