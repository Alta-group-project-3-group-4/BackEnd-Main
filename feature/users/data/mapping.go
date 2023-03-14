package data

import (
	"airbnb/feature/users"
)

func UserEntityToUser(userEntity users.UserEntity) User {
	return User{
		HomestayId: userEntity.HomestayId,
		Name:       userEntity.Name,
		Email:      userEntity.Email,
		Password:   userEntity.Password,
		Address:    userEntity.Address,
		About:      userEntity.About,
	}
}

func UserToUserEntity(user User) users.UserEntity {

	result := users.UserEntity{
		Id:         user.ID,
		HomestayId: user.HomestayId,
		Name:       user.Name,
		Email:      user.Email,
		Password:   user.Password,
		Address:    user.Address,
		About:      user.About,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}

	// if !reflect.ValueOf(user.Team).IsZero() {
	// 	result.Team = teams.TeamEntity{
	// 		Id:   user.Team.ID,
	// 		Name: user.Team.Name,
	// 	}
	// }

	return result
}

func ListUserToUserEntity(user []User) []users.UserEntity {
	var userEntity []users.UserEntity
	for _, v := range user {
		userEntity = append(userEntity, UserToUserEntity(v))
	}
	return userEntity
}
