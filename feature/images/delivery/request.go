package delivery

import (
	img "airbnb/feature/images"
	imag "airbnb/feature/images/data"
)

type Request struct {
	Image_url string `json:"image_url" form:"image_url"`
	HomeID    uint   `json:"home_id " form:"home_id "`
}

func (data *Request) reqToCore() img.ImagesCore {
	return img.ImagesCore{
		Image_url: data.Image_url,
		HomeID:    data.HomeID,
	}
}

func ToCore(data interface{}) *imag.Image {
	res := imag.Image{}

	switch data.(type) {
	case Request:
		cnv := data.(Request)
		res.HomeID = cnv.HomeID
		res.Image_url = cnv.Image_url

	default:
		return nil
	}

	return &res
}
