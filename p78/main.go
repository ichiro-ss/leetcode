package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	lines := []int{1, 2, 3}
	ans := subsets(lines)
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

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
