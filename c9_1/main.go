package main

import "fmt"

func main() {
	a := []int{3, 4, 3, 2, 3, -1, 3, 3}

	fmt.Println(Solution(a))
}

func Solution(A []int) int {
	if len(A) > 0 {
		min := A[0]
		max := 0
		for i := 1; i < len(A); i++ {
			if min > A[i] {
				min = A[i]
			} else if min < A[i] && max < A[i]-min {
				max = A[i] - min
			}
		}
		return max
	}
	return 0
}
