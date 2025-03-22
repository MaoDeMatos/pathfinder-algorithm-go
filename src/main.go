package main

import (
	"fmt"

	"github.com/gookit/color"
)

var mainInput = fiveByFiveInput

func main() {
	solution, err := BreadthFirstSearch(mainInput)

	if err != nil {
		fmt.Println()
		color.Red.Println(err.Error())
	}

	printMatrix(mainInput, solution)
	// fmt.Println()
}
