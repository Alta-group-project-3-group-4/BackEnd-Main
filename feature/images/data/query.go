package data

import (
	img "airbnb/feature/images"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) img.RepositoryInterface {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) Create(input img.ImagesCore) error {
	HomeGorm := CoretoModel(input)
	fmt.Println("gambar", HomeGorm)
	tx := repo.db.Create(&HomeGorm)

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}

	return nil
}

func (repo *Repository) Getall() (data []img.ImagesCore, err error) {
	var image []Image
	tx := repo.db.Unscoped().Find(&image)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var DataCore = ListModelTOCore(image)
	return DataCore, nil

}
