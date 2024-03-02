package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	lines := []int{2, 7, 11, 15}
	t := 9
	ans := twoSum(lines, t)
	fmt.Println(ans)
}

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for {
		if numbers[i]+numbers[j] < target {
			i++
			continue
		}
		if numbers[i]+numbers[j] > target {
			j--
			continue
		}
		break
	}
	return []int{i + 1, j + 1}
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
