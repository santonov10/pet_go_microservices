package user

import (
	"context"
	"github.com/santonov10/microservices/user/internal/app/models"
)

type UseCase interface {
	CreateUser(ctx context.Context, user *models.User) (id string, err error)
	GetUser(ctx context.Context, login, password string) (*models.User, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
}
