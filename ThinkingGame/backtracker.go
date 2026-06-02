package main

func solveBoard(board []string) bool {
	rol, col, found := findEmptySpace(board)
	if !found {
		return true // no empty cells left
	}
	for num := byte('1'); num <= '9'; num++ {
		if isValid(board, rol, col, num) {

		}
	}
}
