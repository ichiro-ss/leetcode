package main

import "fmt"

func main() {
	a := ")("
	fmt.Println(Solution(a))
}

func Solution(S string) int {
	pairs := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	st := []byte{}
	for _, c := range S {
		if _, ok := pairs[byte(c)]; !ok {
			st = append(st, byte(c))
		} else if len(st) > 0 && st[len(st)-1] == pairs[byte(c)] {
			st = st[:len(st)-1]
		} else {
			return 0
		}
	}
	if len(st) > 0 {
		return 0
	}
	return 1
}
