package delivery

import "airbnb/feature/homestay"

type CreateResponse struct {
	Name      string `json:"name"`
	Deskripsi string `json:"deskripsi"`
	Price     int    `json:"price"`
}

type UpdateResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Deskripsi string `json:"deskripsi"`
	Price     int    `json:"price"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "register":
		cnv := core.(homestay.HomestayEntity)
		res = CreateResponse{
			Name:      cnv.Name,
			Price:     int(cnv.Price),
			Deskripsi: cnv.Deskripsi,
		}
	case "update":
		cnv := core.(homestay.HomestayEntity)
		res = UpdateResponse{
			ID:        cnv.Homestay_Id,
			Name:      cnv.Name,
			Price:     int(cnv.Price),
			Deskripsi: cnv.Deskripsi,
		}
	}
	return res
}

func ToResponseHomestay(core interface{}) interface{} {
	var res interface{}
	var arr []UpdateResponse
	val := core.([]homestay.HomestayEntity)
	for _, cnv := range val {
		arr = append(arr, UpdateResponse{
			ID:        cnv.Homestay_Id,
			Name:      cnv.Name,
			Deskripsi: cnv.Deskripsi,
			Price:     int(cnv.Price),
		})
	}
	res = arr
	return res
}
