package usecase

import (
	"context"
	"github.com/santonov10/microservices/tasks/internal/app/models"
	"github.com/santonov10/microservices/tasks/internal/app/task"
	"github.com/stretchr/testify/mock"
)

var _ task.UseCase = (*TaskUseCaseMock)(nil)

type TaskUseCaseMock struct {
	mock.Mock
}

func (u *TaskUseCaseMock) CreateTask(ctx context.Context, taskData *task.InsertTask) (id string, err error) {
	args := u.Called(taskData)

	return args.String(0), args.Error(1)
}

func (u *TaskUseCaseMock) UpdateTask(ctx context.Context, id string, taskData *task.UpdateTask) error {
	args := u.Called(id, taskData)

	return args.Error(0)
}

func (u *TaskUseCaseMock) GetTasks(ctx context.Context, userID string) ([]*models.Task, error) {
	args := u.Called(userID)

	return args.Get(0).([]*models.Task), args.Error(1)
}
