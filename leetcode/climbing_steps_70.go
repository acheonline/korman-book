package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
}

func climbStairs(n int) int {
	if n < 3 { // O(0)
		return n
	}

	dp := make([]int, n+1)    // O(0)
	dp[1] = 1                 // O(0)
	dp[2] = 2                 // O(0)
	for i := 3; i <= n; i++ { // O(n)
		dp[i] = dp[i-1] + dp[i-2] //O(0)
	}
	return dp[n] //O(0)
}

// whole time spent is O(0) + O(0) + O(0) + O(0) + N * (O(0)) + O(0)
// simple record is - c*O(0) + O(n)
// more simple - O(n)
