package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lcsval/boexpenses/pkg/route"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "Expenses",
		})
	})

	route.Expenses(app)
	app.Listen(":3000")
}
