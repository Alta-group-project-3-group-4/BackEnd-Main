package handler

import (
	middlewares "airbnb/app/middleware"
	"airbnb/feature/comment"
	"airbnb/utils/helpers"
	"net/http"
	"strconv"

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

		input := AddCommentRequest{}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "format inputan salah",
			})
		}

		claim := middlewares.ClaimsToken(c)
		userId := claim.Id
		input.UserId = uint(userId)

		err2 := cc.srv.Add(*ToCore(input))
		if err2 != nil {
			return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err2.Error()))
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "berhasil menambahkan ulasan",
		})
	}
}

func (cc *commentControl) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		_id, _ := strconv.Atoi(c.Param("id"))
		id := uint(_id)

		claim := middlewares.ClaimsToken(c)
		userId := uint(claim.Id)

		if err := cc.srv.Delete(userId, id); err != nil {
			return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
		}

		return c.JSON(http.StatusOK, helpers.ResponseSuccess("Delete Data Success", nil))
	}
}

func (cc *commentControl) GetHomestayComments() echo.HandlerFunc {
	return func(c echo.Context) error {
		// claim := middlewares.ClaimsToken(c)
		// userId := uint(claim.Id)

		_id, _ := strconv.Atoi(c.Param("id"))
		homeId := uint(_id)

		res, err := cc.srv.GetHomestayComments(homeId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    ListofCommentToResponse(res),
			"message": "success show user feedback",
		})
	}
}
