package task

import "errors"

var ErrEmptyHeader = errors.New("заголовок не может быть пустым")
var ErrEmptyUserID = errors.New("userID не может быть пустым")
