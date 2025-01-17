package main

import (
	"fmt"

	"github.com/MaoDeMatos/pathfinder-algorithm-go/src/matrix"
)

func main() {
	fmt.Printf("Default Matrix: %v\n", matrix.Default.String())
	fmt.Printf("5x5: %v\n", matrix.FiveByFive.String())
	fmt.Printf("4x4: %v\n", matrix.FourByFour.String())
}
