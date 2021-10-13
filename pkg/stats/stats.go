package stats

import "github.com/bakhodur-nazriev/bank/v2/pkg/types"

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
