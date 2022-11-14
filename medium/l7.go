package easy

func reverse(x int) int {
	var result int
	in := x
	for {
		if in == 0 {
			break
		}
		result = (result * 10) + (in % 10)
		if result < -2147483648 || result > 2147483647 {
			return 0
		}
		in /= 10
	}

	return result
}
