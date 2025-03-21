package main

import (
	"testing"
)

var testInput = Input{
	{
		Node{X: 0, Y: 0, Value: StartNode},
		Node{X: 1, Y: 0, Value: NormalNode},
		Node{X: 2, Y: 0, Value: EndNode},
	},
}

var tests = []struct {
	name     string
	input    Input
	start    Node
	end      Node
	expected []Node
	error    error
}{
	{
		name:     "simplest path",
		input:    testInput,
		expected: testInput[0],
		error:    nil,
	},
	// {
	// 	name:     "no result",
	// 	input:    testInput,
	// 	expected: []Node{},
	// 	error:    ErrNoResult,
	// },
}

func TestBreadthFirstSearch(t *testing.T) {
	for _, test := range tests {
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
