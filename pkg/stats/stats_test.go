package stats

import (
	"github.com/bakhodur-nazriev/bank/v2/pkg/types"
	"reflect"
	"testing"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000},
		{ID: 2, Category: "food", Amount: 2_000},
		{ID: 3, Category: "auto", Amount: 3_000},
		{ID: 4, Category: "food", Amount: 4_000},
		{ID: 5, Category: "food", Amount: 6_000},
	}

	expected := map[types.Category]types.Money{
		"auto": 2_000,
		"food": 4_000,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result){
		t.Errorf("Invalid result, expected: %v, actual: %v", expected, result)
	}
}
