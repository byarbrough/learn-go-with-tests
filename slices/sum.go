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
	var sums []int
	for _, nums := range numsToSum {
		sums = append(sums, Sum(nums))
	}

	return sums
}

// SumAllTails returns sum of end of slice
func SumAllTails(numsToSum ...[]int) []int {
	var sums []int
	for _, nums := range numsToSum {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {

			tail := nums[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
