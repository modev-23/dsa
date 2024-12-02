package problems

import "math"

func findClosestNumber(nums []int) int {
	var min int = nums[0]
	for _, v := range nums {
		if math.Abs(float64(v)) == math.Abs(float64(min)) {
			min = int(math.Max(float64(v), float64(min)))
		}
		if math.Abs(float64(v)) < math.Abs(float64(min)) {
			min = v
		}
	}
	return min
}
