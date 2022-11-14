package medium

func rotate(matrix [][]int) {

	// Transpose the matrix
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}

	// Reverse each row
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[i][len(matrix)-1-j]
			matrix[i][len(matrix)-1-j] = temp
		}
	}
}
