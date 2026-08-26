package conditions

func IsOddNumber(number int) bool {
	if number%2 != 0 {
		return true
	} else {
		return false
	}
}

func IsEvenNumber(number int) bool {
	if number%2 == 0 {
		return true
	} else {
		return false
	}
}

func SortOddNumber(numbers []int) []int {
	var odds []int

	for _, n := range numbers {
		if n%2 != 0 {
			odds = append(odds, n)
		}
	}

	return odds
}

func SortEvenNumber(numbers []int) []int {
	var evens []int

	for _, n := range numbers {
		if n%2 == 0 {
			evens = append(evens, n)
		}
	}

	return evens
}
