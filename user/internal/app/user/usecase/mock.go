package usecase

import (
	"context"
	"github.com/santonov10/microservices/user/internal/app/models"
	"github.com/stretchr/testify/mock"
)

type UserUseCaseMock struct {
	mock.Mock
}

func (u *UserUseCaseMock) CreateUser(ctx context.Context, uData *models.User) (id string, err error) {
	args := u.Called(uData)

	return args.String(0), args.Error(1)
}

func (u *UserUseCaseMock) GetUser(ctx context.Context, login, password string) (*models.User, error) {
	args := u.Called(login, password)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*models.User), args.Error(1)
}

func (u *UserUseCaseMock) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	args := u.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*models.User), args.Error(1)
}
