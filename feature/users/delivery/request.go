package delivery

import "airbnb/feature/users"

type UserRequest struct {
	HomestayId uint   `json:"homestay_id" form:"homestay_id"`
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Address    string `json:"address" form:"address"`
	About      string `json:"about" form:"about"`
}

func UserRequestToUserEntity(userRequest UserRequest) users.UserEntity {
	return users.UserEntity{
		HomestayId: userRequest.HomestayId,
		Name:       userRequest.Name,
		Email:      userRequest.Email,
		Password:   userRequest.Password,
		Address:    userRequest.Address,
		About:      userRequest.About,
	}
}
