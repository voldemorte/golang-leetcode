package medium

import "fmt"

func isValidSudoku(board [][]byte) bool {
	var found = make(map[string]bool)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				// Generate map key strings
				tempRow := fmt.Sprintf("%s is present in row %d", string(board[i][j]), i)
				tempCol := fmt.Sprintf("%s is present in col %d", string(board[i][j]), j)
				tempSubBox := fmt.Sprintf("%s is present in subBox %d, %d", string(board[i][j]), i/3, j/3)

				// Check if generated key is already present in map
				if _, present := found[tempRow]; present {
					return false
				}
				if _, present := found[tempCol]; present {
					return false
				}
				if _, present := found[tempSubBox]; present {
					return false
				}

				// If not present add it to the map
				found[tempRow] = true
				found[tempCol] = true
				found[tempSubBox] = true
			}
		}
	}

	return true
}
