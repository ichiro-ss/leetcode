package main

import (
	"fmt"
)

func main() {
	cost := []int{10, 15, 20}
	ans := minCostClimbingStairs(cost)
	fmt.Println(ans)
}

func minCostClimbingStairs(cost []int) int {
	dp := cost[:]
	dp = append(dp, 0)
	for i := 2; i <= len(cost); i++ {
		dp[i] += min(dp[i-1], dp[i-2])
	}
	return dp[len(cost)]
}
