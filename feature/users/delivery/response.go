package delivery

import (
	"airbnb/feature/users"
)

type UserResponse struct {
	Id         uint   `json:"id"`
	HomestayId uint   `json:"homestay_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	About      string `json:"about"`
	Address    string `json:"address"`
	// Team        *delivery.TeamResponse `json:"team,omitempty"`
}

func UserEntityToUserResponse(userEntity users.UserEntity) UserResponse {
	result := UserResponse{
		Id:         userEntity.Id,
		HomestayId: userEntity.HomestayId,
		Name:       userEntity.Name,
		Email:      userEntity.Email,
		About:      userEntity.About,
		Address:    userEntity.Address,
	}

	// if !reflect.ValueOf(userEntity.Team).IsZero() {
	// 	result.Team = &delivery.TeamResponse{
	// 		Id:   userEntity.Team.Id,
	// 		Name: userEntity.Team.Name,
	// 	}
	// }

	return result
}

func ListUserEntityToUserResponse(dataCore []users.UserEntity) []UserResponse {
	var dataResponses []UserResponse
	for _, v := range dataCore {
		dataResponses = append(dataResponses, UserEntityToUserResponse(v))
	}
	return dataResponses
}
