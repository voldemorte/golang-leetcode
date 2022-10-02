package medium

func majorityElement(nums []int) []int {
	var count = make(map[int]int)
	var result []int

	for _, v := range nums {
		count[v]++
	}

	for k := range count {
		if count[k] > (len(nums) / 3) {
			result = append(result, k)
		}
	}

	return result
}
