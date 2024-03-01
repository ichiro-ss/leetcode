package main

import (
	"container/list"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := []string{"2", "1", "+", "3", "*"}
	ans := evalRPN(lines)
	fmt.Println(ans)
}

func evalRPN(tokens []string) int {
	st := list.New()
	for _, t := range tokens {
		switch t {
		case "+":
			st.PushBack(st.Remove(st.Back()).(int) + st.Remove(st.Back()).(int))
		case "-":
			u, b := st.Remove(st.Back()).(int), st.Remove(st.Back()).(int)
			st.PushBack(b - u)
		case "*":
			st.PushBack(st.Remove(st.Back()).(int) * st.Remove(st.Back()).(int))
		case "/":
			u, b := st.Remove(st.Back()).(int), st.Remove(st.Back()).(int)
			st.PushBack(b / u)
		default:
			v, _ := strconv.Atoi(t)
			st.PushBack(v)
		}
	}
	return st.Back().Value.(int)
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
