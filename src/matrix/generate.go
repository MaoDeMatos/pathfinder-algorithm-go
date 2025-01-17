package matrix

import (
	"math/rand"
)

func GenerateRandomMatrix(rows, cols int) Rows {
	// Create the matrix with default NormalNode
	matrix := make(Rows, rows)
	for i := range matrix {
		matrix[i] = make(Row, cols)
		for j := range matrix[i] {
			matrix[i][j] = Node{
				X:        j,
				Y:        i,
				NodeType: NormalNode,
			}
			// Place walls
			if rand.Float32() < 0.4 {
				matrix[i][j].NodeType = WallNode
			}
		}
	}

	// Set up start and end nodes
	startX, startY := rand.Intn(cols), rand.Intn(rows)
	endX, endY := rand.Intn(cols), rand.Intn(rows)

	matrix[startY][startX].NodeType = StartNode
	matrix[endY][endX].NodeType = EndNode

	return matrix
}
