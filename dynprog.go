package main

import "fmt"

func solveDP(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-1]+nums[i])
	}

	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 3, -2, 5}
	fmt.Println(solveDP(nums))
}
