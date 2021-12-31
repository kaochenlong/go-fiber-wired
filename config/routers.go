package config

import (
	"go_fiber_wired/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	page := new(controllers.PageController)

	app.Get("/", page.Home)
}
