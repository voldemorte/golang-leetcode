package medium

func minimumRounds(tasks []int) int {
	var nums = make(map[int]int)

	for _, v := range tasks {
		nums[v]++
	}

	var result int

	for _, v := range nums {
		for v > 0 {
			switch {
			case v%3 == 0:
				result, v = result+v/3, 0
			case v > 4:
				result++
				v -= 3
			case v%2 == 0:
				result, v = result+v/2, 0
			case v == 1:
				return -1
			}
		}
	}
	return result
}
