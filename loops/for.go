package loops

import "fmt"

func LoopOne() []int {
	var result []int

	i := 1
	for i <= 4 {
		fmt.Println(i)
		result = append(result, i)
		i++
	}

	return result
}

func LoopTwo() []int {
	var result []int
	for j := 1; j <= 4; j++ {
		result = append(result, j)
	}
	return result
}

func LoopThree() []int {
	var result []int
	for n := range 4 {
		result = append(result, n)
	}

	return result
}

func ShowAlphabet(word string) []rune {
	var collection []rune
	for _, char := range word {
		collection = append(collection, char)
	}
	return collection
}
