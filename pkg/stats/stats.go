package stats

import "github.com/pyuldashev912/bank/pkg/types"

// Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	var total types.Money
	for _, payment := range payments {
		total += payment.Amount
	}

	return total / types.Money(len(payments))
}

// TotalInCategory находит сумму покупок в определенной категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var total types.Money
	for _, payment := range payments {
		if payment.Category == category {
			total += payment.Amount
		}
	}
	return total
}
