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
	var k int
	for i, v := range lines {
		if i == len(lines)-1 {
			k, _ = strconv.Atoi(lines[i])
			break
		}
		int_v, _ := strconv.Atoi(v)
		nums = append(nums, int_v)
	}
	ans := topKFrequent(nums, k)
	fmt.Println(ans)
}

func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}
	for _, v := range nums {
		count[v]++
	}
	bucket := make([][]int, len(nums)+1)
	for k, v := range count {
		bucket[v] = append(bucket[v], k)
	}

	ans := []int{}
	for i := len(bucket) - 1; i >= 0; i-- {
		for _, v := range bucket[i] {
			ans = append(ans, v)
			k--
		}
		if k == 0 {
			break
		}
	}
	return ans
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
