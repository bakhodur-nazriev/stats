package stats

import (
	"fmt"

	"github.com/bakhodur-nazriev/bank/pkg/types"
)

func ExampleAvg() {
	myPayment := []types.Payment{
		{
			ID:     1,
			Amount: 5,
		},
		{
			ID:     2,
			Amount: 10,
		},
		{
			ID:     3,
			Amount: 10,
		},
		{
			ID:     4,
			Amount: 10,
		},
	}

	fmt.Println(Avg(myPayment))

	// Output:
	// 8
}

func ExampleTotalInCategory() {
	myPayment := []types.Payment{
		{
			ID:       1,
			Amount:   5,
			Category: "test",
		},
		{
			ID:       2,
			Amount:   10,
			Category: "test",
		},
		{
			ID:       3,
			Amount:   10,
			Category: "test",
		},
		{
			ID:       4,
			Amount:   10,
			Category: "lol",
		},
	}

	fmt.Println(TotalInCategory(myPayment, "test"))

	// Output:
	// 25
}
