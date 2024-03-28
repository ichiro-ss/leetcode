package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	lines := []int{2, 1, 1, 2}
	ans := rob(lines)
	fmt.Println(ans)
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	} else {
		dp := nums[:]
		dp[1] = max(dp[0], dp[1])
		for i := 2; i < len(nums); i++ {
			if dp[i-1] > dp[i-2]+nums[i] {
				dp[i] = dp[i-1]
			} else {
				dp[i] = dp[i-2] + nums[i]
			}
		}
		return dp[len(nums)-1]
	}
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
