package main

import (
	"errors"
	"slices"

	"github.com/MaoDeMatos/pathfinder-algorithm-go/src/internal/util"
)

var (
	ErrNoResult = errors.New("target not found")
	ErrNotFound = errors.New("node not found")
)

var Directions = [4][2]int{
	{0, 1},  // Right
	{0, -1}, // Left
	{1, 0},  // Bottom
	{-1, 0}, // Top
}

type NodeType rune

const (
	StartNode  NodeType = 's'
	EndNode    NodeType = 'e'
	WallNode   NodeType = 'w'
	NormalNode NodeType = 'n'
)

type Node struct {
	X    int
	Y    int
	Type NodeType
}

type Row []Node
type Input []Row

func (in *Input) findNodeByType(nodeType NodeType) (Node, error) {
	for rowIndex := range *in {
		for _, node := range (*in)[rowIndex] {
			if node.Type == nodeType {
				return node, nil
			}
		}
	}

	return Node{}, ErrNotFound
}

func (in *Input) findNodeByCoordinates(x, y int) (Node, error) {
	for rowIndex := range *in {
		for _, node := range (*in)[rowIndex] {
			if node.X == x && node.Y == y {
				return node, nil
			}
		}
	}

	return Node{}, ErrNotFound
}

type Solution []Node

func (s *Solution) contains(node Node) bool {
	return slices.Contains(*s, node)
}

//
// Algorithm
//

func BreadthFirstSearch(input Input) (solution Solution, err error) {
	rootNode, err := input.findNodeByType(StartNode)

	if err != nil {
		return solution, err
	}

	q := util.Queue[Node]{}
	q.Enqueue(rootNode)

	// create a parent map to save the interactions between nodes and recreate the path
	parents := map[Node]Node{} // initialize history

	// while queue has elements, keep iterating
	for q.Size() > 0 {
		// get first element & remove it from the queue
		currentNode, err := q.Dequeue()

		if err != nil {
			return solution, err
		}

		// check if this is the target
		if currentNode.Type == EndNode {
			solution = append(solution, currentNode)

			// build route from the node history
			for parentNode, exists := parents[currentNode]; exists; parentNode, exists = parents[parentNode] {
				solution = append(solution, parentNode)
				if parentNode.Type == StartNode {
					break
				}
			}

			// reverse the route since we built it from end to start
			slices.Reverse(solution)

			return solution, nil
		}

		// find neighbors
		neighbors := []Node{}
		for _, dir := range Directions {
			x, y := dir[1], dir[0]
			neighborNode, notFound := input.findNodeByCoordinates(
				currentNode.X+x,
				currentNode.Y+y,
			)

			// If node exists and is not a wall, you can add it to the list of nodes to add to the queue
			if notFound == nil && neighborNode.Type != WallNode {
				neighbors = append(neighbors, neighborNode)
			}
		}

		for _, neighbor := range neighbors {
			// check if the neighbor has not already been visited
			if _, visited := parents[neighbor]; !visited {
				parents[neighbor] = currentNode // add neighbor to parents map associated to current node value
				q.Enqueue(neighbor)             // add neighbor to the end of the queue
			}
		}
	}

	return solution, ErrNoResult
}
