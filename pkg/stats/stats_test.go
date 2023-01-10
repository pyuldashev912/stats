package stats

import (
	"fmt"

	"github.com/pyuldashev912/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{ID: 1, Amount: 12468, Category: "авто"},
		{ID: 2, Amount: 1500, Category: "аптека"},
		{ID: 3, Amount: 37414, Category: "супермаркет"},
		{ID: 4, Amount: 15400, Category: "кафе"},
	}
	fmt.Println(Avg(payments))

	// Output:
	// 16695
}

func ExampleTotalInCategory() {
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

	fmt.Println(TotalInCategory(payments, "супермаркет"))

	// Output:
	// 77114
}
