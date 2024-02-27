package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

func main() {
	lines := getStdin()

	ans := isAnagram(lines[0], lines[1])
	fmt.Println(ans)
}

func isAnagram(s string, t string) bool {
	s_elem := map[byte]int{}
	t_elem := map[byte]int{}

	for i := 0; i < len(s); i++ {
		s_elem[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		t_elem[t[i]]++
	}
	if reflect.DeepEqual(s_elem, t_elem) {
		return true
	}

	return false
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
