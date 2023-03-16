package delivery

import img "airbnb/feature/images"

type Respon struct {
	ID        uint
	Image_url string
	HomeID    uint
}

func GambarCoreToGambarRespon(dataCore img.ImagesCore) Respon {
	return Respon{
		ID:        dataCore.ID,
		Image_url: dataCore.Image_url,
		HomeID:    dataCore.HomeID,
	}
}
func ListGambarCoreToGambarRespon(dataCore []img.ImagesCore) []Respon {
	var ResponData []Respon

	for _, value := range dataCore {
		ResponData = append(ResponData, GambarCoreToGambarRespon(value))
	}
	return ResponData
}
