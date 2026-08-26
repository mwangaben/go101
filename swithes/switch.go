package swithes

import (
	"fmt"
	"go101/conditions"
)

func SortNumbers(numbers []int, mode string) []int {
	switch mode {
	case "even":
		return conditions.SortEvenNumber(numbers)
	case "odd":
		return conditions.SortOddNumber(numbers)
	default:
		fmt.Printf("failed to Sort %v", numbers)
		var ints []int
		return ints
	}
}

func GiveType(t any) string {
	switch t.(type) {
	case bool:
		return "bool"
	case string:
		return "string"
	case rune:
		return "rune"
	case int:
		return "int"
	default:
		return "unknow type"

	}
}
