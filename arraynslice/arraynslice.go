package arraynslice

func Sum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(numsToSum [][]int) []int {
	var sums []int
	for _, nums := range numsToSum {
		sums = append(sums, Sum(nums))
	}
	return sums
}

func SumAllTails(numsToSum [][]int) []int {
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
