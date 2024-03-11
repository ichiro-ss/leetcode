// codility
package main

import (
	"fmt"
)

func main() {
	a := []int{9, 3, 9, 3, 9, 7, 9}
	ans := Solution(a)
	fmt.Println(ans)
}

func Solution(A []int) int {
	isPair := make(map[int]bool)
	for _, v := range A {
		if _, ok := isPair[v]; ok {
			delete(isPair, v)
		} else {
			isPair[v] = false
		}
	}
	for k := range isPair {
		return k
	}
	return 0
}
