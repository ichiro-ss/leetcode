package main

import (
	"container/list"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	lines := getStdin()

	ans := isValid(lines[0])
	fmt.Println(ans)
}

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{v: list.New()}
}

func (s *Stack) Push(v interface{}) {
	s.v.PushBack(v)
}

func (s *Stack) Pop() interface{} {
	b := s.v.Back()
	if b == nil {
		return nil
	}
	return s.v.Remove(b)
}
func isValid(s string) bool {
	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	st := NewStack()
	for _, v := range s {
		pair, ok := pairs[byte(v)]
		if !ok {
			st.Push(byte(v))
			continue
		}
		if st.v.Len() == 0 {
			return false
		}
		if st.Pop() != pair {
			return false
		}
	}
	return st.v.Len() == 0
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
