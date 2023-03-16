package homestay

import (
	"time"
)

type HomestayEntity struct {
	Homestay_Id uint
	UserId      uint
	Name        string  `validate:"required"`
	Price       float64 `validate:"required"`
	Deskripsi   string  `validate:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ServiceInterface interface {
	Create(input HomestayEntity) error
	GetAll() (data []HomestayEntity, err error)
	Update(id int, input HomestayEntity) error
	Delete(id int) error
	GetUserHome(id uint) (HomestayEntity, error)
}

type RepositoryInterface interface {
	Create(input HomestayEntity) error
	GetAll() (data []HomestayEntity, err error)
	Update(id int, input HomestayEntity) error
	Delete(id int) error
	GetUserHome(id uint) (HomestayEntity, error)
}
