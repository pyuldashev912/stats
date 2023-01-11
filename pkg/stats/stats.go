package stats

import "github.com/pyuldashev912/bank/v2/pkg/types"

// Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	var total types.Money
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			total += payment.Amount
		}
	}

	return total / types.Money(len(payments))
}

// TotalInCategory находит сумму покупок в определенной категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var total types.Money
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		if payment.Category == category {
			total += payment.Amount
		}
	}
	return total
}

// FilterByCategory возвращает платежи в указанной категории
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	filtered := make([]types.Payment, 0, len(payments))
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}

	return filtered
}
