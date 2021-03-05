package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lcsval/boexpenses/pkg/handler"
)

func Expenses(app *fiber.App) {
	var h handler.ExpenseHandler
	r := app.Group("/expenses")
	r.Get("/", h.Index)
}
