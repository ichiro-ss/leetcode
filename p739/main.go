package main

import (
	"container/list"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	lines := []int{73, 74, 75, 71, 69, 72, 76, 73}
	ans := dailyTemperatures(lines)
	fmt.Println(ans)
}

func dailyTemperatures(temperatures []int) []int {
	st := list.New()
	id := list.New()
	ans := [100000]int{}
	for i, t := range temperatures {
		if st.Back() == nil {
			st.PushBack(t)
			id.PushBack(i)
		} else {
			for st.Back() != nil && t > st.Back().Value.(int) {
				st.Remove(st.Back())
				ans[id.Back().Value.(int)] = i - id.Remove(id.Back()).(int)
			}
			st.PushBack(t)
			id.PushBack(i)
		}
	}
	return ans[:len(temperatures)]
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
