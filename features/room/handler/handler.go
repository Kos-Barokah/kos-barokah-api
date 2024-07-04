package handler

import (
	"kos-barokah-api/features/room"
	"kos-barokah-api/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RoomHandler struct {
	s room.RoomServiceInterface
	// sc  customer.CustomerServiceInterface
	jwt helper.JWTInterface
}

func NewHandler(service room.RoomServiceInterface /* sci customer.CustomerServiceInterface,*/, jwt helper.JWTInterface) room.RoomHandlerInterface {
	return &RoomHandler{
		s: service,
		/*sc:  sci,*/
		jwt: jwt,
	}
}

func (rh *RoomHandler) CreateRoom() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input = new(RoomRequest)

		err := c.Bind(input)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Invalid user input", nil))
		}

		var room = new(room.Room)

		room.RoomName = input.RoomName
		room.RoomNumber = input.RoomNumber

		result, err := rh.s.CreateRoom(*room)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		var response = new(RoomResponse)

		response.RoomName = result.RoomName
		response.RoomNumber = result.RoomNumber

		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "alam kontol", response))
	}
}
