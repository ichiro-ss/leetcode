package main

import "fmt"

func main() {
	s := "(())(())())()"
	fmt.Println(Solution(s))
}

func Solution(S string) int {
	cnt := 0
	for _, c := range S {
		switch byte(c) {
		case '(':
			cnt++
		case ')':
			cnt--
		}
		if cnt < 0 {
			return 0
		}
	}
	if cnt != 0 {
		return 0
	}
	return 1
}
