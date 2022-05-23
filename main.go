package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun/driver/pgdriver"
)

var db *sql.DB

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
	var DB_URL string = os.Getenv("DB_URL")

	db = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(DB_URL)))

	app := fiber.New()

	go WebsocketHub()

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		WebSocket(c)
	}))

	app.Get("/message/get", func(c *fiber.Ctx) error {
		var (
			r   []string
			err error
		)

		if r, err = GetMessages(); err != nil {
			log.Println("Error while getting messages: ", err)
			return err
		}

		return c.JSON(fiber.Map{
			"messages": r,
		})

	})

	log.Fatal(app.Listen(":3000"))
}
