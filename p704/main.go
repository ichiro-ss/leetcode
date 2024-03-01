package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	lines := []int{-1, 0, 3, 5, 9, 12}
	t := 9
	ans := search(lines, t)
	fmt.Println(ans)
}

func search(nums []int, target int) int {
	i := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
	if i < len(nums) && nums[i] == target {
		return i
	} else {
		return -1
	}
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
