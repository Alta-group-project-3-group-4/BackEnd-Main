package handler

import (
	"airbnb/feature/comment"
	"airbnb/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type commentControl struct {
	srv comment.CommentService
}

func New(srv comment.CommentService) comment.CommentHandler {
	return &commentControl{
		srv: srv,
	}
}

func (cc *commentControl) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		input := AddCommentRequest{}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "format inputan salah",
			})
		}

		err2 := cc.srv.Add(token, *ToCore(input))
		if err2 != nil {
			return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err2.Error()))
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "berhasil menambahkan ulasan",
		})
	}
}

func (uc *commentControl) Delete() echo.HandlerFunc {
	return nil
}
