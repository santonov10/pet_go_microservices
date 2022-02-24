package usecase

import (
	"context"
	"github.com/santonov10/microservices/user/internal/app/models"
	"github.com/santonov10/microservices/user/internal/app/user"
	"github.com/santonov10/microservices/user/internal/app/user/repository/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserUseCase(t *testing.T) {
	t.Run("normal data", func(t *testing.T) {

		repoMock := new(mock.UserRepoMock)
		UC := NewUserUseCase(repoMock)

		userData := models.NewUserSignIn("NewLogin", "Password")
		userData.ID = "GeneratedID"
		userDataRepoMock := *userData
		userDataRepoMock.Password = UC.passwordHash(userData.Password)
		repoMock.On("Get", userDataRepoMock.Login, userDataRepoMock.Password).Return(userData, nil)
		repoMock.On("GetById", userDataRepoMock.ID).Return(userData, nil)
		repoMock.On("Create", &userDataRepoMock).Return(userData.ID, nil)

		userID, err := UC.CreateUser(context.Background(), userData)
		require.NoError(t, err)
		require.NotEmpty(t, userID)
		userFromRepo, err := UC.GetUser(context.Background(), userData.Login, userData.Password)
		require.NoError(t, err)
		require.Equal(t, userFromRepo.Login, userData.Login)

		findId := "GeneratedID"
		user, err := UC.GetUserByID(context.Background(), findId)
		require.NoError(t, err)
		require.Equal(t, user.ID, findId)
	})

	t.Run("wrong data", func(t *testing.T) {
		repoMock := new(mock.UserRepoMock)
		UC := NewUserUseCase(repoMock)
		userData := models.NewUserSignIn("WrongData", "")
		userDataRepoMock := *userData
		userDataRepoMock.Password = UC.passwordHash(userData.Password)
		userID, err := UC.CreateUser(context.Background(), userData)
		require.ErrorIs(t, err, user.ErrEmptyPassword)
		require.Empty(t, userID)

		findId := "no user"
		repoMock.On("GetById", findId).Return(nil, user.ErrUserNotFound)
		userFound, err := UC.GetUserByID(context.Background(), findId)
		require.ErrorIs(t, err, user.ErrUserNotFound)
		require.Nil(t, userFound)
	})
}
