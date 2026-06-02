package main

import (
	"os"
	"strings"
)

func main() {
	var sudoko string
	for i := 1; i <= 9; i++ {
		sudoko += os.Args[i]
		sudoko += "\n"
	}
	sudokoVal := strings.Split(sudoko, "\n")
	buffer, err := findEmptySpace(sudokoVal)
	if err == nil {

	}
}
