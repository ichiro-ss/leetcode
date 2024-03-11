// codility
package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{0, 1, 0, 1, 1, 0}
	ans := Solution(a)
	fmt.Println(ans)
}

func Solution(A []int) int {
	zeros := 0
	ans := 0
	for _, v := range A {
		if v == 0 {
			zeros++
		} else {
			ans += zeros
		}
		if ans > int(math.Pow(10, 9)) {
			return -1
		}
	}
	return ans
}
