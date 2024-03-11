// codility
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 3, 0, 2}
	ans := Solution(a)
	fmt.Println(ans)
}

func Solution(A []int) int {
	isExist := make(map[int]bool)
	for _, v := range A {
		isExist[v] = true
	}
	if len(isExist) == len(A) {
		for i := 0; i < len(A); i++ {
			if !isExist[i+1] {
				return 0
			}
		}
		return 1
	}
	return 0
}
