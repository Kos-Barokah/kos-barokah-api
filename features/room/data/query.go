package data

import (
	"errors"
	"kos-barokah-api/features/room"

	"gorm.io/gorm"
)

type RoomData struct {
	db *gorm.DB
}

func NewData(db *gorm.DB) room.RoomDataInterface {
	return &RoomData{
		db: db,
	}
}

func (rd *RoomData) CreateRoom(newData room.Room) (*room.Room, error) {

	room := new(room.Room)

	room.RoomName = newData.RoomName
	room.RoomNumber = newData.RoomNumber

	err := rd.db.Create(room).Error

	if err != nil {
		return nil, errors.New("create room failed")
	}

	return room, nil
}
