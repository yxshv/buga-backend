package main

import (
	"log"
	"os"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/joho/godotenv"
)


type client struct{}
type message struct {
	content string
	by      *websocket.Conn
}

var connections = make(map[*websocket.Conn]client)
var register = make(chan *websocket.Conn)
var broadcast = make(chan message)
var unregister = make(chan *websocket.Conn)

func main() {

	godotenv.Load()
	PORT := os.Getenv("PORT")

	app := fiber.New()

	go WebsocketHub()

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		WebSocket(c)
	}))

	if PORT == "" {
		PORT = "3000"
	}

	log.Fatal(app.Listen(fmt.Sprintf(":%s", PORT)))
}
