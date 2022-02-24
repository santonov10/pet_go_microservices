package handlers

import "errors"

var ErrWrongDataFormat = errors.New("ошибка формата данных")
var ErrServiceIsDown = errors.New("сервис не отвечает")
