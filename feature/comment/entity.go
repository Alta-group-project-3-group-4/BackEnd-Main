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
	Add(token interface{}, newComment Core) error
	Delete(token interface{}, ID uint) error
}

type CommentData interface {
	Add(userId uint, newComment Core) error
	Delete(userId, ID uint) error
}
