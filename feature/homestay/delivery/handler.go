package delivery

import (
	"airbnb/feature/homestay"
	"airbnb/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Homehandler struct {
	Service homestay.ServiceInterface
}

func New(srv homestay.ServiceInterface) *Homehandler {
	return &Homehandler{
		Service: srv,
	}
}

func (h *Homehandler) CreateRoom(c echo.Context) error {

	homes := CreateFormat{}
	errBind := c.Bind(&homes)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(errBind.Error()))
	}

	err := h.Service.Create(ToDomain(homes))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(errBind.Error()))
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Create Data Success", ToResponseHomestay(err)))

}

func (uc *Homehandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := uc.Service.GetAll()
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
		}

		return c.JSON(http.StatusOK, helpers.ResponseSuccess("Get Data Success", res))
	}
}

func (uc *Homehandler) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		PostID, _ := strconv.Atoi(c.Param("id"))

		res, err := uc.Service.GetUserHome(uint(PostID))
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
		}

		return c.JSON(http.StatusOK, helpers.ResponseSuccess("Get Data Success", res))
	}
}

func (uc *Homehandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		PostID, _ := strconv.Atoi(c.Param("id"))

		del := uc.Service.Delete(PostID)
		if del != nil {
			return c.JSON(http.StatusBadRequest, helpers.ResponseFail(del.Error()))
		}

		return c.JSON(http.StatusOK, helpers.ResponseSuccess("sukses menghapus buku", del))
	}
}
