package easy

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	var fib = make([]int, n+1)

	fib[0] = 0
	fib[1] = 1

	for i := 2; i < n+1; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[n]
}
