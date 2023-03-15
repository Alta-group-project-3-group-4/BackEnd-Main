package service

import (
	"airbnb/feature/comment"
	"errors"
	"strings"
)

type commentUseCase struct {
	qry comment.CommentData
}

func New(ud comment.CommentData) comment.CommentService {
	return &commentUseCase{
		qry: ud,
	}
}

// perlu mencari extract token
func (cuc *commentUseCase) Add(newComment comment.Core) error {

	err := cuc.qry.Add(newComment)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "data already exist"
		} else {
			msg = "server problem"
		}
		return errors.New(msg)
	}

	return nil
}

func (cuc *commentUseCase) Delete(userId, id uint) error {
	err := cuc.qry.Delete(userId, id)
	if err != nil {
		return err
	}

	return nil
}
