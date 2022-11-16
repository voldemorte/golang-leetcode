package medium

func myAtoi(s string) int {
	var result int
	var sign = 1
	var i int
	intMax, intMin := (1<<31)-1, 1<<31
	var length = len(s)

	// Skip initial white space
	for i < length && s[i] == ' ' {
		i++
	}

	// Check for sign
	if i < length && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// Check for digits
	for i < length && (s[i] >= '0' && s[i] <= '9') {
		digit := int(s[i] - '0')

		// Check for constraints
		if (sign > 0 && result > (intMax-digit)/10) || (sign < 0 && result > (intMin-digit)/10) {
			if sign == 1 {
				return intMax
			} else {
				return -intMin
			}
		}

		// Add the digit to the result
		result = (result * 10) + digit
		i++
	}

	return result * sign
}
