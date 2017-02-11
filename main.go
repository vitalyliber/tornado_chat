package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/middleware"
)

var (
	upgrader = websocket.Upgrader{}
)

func chat(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", msg)

		// Write
		err = ws.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "../public")
	e.File("/", "public/index.html")
	e.GET("/ws", chat)
	e.Logger.Fatal(e.Start(":3000"))
}