package services

import (
	"context"
	"github.com/santonov10/microservices/user/api/grpc/pb"
	"github.com/santonov10/microservices/user/internal/app/models"
	"github.com/santonov10/microservices/user/internal/app/user"
	"github.com/santonov10/microservices/user/internal/app/user/usecase"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func newTestService(useCase user.UseCase, durationSeconds int64) *UserService {
	return &UserService{
		userUC:         useCase,
		signingKey:     []byte("testSigningKey"),
		expireDuration: time.Duration(durationSeconds * int64(time.Second)),
	}
}

func TestUserService_GetUser(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		ctx := context.Background()

		UCMock := new(usecase.UserUseCaseMock)
		var expireSeconds int64 = 1
		service := newTestService(UCMock, expireSeconds)
		pbUser, err := service.GetUser(ctx, &pb.Token{Token: "no token"})
		require.Error(t, err)
		require.Empty(t, pbUser)

		mockUserData := models.NewUserRegistration("test", "pass")
		mockUserDataWithId := *mockUserData
		mockUserDataWithId.ID = "testID"
		regRequest := pb.RegistrationRequest{
			Login:    mockUserData.Login,
			Password: mockUserData.Password,
		}
		UCMock.On("CreateUser", mockUserData).Return(mockUserDataWithId.ID, nil)
		newToken, err := service.Registration(ctx, &regRequest)
		require.NoError(t, err)
		require.NotEmpty(t, newToken)

		UCMock.On("GetUserByID", mockUserDataWithId.ID).Return(&mockUserDataWithId, nil)
		pbUserResponse, err := service.GetUser(ctx, newToken)
		require.NoError(t, err)
		require.Equal(t, pbUserResponse.Id, mockUserDataWithId.ID)

		//просроченый токен
		time.Sleep(time.Duration(expireSeconds+2) * time.Second)
		pbUserResponse, err = service.GetUser(ctx, newToken)
		require.Error(t, err)
		require.Empty(t, pbUserResponse.Id)

		UCMock.On("GetUser", mockUserData.Login, mockUserData.Password).Return(&mockUserDataWithId, nil)
		newToken, err = service.Login(ctx, &pb.LoginRequest{
			Login:    mockUserData.Login,
			Password: mockUserData.Password,
		})
		require.NoError(t, err)
		require.NotEmpty(t, newToken)
	})
}
