package main

func findEmptySpace(s []string) (int, int, bool) {

	for i := 0; i < 9; i++ { // have a grid  of nine rows
		for j := 0; j < 9; j++ {
			if s[i][j] == '.' { // check if the values of each row per column if its empty
				return i, j, true
			}
		}
	}
	return 0, 0, false
}
