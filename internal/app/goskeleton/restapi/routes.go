package restapi

import (
	"github.com/johnnyaustor/golang-skeleton/internal/app/goskeleton/database"
	"github.com/johnnyaustor/golang-skeleton/internal/app/goskeleton/restapi/user"
	"github.com/labstack/echo/v4"
)

func NewRoutes(e *echo.Echo, db *database.DB) {
	user.RegisterHandlers(e, user.NewService(user.NewRepository(db)))
}
