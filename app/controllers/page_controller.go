package controllers

import "github.com/gofiber/fiber/v2"

type PageController struct{}

func (c *PageController) Home(ctx *fiber.Ctx) error {
	return ctx.Render("pages/home", nil)
}
