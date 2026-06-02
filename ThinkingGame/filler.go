package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FillEmptySpace(s []string) *[]string {
	rows, cols, err := extractPositions(s)
	if err == nil {
		for i := 0; i < len(rows); i++ {
			iVal := rows[i]
			jVal := cols[i]
			row := []rune(s[iVal])
			for j := 0; j < 9; j++ {
				row[jVal] = rune('0' + j)
				s[iVal] = string(row)
				valid := validateDigits(s)
				if !valid {

				}
			}
		}
	}

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
