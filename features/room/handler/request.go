package handler

type RoomRequest struct {
	RoomName   string `json:"room_name" form:"room_name"`
	RoomNumber string `json:"room_number" form:"room_number"`
}
