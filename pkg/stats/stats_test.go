package stats

import (
	"reflect"
	"testing"

	"github.com/bakhodur-nazriev/bank/v2/pkg/types"
)

func TestCategoriesTotal(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000_00},
		{ID: 2, Category: "food", Amount: 2_000_00},
		{ID: 3, Category: "auto", Amount: 3_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 5, Category: "fun", Amount: 5_000_00},
	}

	expected := map[types.Category]types.Money{
		"auto": 8_000_00,
		"food": 2_000_00,
		"fun":  5_000_00,
	}

	result := CategoriesTotal(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestFilteredByCategory_nil(t *testing.T) {
	var payments []types.Payment
	result := FilteredByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilteredByCategory_empty(t *testing.T) {
	var payments []types.Payment
	result := FilteredByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilteredByCategory_notFound(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}

	result := FilteredByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilteredByCategory_foundOne(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}

	expected := []types.Payment{
		{ID: 2, Category: "food"},
	}

	result := FilteredByCategory(payments, "food")

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestFilteredByCategory_foundMultiple(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}

	expected := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
	}

	result := FilteredByCategory(payments, "food")

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid result, expected: %v, actual: %v", expected, result)
	}
}

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

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid result, expected: %v, actual: %v", expected, result)
	}
}
