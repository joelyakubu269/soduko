package main

func solveBoard(board []string) bool {
	row, col, found := findEmptySpace(board)
	if !found {
		return true // no empty cells left
	}
	for num := byte('1'); num <= '9'; num++ {
		if isValid(board, row, col, num) {
			rowRunes:= []rune(board[row]) // since strings are immutable in a way that we cant change individual characters
			// i create a slice of runes
			rowRunes[col]= rune(num)
			board[row]= string(rowRunes)
			if solveBoard(board) { // if all cells have been filled
			return true
		}
		// undo backtrack
		rowRunes[col]= '.'
		board[row]= string(rowRunes)
		}
		
	}
	return false
}
