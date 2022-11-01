package easy

func plusOne(digits []int) []int {
	var carry int

	carry = (digits[len(digits)-1] + 1) / 10
	digits[len(digits)-1] = (digits[len(digits)-1] + 1) % 10

	for i := len(digits) - 2; i >= 0; i-- {
		temp := digits[i] + carry
		digits[i] = (digits[i] + carry) % 10
		carry = temp / 10
	}

	if carry == 0 {
		return digits
	}
	digits = append(digits, 0)
	copy(digits[1:], digits[0:])
	digits[0] = carry
	return digits
}
