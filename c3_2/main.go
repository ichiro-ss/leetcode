// codility
package main

import (
	"fmt"
)

func main() {
	a := []int{2, 3, 1, 5}
	ans := Solution(a)
	fmt.Println(ans)
}

func Solution(A []int) int {
	isExist := make([]bool, len(A)+1)
	for _, v := range A {
		isExist[v-1] = true
	}
	for i, v := range isExist {
		if !v {
			return i + 1
		}
	}
	return 0
}
