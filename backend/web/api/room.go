package web

import (
	"net/http"

	"github.com/funobu/teamA/backend/usecases"
	"github.com/labstack/echo/v4"
)

func FindRoomOneById(c echo.Context) error {
	id := c.Param("id")
	room, err := usecases.FindRoomOneById(id)
	if err != nil {
		c.Error(err)
	}
	return c.JSON(http.StatusOK, room)
}
