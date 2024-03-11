// codility
package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{20, -20}
	ans := Solution(a)
	fmt.Println(ans)
}

func Solution(A []int) int {
	sumA := make([]int, len(A))
	if len(A) > 0 {
		sumA[0] = A[0]
	}
	for i := 1; i < len(A); i++ {
		sumA[i] = A[i] + sumA[i-1]
	}
	min := int(math.Abs(float64(sumA[0] - sumA[len(sumA)-1] + sumA[0])))
	for i := 0; i < len(A)-1; i++ {
		diff := int(math.Abs(float64(sumA[i] - sumA[len(sumA)-1] + sumA[i])))
		if diff < min {
			min = diff
		}
	}
	return min
}
