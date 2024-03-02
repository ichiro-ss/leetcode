package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	lines := "A man, a plan, a canal: Panama"
	ans := isPalindrome(lines)
	fmt.Println(ans)
}

func isPalindrome(s string) bool {
	tmp := []byte{}
	for _, c := range s {
		if (int('0') <= int(c) && int(c) <= int('9')) || (int('A') <= int(c) && int(c) <= int('Z')) || (int('a') <= int(c) && int(c) <= int('z')) {
			tmp = append(tmp, byte(c))
		}
	}
	pal := string(tmp)
	fmt.Println(pal)
	low_pal := strings.ToLower(pal)
	for i := 0; i < len(low_pal)/2; i++ {
		if low_pal[i] != low_pal[len(low_pal)-i-1] {
			return false
		}
	}
	return true
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
