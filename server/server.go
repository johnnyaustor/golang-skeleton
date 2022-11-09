package server

import (
	"log"

	"github.com/johnnyaustor/golang-skeleton/server/api"
	"github.com/johnnyaustor/golang-skeleton/server/configuration"
	"github.com/johnnyaustor/golang-skeleton/server/database"
	"github.com/labstack/echo/v4"
)

func init() {
	// read env
	configuration.ReadConfiguration()
}

func RunServer() {
	// Connect to Database
	db := database.ConnectDatabase()

	log.Println("Run Server")
	e := echo.New()

	// Initial Beans
	api.InitialBeans(e, db)

	e.Logger.Fatal(e.Start(":1323"))
}
