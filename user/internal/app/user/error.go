package user

import "errors"

var ErrEmptyPassword = errors.New("пароль не может быть пустым")

var ErrUserNotFound = errors.New("пользователь не найден")
