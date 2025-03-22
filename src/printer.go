package main

import (
	"fmt"

	"github.com/gookit/color"
)

func printMatrix(input Input, solution Solution) {
	var result string

	for rowIndex := range input {
		for _, node := range input[rowIndex] {
			char := ""

			switch node.Type {
			case StartNode:
				char = "S "
			case EndNode:
				char = "E "
			case WallNode:
				char = "X "
			case NormalNode:
				char = ". "
			}

			if solution.contains(node) {
				char = color.FgGreen.Render(char)
			}

			result += char
		}
		result += "\n"
	}

	fmt.Printf("\n%v\n", result)
}
