package comment

import "github.com/labstack/echo/v4"

type Core struct {
	ID          uint
	UserId      uint
	HomeId      uint
	Rate        float32
	Description string
}

type CommentHandler interface {
	Add() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type CommentService interface {
	Add(newComment Core) error
	Delete(userId, id uint) error
}

type CommentData interface {
	Add(newComment Core) error
	Delete(userId, id uint) error
}
