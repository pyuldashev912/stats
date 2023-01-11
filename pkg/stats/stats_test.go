package stats

import (
	"reflect"
	"testing"

	"github.com/pyuldashev912/bank/v2/pkg/types"
)

func TestAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Amount: 12468, Category: "авто"},
		{ID: 2, Amount: 1500, Category: "аптека"},
		{ID: 3, Amount: 37414, Category: "супермаркет"},
		{ID: 4, Amount: 15400, Category: "кафе"},
	}
	result := Avg(payments)
	expected := types.Money(16695)

	if result != expected {
		t.Errorf("Invalid answer: expected %d, got %d", expected, result)
	}
}

func TestTotalInCategory(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Amount: 12468, Category: "авто"},
		{ID: 2, Amount: 1500, Category: "аптека"},
		{ID: 3, Amount: 37414, Category: "супермаркет"},
		{ID: 4, Amount: 15400, Category: "кафе"},
		{ID: 5, Amount: 14700, Category: "супермаркет"},
		{ID: 6, Amount: 45400, Category: "аптека"},
		{ID: 7, Amount: 9800, Category: "кафе"},
		{ID: 8, Amount: 25000, Category: "супермаркет"},
	}

	expected := types.Money(77114)
	result := TotalInCategory(payments, "супермаркет")

	if result != expected {
		t.Errorf("Invalid answer: expected %d, got %d", expected, result)
	}
}

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{Amount: 9575, Category: "home"},
		{Amount: 1200, Category: "taxi"},
		{Amount: 25400, Category: "sweets"},
		{Amount: 16980, Category: "home"},
		{Amount: 9815, Category: "drugstore"},
		{Amount: 1000, Category: "taxi"},
		{Amount: 3500, Category: "taxi"},
		{Amount: 16500, Category: "sweets"},
		{Amount: 31654, Category: "home"},
	}

	got := CategoriesAvg(payments)
	expect := map[types.Category]types.Money{
		"home":      19403,
		"taxi":      1900,
		"sweets":    20950,
		"drugstore": 9815,
	}

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("Error: got %v, expected %v", got, expect)
	}
}

func TestPeriodsDynamic_Simple1(t *testing.T) {
	first := map[types.Category]types.Money{"auto": 10, "food": 20}
	second := map[types.Category]types.Money{"auto": 5, "food": 3}
	expect := map[types.Category]types.Money{"auto": -5, "food": -17}

	result := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Error: got %v, expected %v", result, expect)
	}
}

func TestPeriodsDynamic_Simple2(t *testing.T) {
	first := map[types.Category]types.Money{"auto": 10, "food": 20}
	second := map[types.Category]types.Money{"auto": 20, "food": 20}
	expect := map[types.Category]types.Money{"auto": 10, "food": 0}

	result := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Error: got %v, expected %v", result, expect)
	}
}

func TestPeriodsDynamic_Tricky1(t *testing.T) {
	first := map[types.Category]types.Money{"auto": 10, "food": 20}
	second := map[types.Category]types.Money{"food": 20}
	expect := map[types.Category]types.Money{"auto": -10, "food": 0}

	result := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Error: got %v, expected %v", result, expect)
	}
}

func TestPeriodsDynamic_Tricky2(t *testing.T) {
	first := map[types.Category]types.Money{"auto": 10, "food": 20}
	second := map[types.Category]types.Money{"auto": 10, "food": 25, "mobile": 5}
	expect := map[types.Category]types.Money{"auto": 0, "food": 5, "mobile": 5}

	result := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Error: got %v, expected %v", result, expect)
	}
}
