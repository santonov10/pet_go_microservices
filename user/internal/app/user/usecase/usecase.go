package usecase

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/santonov10/microservices/user/internal/app/models"
	"github.com/santonov10/microservices/user/internal/app/user"
	"github.com/spf13/viper"
)

type UserUseCase struct {
	repo user.Repository

	hashSalt string
}

func (u *UserUseCase) CreateUser(ctx context.Context, uData *models.User) (id string, err error) {
	userC := *uData
	if userC.Password == "" {
		return "", user.ErrEmptyPassword
	}
	userC.Password = u.passwordHash(userC.Password)

	return u.repo.Create(ctx, &userC)
}

func (u *UserUseCase) GetUser(ctx context.Context, login, password string) (*models.User, error) {
	passHash := u.passwordHash(password)
	return u.repo.Get(ctx, login, passHash)
}

func (u *UserUseCase) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return u.repo.GetById(ctx, id)
}

func (u *UserUseCase) passwordHash(pass string) string {
	pwd := md5.New()
	pwd.Write([]byte(pass))
	pwd.Write([]byte(u.hashSalt))
	return fmt.Sprintf("%x", pwd.Sum(nil))
}

func NewUserUseCase(repo user.Repository) *UserUseCase {
	return &UserUseCase{
		repo:     repo,
		hashSalt: viper.GetString("password_hash_salt"),
	}
}
