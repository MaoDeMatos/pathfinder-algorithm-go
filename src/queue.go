package main

import "errors"

var ErrEmptyQueue = errors.New("queue is empty")

type NodeQueue []Node

func (q *NodeQueue) Dequeue() (Node, error) {
	if len(*q) < 1 {
		return Node{}, ErrEmptyQueue
	}

	front := (*q)[0]
	*q = (*q)[1:]

	return front, nil
}

func (q *NodeQueue) Enqueue(value Node) {
	*q = append(*q, value)
}

func (q *NodeQueue) Peek() (Node, error) {
	if len(*q) < 1 {
		return Node{}, ErrEmptyQueue
	}

	return (*q)[0], nil
}

func (q *NodeQueue) Size() int {
	return len(*q)
}
