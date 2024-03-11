package main

import "fmt"

func main() {
	a := []int{4, 9, 6, 8}
	b := []int{1, 1, 1, 0}

	fmt.Println(Solution(a, b))
}

func Solution(A []int, B []int) int {
	river := [][2]int{}
	for i := 0; i < len(A); i++ {
		if B[i] == 0 {
			for len(river) >= 0 {
				if len(river) == 0 {
					river = append(river, [2]int{A[i], B[i]})
					break
				}
				if river[len(river)-1][1] == 0 { // upstream
					river = append(river, [2]int{A[i], B[i]})
					break
				} else { // downstream
					if river[len(river)-1][0] < A[i] { // poped elem is smaller than [i]
						river = river[:len(river)-1]
					} else {
						break
					}
				}
			}
		} else { // B[i] == 1
			river = append(river, [2]int{A[i], B[i]})
		}
	}
	return len(river)
}
