package functions

func Fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * Fact(n-1)
}

func SumSlice(slices []int) int {
	if len(slices) == 0 {
		return 0
	}
	return slices[0] + SumSlice(slices[1:])
}

func Power(base, exp int) int {
	if exp == 0 {
		return 1
	}
	if exp == 1 {
		return base
	}
	if exp%2 == 0 {
		half := Power(base, exp/2)
		return half * half
	}

	return base * Power(base, exp-1)
}
