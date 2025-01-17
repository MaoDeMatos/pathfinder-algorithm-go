package matrix

import (
	"fmt"
)

type nodeType string

const (
	StartNode  nodeType = "s"
	EndNode             = "e"
	WallNode            = "w"
	NormalNode          = "n"
)

type Node struct {
	X        int
	Y        int
	NodeType nodeType
	// visited  bool
}

type Row []Node
type Rows []Row

// Implement the Stringer interface
// Returns a formatted string
func (m Rows) String() string {
	result := ""

	for _, row := range m {
		for _, node := range row {
			switch node.NodeType {
			case StartNode:
				result += "S "
			case EndNode:
				result += "E "
			case WallNode:
				result += "W "
			default:
				result += ". "
			}
		}
		result += "\n"
	}

	return fmt.Sprintf("\n%s", result)
}
