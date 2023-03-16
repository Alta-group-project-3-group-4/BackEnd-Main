package delivery

import img "airbnb/feature/images"

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
