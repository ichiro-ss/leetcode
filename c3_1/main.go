// codility
package main

import (
	"fmt"
)

func main() {
	x, y, d := 10, 85, 30
	ans := Solution(x, y, d)
	fmt.Println(ans)
}

func Solution(X, Y, D int) int {
	return (Y - X + D - 1) / D
}
