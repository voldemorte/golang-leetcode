package easy

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		return true
	}

	search := []rune(s)

	for i := 0; i < len(s); i++ {
		temp := search[0]
		copy(search, search[1:])
		search[len(s)-1] = temp
		if string(search) == goal {
			return true
		}
	}
	return false
}
