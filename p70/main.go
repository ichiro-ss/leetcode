package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := getStdin()
	n, _ := strconv.Atoi(lines[0])
	ans := climbStairs(n)
	fmt.Println(ans)
}

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	if len(dp) > 2 {
		dp[2] = 1
	}
	for i := 1; i < n; i++ {
		if i+1 <= n {
			dp[i+1] += dp[i]
		}
		if i+2 <= n {
			dp[i+2] += dp[i]
		}
	}
	return dp[n]
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
