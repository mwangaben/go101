package ranges

func Sums(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func KVs(theMaps map[string]string) ([]string, []string) {
	var keys []string
	var values []string

	for k, v := range theMaps {
		values = append(values, v)
		keys = append(keys, k)
	}
	return keys, values
}
