package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FillEmptySpace(board []string) []string {
	s1, err := findEmptySpace(board)

	rows, cols, err := extractPositions(s1)
	if err != nil {
		return board
	}
	for i := 0; i < len(rows); i++ {
		r := rows[i]
		c := cols[i]

		row := []rune(board[r]) // because strings are immutable and so you cannot change the individual values i cast s into a slice of runes
		placed := false
		for num := byte('1'); num <= '9'; num++ {
			if isValid(board, r, c, num) {
				row[c] = rune(num)     // this enables to put a valid value at a position in the slice of runes
				board[r] = string(row) // make changes to the slice of strings at the positions, where dots where found
				placed = true
				break
			}

		}
		if !placed {
			row[c] = '.'
			board[r] = string(row)
		}
	}

	return board
}
func extractPositions(s []string) ([]int, []int, error) {
	rows := []int{}
	cols := []int{}
	for _, r := range s {
		parts := strings.Split(r, ",")
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid format")
		}

		firstNum, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}
		rows = append(rows, firstNum)

		secondNum, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}
		cols = append(cols, secondNum)
	}
	return rows, cols, nil
}
