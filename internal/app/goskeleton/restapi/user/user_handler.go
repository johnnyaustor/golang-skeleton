package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func RegisterHandlers(e *echo.Echo, service Service) {
	h := handler{
		service: service,
	}
	e.GET("/users", h.get)
}

func (h handler) get(c echo.Context) error {
	return c.String(http.StatusOK, h.service.Get())
}
