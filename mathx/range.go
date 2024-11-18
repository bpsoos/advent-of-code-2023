package mathx

import "math"

func Range(start int, length int) []int {
	nums := make([]int, length)
	for i := range length {
		nums[i] = AddInt(start, i)
	}
	return nums
}

func AddInt(a, b int) int {
	if b > 0 {
		if a > math.MaxInt-b {
			println(a, b, math.MaxInt)
			panic("int overflow")
		}
	} else {
		if a < math.MinInt-b {
			println(a, b, math.MaxInt)
			panic("int overflow")
		}
	}

	return a + b
}
