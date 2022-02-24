package mock

import (
	"context"
	"github.com/santonov10/microservices/user/internal/app/models"
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (u UserRepoMock) Delete(ctx context.Context, id string) error {
	args := u.Called(id)

	return args.Error(0)
}

func (u UserRepoMock) Create(ctx context.Context, user *models.User) (id string, err error) {
	args := u.Called(user)

	return args.String(0), args.Error(1)
}

func (u UserRepoMock) Get(ctx context.Context, login, password string) (*models.User, error) {
	args := u.Called(login, password)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*models.User), args.Error(1)
}

func (u UserRepoMock) GetById(ctx context.Context, id string) (*models.User, error) {
	args := u.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*models.User), args.Error(1)
}
