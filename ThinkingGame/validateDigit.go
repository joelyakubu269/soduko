package main

func isValid(board []string, row, col int, val byte) bool {
	// row check
	for i := 0; i < 9; i++ {
		if board[row][i] == val { // checks a particular row from column 0 to 8
			return false
		}
	}
	// column check
	for j := 0; j < 9; j++ {
		if board[j][col] == val { // checks a particular column from rows 0 to 8
			return false
		}
	}

	// box check
	boxrow := (row / 3) * 3
	boxcol := (col / 3) * 3
	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			if board[boxrow+k][boxcol+k] == val {
				return false
			}
		}

	}
	return true
}

// func validateDigits(s []string) bool {
// 	valid := true
// 	for i := 0; i < 9; i++ { // checking the rows
// 		for j := 0; j < 9; j++ {
// 			for k := j+i; k < 9; k++ {
// 				if s[i][j] == s[i][k] && s[i][j]== '.' || s[j][i] == s[j][k] && s[j][i] == '.' { // making sure duplicates are not dots
// 					continue
// 				}
// 				if s[i][j] == s[i][k] || s[j][i] == s[j][k] {
// 					valid = false
// 					return valid
// 				}
// 			}
// 		}
// 	}
// 	return valid
// }
