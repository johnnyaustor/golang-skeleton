package goskeleton

import (
	"log"

	"github.com/johnnyaustor/golang-skeleton/internal/app/goskeleton/configuration"
	"github.com/johnnyaustor/golang-skeleton/internal/app/goskeleton/database"
	"github.com/johnnyaustor/golang-skeleton/internal/app/goskeleton/restapi"
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
	restapi.NewRoutes(e, db)

	e.Logger.Fatal(e.Start(":1323"))
}
