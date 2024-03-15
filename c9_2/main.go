package main

import "fmt"

func main() {
	a := []int{-3, -2, -1}

	fmt.Println(Solution(a))
}

func Solution(A []int) int {
	max := A[0]
	cur := A[0]
	for i := 1; i < len(A); i++ {
		fmt.Println("cur:", cur, "max:", max)
		if cur < 0 {
			cur = 0
		}
		if max < cur+A[i] {
			max = cur + A[i]
		}
		cur += A[i]
	}
	return max
}
