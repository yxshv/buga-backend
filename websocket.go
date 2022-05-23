package main

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

func WebsocketHub() {
	for {
		select {
		case c := <-register:
			connections[c] = client{}
			log.Println("client registered")
		case msg := <-broadcast:
			for c := range connections {
				if c == msg.by {
					continue
				}
				if err := c.WriteMessage(websocket.TextMessage, []byte(msg.content)); err != nil {
					log.Println("Error while sending message: ", err)

					c.WriteMessage(websocket.CloseMessage, []byte{})
					c.Close()
					delete(connections, c)
				}
			}
		case c := <-unregister:
			delete(connections, c)

			log.Println("client unregistered")
		}
	}
}

func WebSocket(c *websocket.Conn) {
	defer func() {
		unregister <- c
		c.Close()
	}()

	register <- c
}
