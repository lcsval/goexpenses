package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lcsval/boexpenses/pkg/entity"
)

type ExpenseHandler struct {
}

func (h ExpenseHandler) Index(ctx *fiber.Ctx) error {
	expenses := []entity.Expense{
		{
			ID:         "1",
			Title:      "Lunch at MyFood",
			Total:      14.95,
			Attachment: "photo.jpg",
			CreatedAt:  time.Now(),
		},
	}
	return ctx.JSON(fiber.Map{"data": expenses})
}
