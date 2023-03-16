package service

import (
	image "airbnb/feature/images"
	"errors"

	"github.com/go-playground/validator"
)

type HomeService struct {
	Repository image.RepositoryInterface
	validasi   *validator.Validate
}

func NewGambar(repo image.RepositoryInterface) image.ServiceInterface {
	return &HomeService{
		Repository: repo,
		validasi:   validator.New(),
	}
}

func (service *HomeService) Create(input image.ImagesCore) error {
	if validateERR := service.validasi.Struct(input); validateERR != nil {
		return validateERR
	}

	errCreate := service.Repository.Create(input)
	if errCreate != nil {
		return errors.New("GAGAL MENAMBAH DATA , QUERY ERROR")
	}

	return nil
}

func (service *HomeService) Getall() (data []image.ImagesCore, err error) {
	data, err = service.Repository.Getall()
	return
}
