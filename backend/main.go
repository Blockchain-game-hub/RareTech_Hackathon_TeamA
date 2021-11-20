package main

import (
	am "github.com/funobu/teamA/backend/auth"
	api "github.com/funobu/teamA/backend/web/api"
	"github.com/funobu/teamA/backend/web/ws"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	authorized := e.Group("/api")
	authorized.GET("/room:id", api.FindRoomOneById, am.Validate())

	// process game
	match := e.Group("/match")
	match.Any("/room:id", func(c echo.Context) error {
		room := c.Param("id")
		server := ws.Connect(room)
		server.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":8000"))
}
