package easy

func majorityElement(nums []int) int {
	var count = make(map[int]int)
	var result = 0
	for _, v := range nums {
		count[v]++

		if count[v] > (len(nums) / 2) {
			result = v
		}
	}

	return result
}
