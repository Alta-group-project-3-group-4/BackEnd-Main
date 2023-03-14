package data

import (
	homestay "airbnb/feature/homestay"
	"errors"

	"gorm.io/gorm"
)

type HomeRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) homestay.RepositoryInterface {
	return &HomeRepository{
		db: db,
	}
}

// CreateRoom implements room.RepositoryInterface
func (repo *HomeRepository) Create(input homestay.HomestayEntity) error {

	HomeGorm := HomestayEntityToHome(input)
	tx := repo.db.Create(&HomeGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func (repo *HomeRepository) GetAll() (data []homestay.HomestayEntity, err error) {
	var home []Homestay //mengambil data gorm model(model.go)
	tx := repo.db.Unscoped().Find(&home)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var DataCore = ListHomeToHomeEntity(home) //mengambil data dari gorm model(file repository(model.go))

	return DataCore, nil

}

// UpdateRoom implements room.RepositoryInterface
func (repo *HomeRepository) Update(id int, input homestay.HomestayEntity) error {
	room := HomestayEntityToHome(input)
	tx := repo.db.Model(room).Where("id = ?", id).Updates(room)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
func (repo *HomeRepository) Delete(id int) error {
	var home Homestay
	tx := repo.db.Delete(&home, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// GetUserRoom implements room.RepositoryInterface
func (repo *HomeRepository) GetUserHome(id uint) (homestay.HomestayEntity, error) {
	home := Homestay{}

	tx := repo.db.Preload("Homestay").First(&home, id)
	if tx.Error != nil {
		return homestay.HomestayEntity{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return homestay.HomestayEntity{}, errors.New("id not found")

	}
	return HomeToHomestayEntity(home), nil
}
