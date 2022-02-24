package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/santonov10/microservices/user/api/grpc/pb"
	"github.com/santonov10/microservices/user/internal/app/models"
	"github.com/santonov10/microservices/user/internal/app/user"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"time"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID string
}

type UserService struct {
	pb.UnimplementedUserServiceServer

	userUC user.UseCase

	signingKey     []byte
	expireDuration time.Duration
}

func (u *UserService) Registration(ctx context.Context, request *pb.RegistrationRequest) (*pb.Token, error) {
	userData := models.NewUserRegistration(request.Login, request.Password)
	userID, err := u.userUC.CreateUser(ctx, userData)
	if err != nil {
		return &pb.Token{}, err
	}
	token, err := u.createToken(userID)
	if err != nil {
		return &pb.Token{}, err
	}

	return &pb.Token{Token: token}, err
}

func (u *UserService) Login(ctx context.Context, request *pb.LoginRequest) (*pb.Token, error) {
	us, err := u.userUC.GetUser(ctx, request.Login, request.Password)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			return &pb.Token{}, ErrUserNotFound
		}
		return &pb.Token{}, err
	}

	token, err := u.createToken(us.ID)
	if err != nil {
		return &pb.Token{}, err
	}

	return &pb.Token{Token: token}, err
}

func (u *UserService) GetUser(ctx context.Context, token *pb.Token) (*pb.UserResponse, error) {
	userID, err := u.parseToken(token.Token)
	if err != nil {
		return &pb.UserResponse{}, err
	}

	user, err := u.userUC.GetUserByID(ctx, userID)
	if err != nil {
		return &pb.UserResponse{}, err
	}

	return &pb.UserResponse{
		Id:    user.ID,
		Login: user.Login,
	}, nil
}

func NewUserService(userUC user.UseCase) *UserService {
	tokenTtlSeconds, err := strconv.ParseInt(viper.GetString("token_ttl_seconds"), 10, 64)
	if err != nil {
		log.Fatal("ошибка преобразования в число token_ttl_seconds")
	}
	return &UserService{
		userUC:         userUC,
		signingKey:     []byte(viper.GetString("signing_key")),
		expireDuration: time.Duration(tokenTtlSeconds * int64(time.Second)),
	}
}

func (u *UserService) createToken(userID string) (token string, err error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(u.expireDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: userID,
	})

	return jwtToken.SignedString(u.signingKey)
}

func (u *UserService) parseToken(accessToken string) (userID string, err error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return u.signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		return claims.UserID, nil
	}

	return "", fmt.Errorf("неверный токен")
}
