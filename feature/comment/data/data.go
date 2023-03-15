package data

import (
	"airbnb/feature/comment"
	"errors"
	"log"

	"gorm.io/gorm"
)

type commentQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) comment.CommentData {
	return &commentQuery{
		db: db,
	}
}

func (uq *commentQuery) Add(newComment comment.Core) error {
	cnv := CoreToData(newComment)
	if err := uq.db.Create(&cnv).Error; err != nil {
		log.Println("add comment query error", err.Error())
		return err
	}

	return nil
}

func (uq *commentQuery) Delete(userId, id uint) error {
	qry := uq.db.Where("user_id = ?", userId).Delete(&Comment{}, id)

	if qry.RowsAffected <= 0 {
		log.Println("no rows affected")
		return errors.New("data not found")
	}

	if err := qry.Error; err != nil {
		log.Println("delete comment query error")
		return errors.New("data not found")
	}
	return nil
}

func (uq *commentQuery) GetHomestayComments(homeId uint) ([]comment.Core, error) {
	res := []HomeComment{}
	// err := uq.db.Where("user_id = ?", userID).Find(&res).Error
	err := uq.db.Raw("SELECT comments.id, comments.user_id, comments.rate, comments.description FROM comments WHERE products.deleted_at is NULL AND comments.home_id=?", homeId).Scan(&res).Error

	if err != nil {
		log.Println("query error", err.Error())
		return []comment.Core{}, errors.New("server error")
	}

	result := ListHomeCommentToCore(res)
	return result, nil
}
