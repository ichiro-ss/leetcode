package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	lines := getStdin()
	var board [][]byte
	for i, v := range lines {
		for _, w := range v {
			board[i] = append(board[i], byte(w))
		}
	}
	ans := isValidSudoku(board)
	fmt.Println(ans)
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		row := make(map[byte]int)
		col := make(map[byte]int)

		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				col[board[i][j]]++
				if col[board[i][j]] == 2 {
					return false
				}
			}
			if board[j][i] != '.' {
				row[board[j][i]]++
				if row[board[j][i]] == 2 {
					return false
				}
			}
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			square := make(map[byte]int)
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if board[i+k][j+l] != '.' {
						square[board[i+k][j+l]]++
						if square[board[i+k][j+l]] == 2 {
							return false
						}
					}
				}
			}
		}
	}
	return true
}

func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
