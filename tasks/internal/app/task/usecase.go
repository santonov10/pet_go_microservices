package task

import (
	"context"
	"github.com/santonov10/microservices/tasks/internal/app/models"
)

type InsertTask struct {
	Header      string
	Description string
	UserID      string
}

type UpdateTask struct {
	Header      string
	Description string
}

type UseCase interface {
	CreateTask(ctx context.Context, task *InsertTask) (id string, err error)
	UpdateTask(ctx context.Context, id string, task *UpdateTask) error
	GetTasks(ctx context.Context, userID string) ([]*models.Task, error)
}
