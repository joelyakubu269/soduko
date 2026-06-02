package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FillEmptySpace(s []string) []string {
	s1, err := findEmptySpace(s)

	rows, cols, err := extractPositions(s1)
	if err != nil {
		return s
	}
	for i := 0; i < len(rows); i++ {
		iVal := rows[i]
		jVal := cols[i]
		found := false
		row := []rune(s[iVal]) // because strings are immutable and so you cannot change the individual values i cast s into a slice of runes
		for j := 0; j < 9; j++ {
			row[jVal] = rune('0' + j)
			s[iVal] = string(row)
			if validateDigits(s) {
				found = true
				break
			}

		}
		if !found {
			row[jVal] = '.'
			s[iVal] = string(row)
		}
	}

	return s
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
