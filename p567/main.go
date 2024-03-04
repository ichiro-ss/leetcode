package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

func main() {
	ans := checkInclusion("ab", "eidbaooo")
	fmt.Println(ans)
}

func checkInclusion(s1 string, s2 string) bool {
	cnt1 := make(map[byte]int, 0)
	for _, v := range s1 {
		cnt1[byte(v)]++
	}
	l, r := 0, len(s1)-1
	cnt2 := make(map[byte]int, 0)
	for i, v := range s2 {
		if i > r {
			break
		}
		cnt2[byte(v)]++
	}
	// fmt.Println(cnt1, cnt2)
	for r < len(s2) {
		if reflect.DeepEqual(cnt1, cnt2) {
			return true
		}
		if r+1 < len(s2) {
			cnt2[s2[l]]--
			if cnt2[s2[l]] == 0 {
				delete(cnt2, s2[l])
			}
			cnt2[s2[r+1]]++
			// fmt.Println(cnt2)
		}
		l++
		r++
	}
	return false
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
