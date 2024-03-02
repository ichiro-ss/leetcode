package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func main() {
	lines := []int{3, 6, 7, 11}
	t := 8
	ans := minEatingSpeed(lines, t)
	fmt.Println(ans)
}

func minEatingSpeed(piles []int, h int) int {
	l := 1
	r := int(math.Pow(10, 9))
	ans := 1
	for l <= r {
		m := (r + l) / 2
		cnt := 0
		for _, v := range piles {
			cnt += (v + m - 1) / m
		}
		if cnt > h {
			l = m + 1
		} else {
			ans = m
			r = m - 1
		}
	}
	return ans
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
