package main

func validateDigits(s []string) bool {
	valid := true
	for i := 0; i < 9; i++ { // checking the rows
		for j := 0; j < 9; j++ {
			for k := 0; k < 9; k++ {
				if s[i][j] == s[j][k] {
					valid = false
					return valid
				}
			}
		}
	}
	return valid
}
