package stats

import (
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
