package easy

//https://leetcode.com/problems/sqrtx/description

func mySqrt(x int) int {
	l, r := 0, x
	res := 0
	m := 0

	for l <= r {
		m = l + ((r - l) / 2)

		if (m * m) > x {
			r = m - 1
		} else if (m * m) < x {
			l = m + 1
			res = m
		} else {
			return m
		}
	}

	return res
}
