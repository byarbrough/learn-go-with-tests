package slices

// Sum returns the sum of all numbers in slice
func Sum(numbers []int) int {

	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}
