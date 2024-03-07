// codility_
package main

import (
	"fmt"
	"math"
)

func main() {
	n := 32
	ans := Solution(n)
	fmt.Println(ans)
}

func Solution(N int) int {
	max, cur := 0, 0
	dig := 0.0
	for {
		if N <= int(math.Pow(2.0, dig)-1.0) {
			break
		}
		dig++
	}
	for i := int(dig) - 1; i >= 0; i-- {
		if N>>uint(i)&1 == 1 {
			max = int(math.Max(float64(max), float64(cur)))
			cur = 0
		} else {
			cur++
		}
	}
	return max
}
