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

// CategoriesTotal возвращает сумму платежей по каждой категории
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}

	return categories
}

// CategoriesAvg считает среднюю сумму платежа по каждой категории
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	result := CategoriesTotal(payments)
	paymentCategoryCounter := map[types.Category]int{}

	// Находим количество покупок по каждой категории
	for _, payment := range payments {
		paymentCategoryCounter[payment.Category]++
	}

	for key := range result {
		result[key] /= types.Money(paymentCategoryCounter[key])
	}

	return result
}

// PeriodsDynamic сравнивает расходы по категориям за два периода
func PeriodsDynamic(
	first map[types.Category]types.Money, second map[types.Category]types.Money,
) map[types.Category]types.Money {

	// Создаем глубокую копию второго периода
	result := make(map[types.Category]types.Money, len(second))
	for k, v := range second {
		result[k] = v
	}

	for key := range first {
		if _, ok := result[key]; ok {
			result[key] -= first[key]
		} else {
			result[key] = -first[key]
		}
	}

	return result
}
