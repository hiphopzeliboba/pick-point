package model

import "errors"

// ErrNotFound возвращается, когда запрашиваемая запись не найдена в базе данных
var ErrNotFound = errors.New("requested record not found")

// ErrActiveIntakeExists возвращается при попытке создать новую приёмку, когда уже есть активная
var ErrActiveIntakeExists = errors.New("active intake already exists for this pick point")

// ErrIntakeAlreadyClosed возвращается при попытке закрыть уже закрытую приёмку
var ErrIntakeAlreadyClosed = errors.New("intake is already closed")

// ErrInvalidCity возвращается при попытке создать ПВЗ в неподдерживаемом городе
var ErrInvalidCity = errors.New("city is not supported")

// ErrInvalidRole возвращается при несответсвии роли пользователя
var ErrInvalidRole = errors.New("invalid user role")

// ErrEmailTaken возвращается при занятом имейле при регистрации
var ErrEmailTaken = errors.New("email already taken")

// ErrInvalidCredentials возвращается при неподходящих кредах при авторизации
var ErrInvalidCredentials = errors.New("invalid user credentials")

// ErrUnauthorized возвращается при попытке выполнить действие без необходимых прав
var ErrUnauthorized = errors.New("unauthorized access")

// ErrForbidden возвращается при попытке выполнить действие с недостаточными правами
var ErrForbidden = errors.New("forbidden action")

// ErrIntakeNotActive возвращается при попытке добавить товар в неактивную приёмку
var ErrIntakeNotActive = errors.New("intake is not active")

// ErrInvalidItemOrder возвращается при попытке удалить товар не в порядке LIFO
var ErrInvalidItemOrder = errors.New("can only remove last added item (LIFO order)")
