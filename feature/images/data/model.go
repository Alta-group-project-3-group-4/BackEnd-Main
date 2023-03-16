package data

import (
	home "airbnb/feature/homestay/data"
	img "airbnb/feature/images"

	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Image_url string
	HomeID    uint
	Home      *home.Homestay `gorm:"foreignKey:HomeID"`
}

func CoretoModel(dataCore img.ImagesCore) Image {
	classGorm := Image{

		Image_url: dataCore.Image_url,
		HomeID:    dataCore.HomeID,
	}
	return classGorm
}

func (dataModel Image) ModeltoCore() img.ImagesCore {
	return img.ImagesCore{
		Image_url: dataModel.Image_url,
		HomeID:    dataModel.HomeID,
	}

}

func ListModelTOCore(dataModel []Image) []img.ImagesCore { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []img.ImagesCore
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModeltoCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
