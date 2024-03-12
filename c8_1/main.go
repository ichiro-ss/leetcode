package main

import "fmt"

func main() {
	a := []int{3, 4, 3, 2, 3, -1, 3, 3}

	fmt.Println(Solution(a))
}

func Solution(A []int) int {
	cnt := make(map[int]int)
	half := len(A) / 2
	for i, v := range A {
		cnt[v]++
		if half < cnt[v] {
			return i
		}
	}
	return -1
}
