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
	for _, v := range lines {
		int_v, _ := strconv.Atoi(v)
		nums = append(nums, int_v)
	}
	ans := productExceptSelf(nums)
	fmt.Println(ans)
}

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	pre := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		ans[i] *= pre
		pre *= nums[i]
	}
	return ans
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
