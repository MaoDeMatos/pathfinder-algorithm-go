package main

import (
	"errors"
	"slices"
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

func BreadthFirstSearch(input Input) (route []Node, err error) {
	rootNode, err := findNodeByType(input, StartNode)

	if err != nil {
		return route, err
	}

	q := NodeQueue{}
	q.Enqueue(rootNode)

	// create a parent map to save the interactions between nodes and recreate the path
	parents := map[Node]Node{} // initialize history

	// while queue has elements, keep iterating
	for q.Size() > 0 {
		// get first element & remove it from the queue
		currentNode, err := q.Dequeue()

		if err != nil {
			return route, err
		}

		// check if this is the target
		if currentNode.Type == EndNode {
			route = append(route, currentNode)

			// build route from the node history
			for parentNode, exists := parents[currentNode]; exists; parentNode, exists = parents[parentNode] {
				route = append(route, parentNode)
				if parentNode.Type == StartNode {
					break
				}
			}

			// reverse the route since we built it from end to start
			slices.Reverse(route)

			return route, nil
		}

		// find neighbors
		neighbors := []Node{}
		for _, dir := range Directions {
			x, y := dir[1], dir[0]
			neighborNode, notFound := findNodeByCoordinates(
				input,
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

	return route, ErrNoResult
}

func findNodeByType(input Input, nodeType NodeType) (Node, error) {
	for rowIndex := range input {
		for _, node := range input[rowIndex] {
			if node.Type == nodeType {
				return node, nil
			}
		}
	}

	return Node{}, ErrNotFound
}

func findNodeByCoordinates(input Input, x, y int) (Node, error) {
	for rowIndex := range input {
		for _, node := range input[rowIndex] {
			if node.X == x && node.Y == y {
				return node, nil
			}
		}
	}

	return Node{}, ErrNotFound
}
