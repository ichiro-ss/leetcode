package main

import "fmt"

func main() {
	a := []int{4, 3, 4, 4, 4, 2}

	fmt.Println(Solution(a))
}

func Solution(A []int) int {
	idx := make(map[int][]int)
	half := len(A) / 2
	leader := 0
	for i, v := range A {
		idx[v] = append(idx[v], i)
		if half < len(idx[v]) {
			leader = v
		}
	}
	idxLeader := 0
	ans := 0
	for i := 0; i < len(A); i++ {
		if idxLeader+1 < len(idx[leader]) && i == idx[leader][idxLeader+1] {
			idxLeader++
		}
		// fmt.Println(i, ":", idxLeader+1, ",", len(idx[leader])-idxLeader-1, "-", idxLeader)
		if idxLeader+1 > (i+1)/2 && len(idx[leader])-idxLeader-1 > (len(A)-i-1)/2 {
			ans++
		}
	}
	return ans
}
