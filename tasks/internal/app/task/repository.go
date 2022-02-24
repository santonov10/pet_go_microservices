package task

import (
	"context"
	"github.com/santonov10/microservices/tasks/internal/app/models"
)

type Repository interface {
	Create(ctx context.Context, task *InsertTask) (id string, err error)
	Update(ctx context.Context, id string, task *UpdateTask) (err error)
	GetAllForUser(ctx context.Context, userID string) ([]*models.Task, error)
}
