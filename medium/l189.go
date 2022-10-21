package medium

func rotate(nums []int, k int) {
	if k == len(nums) {
		return
	}

	var rotate int
	var result = make([]int, len(nums))
	for k > len(nums) {
		k = k - len(nums)
	}

	rotate = k

	copy(result, nums[(len(nums)-rotate):])
	copy(result[rotate:], nums[:(len(nums)-rotate)])
	copy(nums, result)
}
