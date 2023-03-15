package handler

import "airbnb/feature/comment"

type AddCommentRequest struct {
	UserId      uint
	HomeId      uint    `json:"home_id" form:"home_id"`
	Rate        float32 `json:"rate" form:"rate"`
	Description string  `json:"description" form:"description"`
}

func ToCore(data interface{}) *comment.Core {
	res := comment.Core{}

	switch data.(type) {
	case AddCommentRequest:
		cnv := data.(AddCommentRequest)
		res.HomeId = cnv.HomeId
		res.Rate = cnv.Rate
		res.Description = cnv.Description
		res.UserId = cnv.UserId
	default:
		return nil
	}
	return &res
}
