package easy

func containsDuplicate(nums []int) bool {
	var count = make(map[int]int)

	for _, v := range nums {
		// Check if number already exists in the map
		_, exists := count[v]
		if exists { // If exits then we have a duplicate and return
			return true
		}

		// Else add the number to map.
		count[v]++
	}

	// Reached the end of array so no duplicates found.
	return false
}
