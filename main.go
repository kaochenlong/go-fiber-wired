package main

import (
	"go_fiber_wired/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func main() {
	config.BuildAssets()

	app := fiber.New(fiber.Config{
		Views:         html.New("app/views", ".html"),
		ViewsLayout:   "layouts/application",
		CaseSensitive: true,
	})

	app.Use(logger.New())

	config.SetupRoutes(app)
	config.SetupWebsocket(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
