package main

import (
	"fmt"

	"github.com/gookit/color"
)

func printMatrix(input Input, solution Solution) {
	var result string

	for _, row := range input {
		for _, node := range row {
			char := string(node.Type)

			if solution.contains(node) {
				char = color.FgGreen.Render(char)
			}

			result += char + " "
		}
		result += "\n"
	}

	fmt.Printf("\n%v\n", result)
}
