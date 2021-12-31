package config

import (
	"go_fiber_wired/app/controllers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func SetupRoutes(app *fiber.App) {
	page := new(controllers.PageController)

	app.Get("/", page.Home)
}

func SetupWebsocket(app *fiber.App) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		var msgType int
		var msg []byte
		var err error

		for {
			if msgType, msg, err = c.ReadMessage(); err != nil {
				log.Println("read: ", err)
				break
			}

			log.Printf("Message Type: %d, rece: %s", msgType, msg)

			if err = c.WriteMessage(msgType, msg); err != nil {
				log.Println("write: ", err)
				break
			}
		}
	}))
}
