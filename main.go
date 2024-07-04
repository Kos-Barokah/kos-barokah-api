package main

import (
	"fmt"
	"kos-barokah-api/configs"
	"kos-barokah-api/helper"
	email "kos-barokah-api/helper/email"
	encrypt "kos-barokah-api/helper/encrypt"
	"kos-barokah-api/routes"
	"kos-barokah-api/utils/database"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	dataUser "kos-barokah-api/features/users/data"
	handlerUser "kos-barokah-api/features/users/handler"
	serviceUser "kos-barokah-api/features/users/service"

	dataRoom "kos-barokah-api/features/room/data"
	handlerRoom "kos-barokah-api/features/room/handler"
	serviceRoom "kos-barokah-api/features/room/service"
)

func main() {
	e := echo.New()

	var config = configs.InitConfig()

	db, err := database.InitDB(*config)
	if err != nil {
		e.Logger.Fatal("Cannot run database: ", err.Error())
	}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Endpoint not found", nil))
	})

	e.GET("/api", func(c echo.Context) error {
		return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Endpoint not found", nil))
	})

	e.GET("/api/v1", func(c echo.Context) error {
		return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Endpoint not found", nil))
	})

	var encrypt = encrypt.New()
	var email = email.New(*config)
	// var bucket = bucket.InitBucket(*config)

	jwtInterface := helper.New(config.Secret, config.RefSecret)

	userModel := dataUser.NewData(db)
	roomModel := dataRoom.NewData(db)

	userServices := serviceUser.NewService(userModel, jwtInterface, email, encrypt)
	roomServices := serviceRoom.NewService(roomModel, jwtInterface, email, encrypt)

	userController := handlerUser.NewHandler(userServices, jwtInterface)
	roomController := handlerRoom.NewHandler(roomServices, jwtInterface)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	group := e.Group("/api/v1")

	routes.RouteUser(group, userController, *config)
	routes.RouteRoom(group, roomController, *config)

	fmt.Println("tes")

	e.Logger.Debug(db)

	e.Logger.Info(fmt.Sprintf("Listening in port :%d", config.ServerPort))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
