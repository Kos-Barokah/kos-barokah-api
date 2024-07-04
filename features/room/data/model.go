package data

import (
	"gorm.io/gorm"
)

type Room struct {
	*gorm.Model
	RoomName   string `gorm:"column:room_name;type:varchar(255)"`
	RoomNumber string `gorm:"column:room_number;type:varchar(255)"`
}
