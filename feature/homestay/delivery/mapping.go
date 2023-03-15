package delivery

import (
	home "airbnb/feature/homestay"
	homestay "airbnb/feature/homestay/data"
)

func RoomEntityToRoom(homeEntity home.HomestayEntity) homestay.Homestay {
	result := homestay.Homestay{
		Name:        homeEntity.Name,
		Price:       homeEntity.Price,
		Description: homeEntity.Deskripsi,
	}

	return result
}

func RoomToRoomEntity(room homestay.Homestay) home.HomestayEntity {
	result := home.HomestayEntity{
		Name:      room.Name,
		Price:     room.Price,
		Deskripsi: room.Description,
	}

	// if !reflect.ValueOf(room.User).IsZero() {
	// 	result.User = users.UserEntity{
	// 		Name:  room.User.Name,
	// 		Email: room.User.Email,
	// 	}
	// }

	return result
}
