package main

import (
	"fmt"

	"github.com/MaoDeMatos/pathfinder-algorithm-go/src/matrix"
)

func main() {
	fmt.Printf("Default Matrix: %s\n", matrix.Simple)
	fmt.Printf("5x5: %s\n", matrix.FiveByFive)
	fmt.Printf("4x4: %s\n", matrix.FourByFour)

	fmt.Printf("random: %s\n", matrix.GenerateRandomMatrix(8, 8))
}
