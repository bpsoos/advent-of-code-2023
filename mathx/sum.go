package mathx

func Sum(nums []int) int {
	result := 0

	for _, num := range nums {
		result = AddInt(result, num)
	}

	return result
}
