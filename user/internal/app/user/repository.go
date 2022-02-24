package user

import (
	"context"
	"github.com/santonov10/microservices/user/internal/app/models"
)

type Repository interface {
	Create(ctx context.Context, user *models.User) (id string, err error)
	Get(ctx context.Context, login, password string) (*models.User, error)
	GetById(ctx context.Context, id string) (*models.User, error)
	Delete(ctx context.Context, id string) error
}
