package gambar

import (
	"time"
)

type ImagesCore struct {
	ID        uint
	Image_url string `validate:"required"`
	HomeID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	Create(input ImagesCore) error
	Getall() (data []ImagesCore, err error)
}

type RepositoryInterface interface {
	Create(input ImagesCore) error
	Getall() (data []ImagesCore, err error)
}
