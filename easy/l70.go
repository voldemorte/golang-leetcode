package easy

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	var dp = make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
