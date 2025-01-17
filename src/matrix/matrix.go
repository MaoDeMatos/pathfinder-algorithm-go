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

func (m Rows) String() string {
	result := ""

	for _, row := range m {
		for _, node := range row {
			// result += string(node.NodeType)
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

	// // Remove last newline
	// result = strings.TrimRight(result, "\n")

	return fmt.Sprintf("\n%s", result)
}
