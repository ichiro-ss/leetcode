// codility
package main

import (
	"fmt"
)

func main() {
	a := []int{}
	k := 3
	ans := Solution(a, k)
	fmt.Println(ans)
}

func Solution(A []int, K int) []int {
	if len(A) > 0 {
		for i := 0; i < K; i++ {
			A = append(A[len(A)-1:], A[:len(A)-1]...)
		}
	}
	return A
}
