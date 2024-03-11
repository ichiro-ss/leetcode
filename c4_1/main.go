// codility
package main

import (
	"fmt"
)

func main() {
	x := 5
	a := []int{1, 3, 1, 4, 2, 3, 5, 4}
	ans := Solution(x, a)
	fmt.Println(ans)
}

func Solution(X int, A []int) int {
	isExist := make(map[int]bool)
	for i, v := range A {
		isExist[v] = true
		if len(isExist) == X {
			return i
		}
	}
	return -1
}
