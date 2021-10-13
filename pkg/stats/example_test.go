package stats

import (
	"fmt"

	"github.com/bakhodur-nazriev/bank/v2/pkg/types"
)

func ExampleAvg() {
	myPayment := []types.Payment{
		{
			ID:     1,
			Amount: 5,
			Status: types.StatusFail,
		},
		{
			ID:     2,
			Amount: 10,
			Status: types.StatusOk,
		},
		{
			ID:     3,
			Amount: 10,
			Status: types.StatusOk,
		},
		{
			ID:     4,
			Amount: 10,
			Status: types.StatusOk,
		},
	}

	fmt.Println(Avg(myPayment))

	// Output:
	// 10
}

func ExampleTotalInCategory() {
	myPayment := []types.Payment{
		{
			ID:       1,
			Amount:   5,
			Category: "test",
			Status:   types.StatusFail,
		},
		{
			ID:       2,
			Amount:   10,
			Category: "test",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   10,
			Category: "test",
			Status:   types.StatusOk,
		},
		{
			ID:       4,
			Amount:   10,
			Category: "test",
			Status:   types.StatusOk,
		},
	}

	fmt.Println(TotalInCategory(myPayment, "test"))

	// Output:
	// 30
}
