package main

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	upgrader = websocket.Upgrader{}
	room     = Room{}
)

func wsHandler(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	client := &Client{Ws: ws}
	room.AddClient(client)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}

		room.Broadcast(msg)
		fmt.Printf("%s\n", msg)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong \n")
	})

	e.Static("/", "./public")
	e.GET("/ws", wsHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
