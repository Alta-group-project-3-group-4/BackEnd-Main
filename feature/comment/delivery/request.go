package handler

import "airbnb/feature/comment"

type AddCommentRequest struct {
	UserId      uint
	HomeId      uint    `json:"id_homestay" form:"id_homestay"`
	Rate        float32 `json:"penilaian" form:"penilaian"`
	Description string  `json:"ulasan" form:"ulasan"`
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
