package arraynslice

func Sum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(numsToSum [][]int) []int {
	lengthOfNumbers := len(numsToSum)
	sums := make([]int, lengthOfNumbers)

	for i, nums := range numsToSum {
		sums[i] = Sum(nums)
	}
	return sums
}
