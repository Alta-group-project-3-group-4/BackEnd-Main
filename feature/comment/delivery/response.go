package handler

import "airbnb/feature/comment"

type CommentResponse struct {
	ID          uint    `json:"id"`
	UserId      uint    `json:"user_id"`
	Rate        float32 `json:"rate"`
	Description string  `json:"description"`
}

func ToCommentResponse(data comment.Core) CommentResponse {
	return CommentResponse{
		ID:          data.ID,
		UserId:      data.UserId,
		Rate:        data.Rate,
		Description: data.Description,
	}
}

func ListofCommentToResponse(dataCore []comment.Core) []CommentResponse {
	var DataResponse []CommentResponse

	for _, value := range dataCore {
		DataResponse = append(DataResponse, ToCommentResponse(value))
	}
	return DataResponse
}
