package services

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrUserNotFound = status.Error(codes.NotFound, "пользователь не найден")
