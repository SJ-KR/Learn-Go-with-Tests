package arrays

func Sum(nums [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += nums[i]
	}
	return sum
}
