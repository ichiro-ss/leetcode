package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	lines := getStdin()
	ans := groupAnagrams(lines)
	fmt.Println(ans)
}

func groupAnagrams(strs []string) [][]string {
	ga := map[[26]int][]string{}
	for _, v := range strs {
		ana := [26]int{0}
		for _, w := range v {
			ana[int(w)-int('a')]++
		}
		ga[ana] = append(ga[ana], v)
	}
	var ans [][]string
	for _, v := range ga {
		ans = append(ans, v)
	}
	return ans
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
