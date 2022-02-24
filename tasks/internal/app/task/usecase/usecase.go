package usecase

import (
	"context"
	"github.com/santonov10/microservices/tasks/internal/app/models"
	"github.com/santonov10/microservices/tasks/internal/app/task"
)

var _ task.UseCase = (*TaskUseCase)(nil)

type TaskUseCase struct {
	repo task.Repository
}

func (u *TaskUseCase) CreateTask(ctx context.Context, taskData *task.InsertTask) (id string, err error) {
	if taskData.Header == "" {
		return "", task.ErrEmptyHeader
	}
	if taskData.UserID == "" {
		return "", task.ErrEmptyUserID
	}
	return u.repo.Create(ctx, taskData)
}

func (u *TaskUseCase) UpdateTask(ctx context.Context, id string, taskData *task.UpdateTask) error {
	if taskData.Header == "" {
		return task.ErrEmptyHeader
	}
	return u.repo.Update(ctx, id, taskData)
}

func (u *TaskUseCase) GetTasks(ctx context.Context, userID string) ([]*models.Task, error) {
	return u.repo.GetAllForUser(ctx, userID)
}

func NewTaskUseCase(repo task.Repository) *TaskUseCase {
	return &TaskUseCase{
		repo: repo,
	}
}
