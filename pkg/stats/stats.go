package stats

import (
	"github.com/bakhodur-nazriev/bank/v2/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) (avgSum types.Money) {
	var total types.Money
	var sum []types.Money

	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		} else {
			total += payment.Amount
		}

		sum = append(sum, total)
	}

	avgSum = total / types.Money(len(sum))
	return
}

// TotalInCategory находит сумму покупок в определённой каторегории.
func TotalInCategory(payments []types.Payment, category types.Category) (res types.Money) {
	for _, payment := range payments {
		if category == payment.Category {
			if payment.Status == types.StatusFail {
				continue
			} else {
				res += payment.Amount
			}
		}
	}

	return
}

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
