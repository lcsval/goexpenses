package validation

import "github.com/lcsval/goexpenses/pkg/entity"

func CreateOrUpdate(expense entity.Expense) []string {
	var messages []string

	if !isValid(expense.Title) {
		messages = append(messages, "Tile is required")
	}

	if !isValid(expense.Total) {
		messages = append(messages, "Total should be greater than zero")
	}

	return messages
}
