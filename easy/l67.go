package easy

//https://leetcode.com/problems/add-binary/
func addBinary(a string, b string) string {
	var maxLen int
	var carry rune

	aRune := []rune(a)
	bRune := []rune(b)

	// Find the largest string
	// Make the lengths of strings equal by prepending 0
	if len(a) > len(b) {
		maxLen = len(aRune)
		for i := 0; i < len(a)-len(b); i++ {
			bRune = append([]rune{'0'}, bRune...)
		}
	} else {
		maxLen = len(b)
		for i := 0; i < len(b)-len(a); i++ {
			aRune = append([]rune{'0'}, aRune...)
		}
	}
	carry = rune('0')

	result := make([]rune, 0, maxLen+1)

	for i := maxLen - 1; i >= 0; i-- {
		if carry == rune('1') && aRune[i] == rune('1') && bRune[i] == rune('1') {
			result = append([]rune{'1'}, result...)

		} else if carry == rune('0') && (aRune[i] == rune('1') && bRune[i] == rune('1')) {
			carry = rune('1')
			result = append([]rune{'0'}, result...)

		} else if carry == rune('1') && aRune[i] == rune('0') && bRune[i] == rune('0') {
			result = append([]rune{'1'}, result...)
			carry = rune('0')

		} else if carry == rune('1') && ((aRune[i] == rune('1') && bRune[i] == rune('0')) || (aRune[i] == rune('0') && bRune[i] == rune('1'))) {
			result = append([]rune{'0'}, result...)

		} else if carry == rune('0') && ((aRune[i] == rune('1') && bRune[i] == rune('0')) || (aRune[i] == rune('0') && bRune[i] == rune('1'))) {
			result = append([]rune{'1'}, result...)
		} else if carry == rune('0') && aRune[i] == rune('0') && bRune[i] == rune('0') {
			result = append([]rune{'0'}, result...)
		}
	}
	if carry == rune('1') {
		result = append([]rune{'1'}, result...)
	}

	return string(result)
}
