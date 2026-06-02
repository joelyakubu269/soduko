package main

import (
	"strconv"
	"strings"
)

func FillEmptySpace(s []string) *[]string {

}
func extractPositions(s []string) ([]int, []int, error) {
	rows := []int{}
	cols := []int{}
	for i, r := range s {
		parts := strings.Split(string(r[i]), ",")

		firstNum, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}
		rows = append(rows, firstNum)

		secondNum, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}
		cols = append(cols, secondNum)
	}
	return rows, cols, nil
}
