package main

import "fmt"

func findEmptySpace(s []string) ([]string, error) {
	buffer := []string{}
	for i := 0; i < 9; i++ { // have a grid  of nine rows
		for j := 0; j < 9; j++ {
			if s[i][j] == '.' { // check if the values of each row per column if its empty
				buffer = append(buffer, fmt.Sprintf("%d,%d", i, j)) // append the position in a slice of strings
			}
		}
	}
	return buffer, nil
}
