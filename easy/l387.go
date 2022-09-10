package easy

func firstUniqChar(s string) int {
	var LetterCount = make(map[rune]int)

	for _, j := range s {
		LetterCount[rune(j)]++
	}

	for i, j := range s {
		if LetterCount[rune(j)] == 1 {
			return i
		}
	}

	return -1
}
