package easy

func arrangeCoins(n int) int {
	var rows int

	for n > 0 {
		rows++
		n -= rows
	}

	if n == 0 {
		return rows
	}

	return rows - 1
}
