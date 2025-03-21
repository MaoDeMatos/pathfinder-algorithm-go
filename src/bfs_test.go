package main

import (
	"testing"
)

var simpleInput = Input{
	{
		Node{X: 0, Y: 0, Type: StartNode},
		Node{X: 1, Y: 0, Type: NormalNode},
		Node{X: 2, Y: 0, Type: EndNode},
	},
}

var errorInput = Input{
	{
		Node{X: 0, Y: 0, Type: StartNode},
		Node{X: 1, Y: 0, Type: NormalNode},
		Node{X: 2, Y: 0, Type: NormalNode},
	},
}

var fourByFourInput = Input{
	{
		Node{X: 0, Y: 0, Type: StartNode},
		Node{X: 1, Y: 0, Type: NormalNode},
		Node{X: 2, Y: 0, Type: WallNode},
		Node{X: 3, Y: 0, Type: NormalNode},
	},
	{
		Node{X: 0, Y: 1, Type: NormalNode},
		Node{X: 1, Y: 1, Type: WallNode},
		Node{X: 2, Y: 1, Type: NormalNode},
		Node{X: 3, Y: 1, Type: NormalNode},
	},
	{
		Node{X: 0, Y: 2, Type: NormalNode},
		Node{X: 1, Y: 2, Type: NormalNode},
		Node{X: 2, Y: 2, Type: WallNode},
		Node{X: 3, Y: 2, Type: WallNode},
	},
	{
		Node{X: 0, Y: 3, Type: WallNode},
		Node{X: 1, Y: 3, Type: NormalNode},
		Node{X: 2, Y: 3, Type: NormalNode},
		Node{X: 3, Y: 3, Type: EndNode},
	},
}

var testCases = []struct {
	name     string
	input    Input
	expected []Node
	error    error
}{
	{
		name:     "simplest input",
		input:    simpleInput,
		expected: simpleInput[0],
		error:    nil,
	},
	{
		name:     "error input - no end node",
		input:    errorInput,
		expected: nil,
		error:    ErrNotFound,
	},
	{
		name:  "4x4 input",
		input: fourByFourInput,
		expected: []Node{
			{X: 0, Y: 0, Type: StartNode},
			{X: 0, Y: 1, Type: NormalNode},
			{X: 0, Y: 2, Type: NormalNode},
			{X: 1, Y: 2, Type: NormalNode},
			{X: 1, Y: 3, Type: NormalNode},
			{X: 2, Y: 3, Type: NormalNode},
			{X: 3, Y: 3, Type: EndNode},
		},
		error: nil,
	},
}

func TestBreadthFirstSearch(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result, err := BreadthFirstSearch(test.input)

			if !areNodesListsEqual(test.expected, result) {
				t.Errorf("expected result: %v, got %v. expected error : %v, got %v", test.expected, result, test.error, err)
			}
		})
	}
}

func areNodesListsEqual(a, b []Node) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i].X != b[i].X || a[i].Y != b[i].Y {
			return false
		}
	}

	return true
}
