package easy

func isPalindrome(s string) bool {
	var Letters []rune

	for _, v := range s {
		if (v >= '0' && v <= '9') || (v >= 'a' && v <= 'z') {
			Letters = append(Letters, v)
		} else if v >= 'A' && v <= 'Z' {
			Letters = append(Letters, (v + 32))
		}
	}

	for i, j := 0, (len(Letters) - 1); i <= j; i, j = i+1, j-1 {
		if Letters[i] != Letters[j] {
			return false
		}
	}

	return true
}
