package api

import (
	"github.com/johnnyaustor/golang-skeleton/server/api/user"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitialBeans(e *echo.Echo, db *gorm.DB) {
	user.RegisterHandlers(e, user.NewService(user.NewRepository(db)))
}
