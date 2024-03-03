package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func main() {
	lines := "abba"
	ans := lengthOfLongestSubstring(lines)
	fmt.Println(ans)
}

func lengthOfLongestSubstring(s string) int {
	places := make(map[byte]int, 0)
	l, r := 0, 0
	ans := 0
	for r < len(s) {
		if p, ok := places[s[r]]; ok {
			l = int(math.Max(float64(l), float64(p+1)))
		}
		ans = int(math.Max(float64(ans), float64(r-l+1)))
		places[s[r]] = r
		r++
	}
	return ans
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
