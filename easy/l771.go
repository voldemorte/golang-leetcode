package easy

func numJewelsInStones(jewels string, stones string) int {
	var result int
	for i := 0; i < len(jewels); i++ {
		for j := 0; j < len(stones); j++ {
			if jewels[i] == stones[j] {
				result++
			}
		}
	}

	return result
}
