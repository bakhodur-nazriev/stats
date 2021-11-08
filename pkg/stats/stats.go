package stats

import (
	"github.com/bakhodur-nazriev/bank/v2/pkg/types"
)

// CategoriesTotal возвращает сумму платежей пл каждой категории.
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}

	return categories
}

// FilteredByCategory возвращает платежи в указанной категории.
func FilteredByCategory(payments []types.Payment, categoty types.Category) []types.Payment {
	var filtered []types.Payment

	for _, payment := range payments {
		if payment.Category == categoty {
			filtered = append(filtered, payment)
		}
	}

	return filtered
}

// CategoriesAvg рассчитывает среднюю сумму категории платежа.
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	total := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
		total[payment.Category]++
	}

	for key, category := range categories {
		categories[key] = category / total[key]
	}

	return categories
}

// PeriodsDynamic динамический рассчитывает периуд затрать
func PeriodsDynamic(first, second map[types.Category]types.Money) map[types.Category]types.Money {
	dynamic := map[types.Category]types.Money{}

	for key := range second {
		dynamic[key] = second[key] - first[key]
	}

	for key := range first {
		dynamic[key] = second[key] - first[key]
	}

	return dynamic
}