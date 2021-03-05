package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/lcsval/goexpenses/pkg/connector"
	"github.com/lcsval/goexpenses/pkg/route"
)

func main() {
	os.Setenv("DB_NAME", "expenses")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASS", "root")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_MIGRATE", "false")

	fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_PASS:", os.Getenv("DB_PASS"))
	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_MIGRATE:", os.Getenv("DB_MIGRATE"))

	app := fiber.New()
	db := connector.Connect()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "Expenses",
		})
	})

	route.Expenses(app, db)
	app.Listen(":3000")
}
