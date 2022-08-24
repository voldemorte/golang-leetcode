package easy

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	uniq := make(map[int]int)

	var i, k = 0, len(nums)

	for i < k {
		if uniq[nums[i]] > 0 {
			copy(nums[i:], nums[(i+1):])
			k--
		} else {
			uniq[nums[i]]++
			i++
		}
	}

	return k
}
