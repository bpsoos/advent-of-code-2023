package mathx

func Multiply(nums []int) int {
	result := 1

	for i := 0; i < len(nums); i++ {
		result = result * nums[i]
	}

	return result
}
