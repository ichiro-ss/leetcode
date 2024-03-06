package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	lines := []int{1, 2, 3}
	// ans := subsets(lines)
	ans := subsets2(lines)
	fmt.Println(ans)
}

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	for b := 0; b < (1 << uint(len(nums))); b++ {
		tmp := make([]int, 0)
		for i := 0; i < len(nums); i++ {
			if (b>>uint(i))&1 == 1 {
				tmp = append(tmp, nums[i])
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}

func subsets2(nums []int) [][]int {
	ans := make([][]int, 0)
	cur := make([]int, 0)
	var dfs func(i int)
	dfs = func(i int) {
		if i >= len(nums) {
			ans = append(ans, append([]int{}, cur...))
			fmt.Println(ans)
			return
		}
		cur = append(cur, nums[i])
		dfs(i + 1)
		cur = cur[:len(cur)-1]
		dfs(i + 1)
	}
	dfs(0)
	return ans
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
