package main

import "fmt"

func main() {
	a := []int{2, 1, 1, 2, 3, 1}
	fmt.Println(Solution(a))
}

func Solution(A []int) int {
	// Implement your solution here
	isExist := make(map[int]bool)
	for _, v := range A {
		isExist[v] = true
	}
	return len(isExist)
}
