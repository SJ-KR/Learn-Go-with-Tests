package arrays

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
func SumAll(numbersToSum ...[]int) []int {
	length := len(numbersToSum)
	sums := make([]int, length)

	for i := 0; i < length; i++ {
		sums[i] = Sum(numbersToSum[i])
	}
	return sums
}
