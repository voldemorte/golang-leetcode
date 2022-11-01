package easy

func intersect(nums1 []int, nums2 []int) []int {
	var ans []int
	var intersection = make(map[int]int)

	for _, v := range nums1 {
		intersection[v]++
	}

	for _, v := range nums2 {
		_, exists := intersection[v]

		if exists && (intersection[v] > 0) {
			ans = append(ans, v)
			intersection[v]--
		}
	}

	return ans
}
