package room

import "github.com/labstack/echo/v4"

type Room struct {
	RoomName   string `gorm:"column:room_name;type:varchar(255)" json:"room_name"`
	RoomNumber string `gorm:"column:room_number;type:varchar(255)" json:"room_number"`
}

type RoomHandlerInterface interface {
	CreateRoom() echo.HandlerFunc
}

type RoomServiceInterface interface {
	CreateRoom(newData Room) (*Room, error)
}

type RoomDataInterface interface {
	CreateRoom(newData Room) (*Room, error)
}
