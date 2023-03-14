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
func (cuc *commentUseCase) Add(token interface{}, newComment comment.Core) error {
	userId := helper.ExtractToken(token)
	if userId <= 0 {
		return errors.New("id tidak valid")
	}

	err := cuc.qry.Add(uint(userId), newComment)
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

func (cuc *commentUseCase) Delete(token interface{}, ID uint) error {
	return nil
}
