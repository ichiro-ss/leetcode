package main

import (
	"container/list"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func main() {
	q := []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"}
	v := [][]int{{}, {-2}, {0}, {-3}, {}, {}, {}, {}}
	minStack(q, v)
}

type MinStack struct {
	st    *list.List
	minSt *list.List
}

func Constructor() MinStack {
	newSt := MinStack{st: list.New(), minSt: list.New()}
	return newSt
}

func (this *MinStack) Push(val int) {
	this.st.PushBack(val)

	e := this.minSt.Back()
	if e != nil {
		this.minSt.PushBack(int(math.Min(float64(val), float64(e.Value.(int)))))
	} else {
		this.minSt.PushBack(val)
	}
}

func (this *MinStack) Pop() {
	s := this.st.Back()
	if s != nil {
		this.st.Remove(s)
	}

	m := this.minSt.Back()
	if m != nil {
		this.minSt.Remove(m)
	}
}

func (this *MinStack) Top() int {
	return this.st.Back().Value.(int)
}

func (this *MinStack) GetMin() int {
	return this.minSt.Back().Value.(int)
}

func minStack(q []string, val [][]int) {
	var st MinStack
	for i, l := range q {
		switch l {
		case "MinStack":
			st = Constructor()
		case "push":
			st.Push(val[i][0])
			fmt.Println(val[i][0], "pushed")
		case "pop":
			st.Pop()
		case "top":
			fmt.Println(st.Top())
		case "getMin":
			fmt.Println(st.GetMin())
		}

	}
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
