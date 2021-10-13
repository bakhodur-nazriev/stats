package stats

import "github.com/bakhodur-nazriev/bank/pkg/types"

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) (avgSum types.Money) {
	var total types.Money

	for _, payment := range payments {
		total += payment.Amount
	}

	avgSum = total / types.Money(len(payments))
	return
}

// TotalInCategory находит сумму покупок в определённой каторегории.
func TotalInCategory(payments []types.Payment, category types.Category) (res types.Money) {
	for _, payment := range payments {
		if category == payment.Category {
			res += payment.Amount
		}
	}

	return
}
