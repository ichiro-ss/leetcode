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
	var nums []int
	var target int
	for i, v := range lines {
		if i == len(lines)-1 {
			target, _ = strconv.Atoi(lines[i])
			break
		}
		int_v, _ := strconv.Atoi(v)
		nums = append(nums, int_v)
	}
	ans := twoSum(nums, target)
	fmt.Println(ans)
}

func twoSum(nums []int, target int) []int {
	var ans = []int{0, 1}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ans[0] = i
				ans[1] = j
				return ans
			}
		}
	}
	return ans
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
