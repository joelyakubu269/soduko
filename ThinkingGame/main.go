package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var sudoko string
	//var s2 []string
	for i := 1; i <= 9; i++ {
		sudoko += os.Args[i]
		sudoko += "\n"
	}
	sudokoVal := strings.Split(sudoko, "\n")
	buffer, err := findEmptySpace(sudokoVal)
	if err != nil {
		println(err)
	}
	s := FillEmptySpace(buffer)
	valid := validateDigits(s)
	if valid {
		fmt.Println(len(s))

	}
}
