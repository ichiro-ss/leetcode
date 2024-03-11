package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	a := []int{2, 1, 1, 2, 3, 1}
	fmt.Println(Solution(a))
}

func Solution(A []int) int {
	sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })
	return int(math.Max(float64(A[0]*A[1]*A[len(A)-1]), float64(A[len(A)-3]*A[len(A)-2]*A[len(A)-1])))
}
