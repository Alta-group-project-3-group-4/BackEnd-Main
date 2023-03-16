package delivery

type ReservationRequest struct {
	HomestayId uint `json:"homestay_id" form:"homestay_id"`
	UserId     uint
	CheckIn    string `json:"check_in" form:"check_in"`
	CheckOut   string `json:"check_out" form:"check_out"`
	Price      int    `json:"price" form:"price"`
	Bank       string `json:"bank" form:"bank"`
}
