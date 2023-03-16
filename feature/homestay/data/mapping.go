package data

import (
	homestay "airbnb/feature/homestay"
)

func HomestayEntityToHome(homeEntity homestay.HomestayEntity) Homestay {
	return Homestay{
		UserId:      homeEntity.UserId,
		Name:        homeEntity.Name,
		Price:       homeEntity.Price,
		Description: homeEntity.Deskripsi,
	}
}

func HomeToHomestayEntity(home Homestay) homestay.HomestayEntity {

	result := homestay.HomestayEntity{
		Name:      home.Name,
		Price:     home.Price,
		Deskripsi: home.Description,
		CreatedAt: home.CreatedAt,
		UpdatedAt: home.UpdatedAt,
	}

	// if !reflect.ValueOf(home.User).IsZero() {
	// 	result.User = users.UserEntity{
	// 		HomestayId: 0,
	// 		Name:       home.User.Name,
	// 		Email:      home.User.Email,
	// 		Homestay:   homestay.HomestayEntity{},
	// 	}
	// }

	return result
}

func ListHomeToHomeEntity(home []Homestay) []homestay.HomestayEntity {
	var homeEntity []homestay.HomestayEntity
	for _, v := range home {
		homeEntity = append(homeEntity, HomeToHomestayEntity(v))
	}
	return homeEntity
}
