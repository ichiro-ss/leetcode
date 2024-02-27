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
	ans := containsDuplicate(nums)
	fmt.Println(ans)
}

func containsDuplicate(nums []int) bool {
	is_exist := map[int]bool{}

	for _, v := range nums {
		if _, ok := is_exist[v]; ok {
			return true
		}
		is_exist[v] = true
	}
	return false
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
