package medium

import "math"

func reorderedPowerOf2(n int) bool {
	var pow2digits = [][]int{}

	for i := 0; i <= 29; i++ { // (2**29 is the closest pow 2 less than 10**9, as in exercise constraints)
		pow2value := math.Pow(2, float64(i))
		pow2digits = append(pow2digits, createDigitsArray(int(pow2value)))
	}

	inputDigits := createDigitsArray(n)

	// compare the generated digits with input digits
	for _, pow2digit := range pow2digits {
		if compareDigits(pow2digit, inputDigits) {
			return true
		}
	}

	return false
}

func compareDigits(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func createDigitsArray(n int) []int {
	digits := make([]int, 10)
	for n > 0 {
		digits[n%10]++
		n /= 10
	}
	return digits
}
