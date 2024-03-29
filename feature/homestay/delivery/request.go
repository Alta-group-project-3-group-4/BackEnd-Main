package delivery

import (
	"airbnb/feature/homestay"
	"time"
)

type CreateFormat struct {
	UserId    uint
	Name      string `json:"name" form:"name"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	Price     int    `json:"price" form:"price"`
}

type UpdateFormat struct {
	UserId    uint
	ID        uint   `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	Price     int    `json:"price" form:"price"`
}
type GetId struct {
	id uint `param:"id"`
}

func ToDomain(i interface{}) homestay.HomestayEntity {
	switch i.(type) {
	case CreateFormat:
		cnv := i.(CreateFormat)
		return homestay.HomestayEntity{
			Name:      cnv.Name,
			Price:     float64(cnv.Price),
			Deskripsi: cnv.Deskripsi,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}
	case GetId:
		cnv := i.(GetId)
		return homestay.HomestayEntity{Homestay_Id: cnv.id}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return homestay.HomestayEntity{
			Name:      cnv.Name,
			Price:     float64(cnv.Price),
			Deskripsi: cnv.Deskripsi}

	}
	return homestay.HomestayEntity{}
}
