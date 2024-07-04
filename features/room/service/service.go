package service

import (
	"kos-barokah-api/features/room"
	"kos-barokah-api/helper"
	email "kos-barokah-api/helper/email"
	encrypt "kos-barokah-api/helper/encrypt"
)

type RoomService struct {
	d     room.RoomDataInterface
	j     helper.JWTInterface
	e     encrypt.HashInterface
	email email.EmailInterface
}

func NewService(data room.RoomDataInterface, jwt helper.JWTInterface, email email.EmailInterface, encrypt encrypt.HashInterface) room.RoomServiceInterface {
	return &RoomService{
		d:     data,
		j:     jwt,
		email: email,
		e:     encrypt,
	}
}

func (rs *RoomService) CreateRoom(newData room.Room) (*room.Room, error) {

	result, err := rs.d.CreateRoom(newData)

	if err != nil {
		return nil, err
	}

	return result, nil
}
