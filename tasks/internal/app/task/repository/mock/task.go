package mock

import (
	"context"
	"github.com/santonov10/microservices/tasks/internal/app/models"
	"github.com/santonov10/microservices/tasks/internal/app/task"
	"github.com/stretchr/testify/mock"
)

var _ task.Repository = (*TaskRepoMock)(nil)

type TaskRepoMock struct {
	mock.Mock
}

func (u TaskRepoMock) Create(ctx context.Context, task *task.InsertTask) (id string, err error) {
	args := u.Called(task)

	return args.String(0), args.Error(1)
}

func (u TaskRepoMock) Update(ctx context.Context, id string, task *task.UpdateTask) (err error) {
	args := u.Called(id, task)

	return args.Error(0)
}

func (u TaskRepoMock) GetAllForUser(ctx context.Context, userID string) ([]*models.Task, error) {
	args := u.Called(userID)

	return args.Get(0).([]*models.Task), args.Error(1)
}
