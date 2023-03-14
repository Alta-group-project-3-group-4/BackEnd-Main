package data

import (
	"airbnb/feature/comment"
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

func (uq *commentQuery) Add(userId uint, newComment comment.Core) error {
	cnv := CoreToData(newComment)
	cnv.UserId = userId
	if err := uq.db.Create(&cnv).Error; err != nil {
		log.Println("add comment query error", err.Error())
		return err
	}

	return nil
}

func (uq *commentQuery) Delete(userId, ID uint) error {
	return nil
}
