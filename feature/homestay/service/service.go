package service

import (
	home "airbnb/feature/homestay"
	"errors"

	"github.com/go-playground/validator/v10"
)

type HomeService struct {
	homeRepository home.RepositoryInterface
	validasi       *validator.Validate
}

func NewRoom(repo home.RepositoryInterface) home.ServiceInterface {
	return &HomeService{
		homeRepository: repo,
		validasi:       validator.New(),
	}
}

// CreateRoom implements room.ServiceInterface
func (service *HomeService) Create(input home.HomestayEntity) error {

	err := service.homeRepository.Create(input)

	if err != nil {
		return err
	}
	return nil
}
func (service *HomeService) GetAll() (data []home.HomestayEntity, err error) {
	data, err = service.homeRepository.GetAll()
	return
}

func (service *HomeService) Update(id int, input home.HomestayEntity) error {
	err := service.homeRepository.Update(id, input)
	if err != nil {
		return errors.New("id not found")
	}
	return nil
}

func (service *HomeService) Delete(id int) error {
	err := service.homeRepository.Delete(id)
	return err
}

func (service *HomeService) GetUserHome(id uint) (home.HomestayEntity, error) {
	data, err := service.homeRepository.GetUserHome(id)
	return data, err
}
