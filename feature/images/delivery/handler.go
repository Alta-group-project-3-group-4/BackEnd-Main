package delivery

import (
	img "airbnb/feature/images"
	"airbnb/utils/helpers"
	helper "airbnb/utils/s3"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Deliv struct {
	Service img.ServiceInterface
}

func New(service img.ServiceInterface) *Deliv {
	return &Deliv{
		Service: service,
	}

}

func (delivery *Deliv) CreateGambar(c echo.Context) error {

	image := Request{}
	homeid := c.FormValue("home_id")
	id, _ := strconv.Atoi(homeid)
	image.HomeID = uint(id)
	file, _ := c.FormFile("image_url")

	if file != nil {

		res, err := helper.GetUrlImagesFromAWS(*file)
		fmt.Println("print helper s3", res)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error insert into database"+err.Error()))
		}
		image.Image_url = res
	} else {
		image.Image_url = "https://rischiyudap1.s3.ap-southeast-1.amazonaws.com/homestay/Albert+einstein+1.jpg"
	}
	result := image.reqToCore()

	err := delivery.Service.Create(result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error insert into database"+err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Success Add Gambar", nil))

}
func (delivery *Deliv) Getall(c echo.Context) error {

	result, err := delivery.Service.Getall() //memanggil fungsi service yang ada di folder service

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("erorr read data"))
	}
	var ResponData = ListGambarCoreToGambarRespon(result)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("berhasil membaca  user", ResponData))

}
