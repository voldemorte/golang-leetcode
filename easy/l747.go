package easy

// https://leetcode.com/problems/largest-number-at-least-twice-of-others/
func dominantIndex(nums []int) int {
	// check for max and second max and store max index
	var max, second, imax = nums[0], 0, 0

	// loop to find the max and second max in the array
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			second = max
			max, imax = nums[i], i
		} else if nums[i] > second {
			second = nums[i]
		}
	}

	// Check if the condition between max and second max is satisfied
	if 2*second <= max {
		return imax
	}

	return -1
}
