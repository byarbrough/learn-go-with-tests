package slices

// Sum returns the sum of all numbers in slice
func Sum(numbers []int) int {

	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}

// SumAll returns sums of multiple slices
func SumAll(numsToSum ...[]int) []int {

	lenNumsToSum := len(numsToSum)
	sums := make([]int, lenNumsToSum)

	for i, nums := range numsToSum {
		sums[i] = Sum(nums)
	}

	return sums
}
