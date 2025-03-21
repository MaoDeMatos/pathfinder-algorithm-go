package util

import "errors"

var ErrEmptyQueue = errors.New("queue is empty")

type Queue[T any] []T

func (q *Queue[T]) Dequeue() (T, error) {
	if len(*q) < 1 {
		var zeroValue T
		return zeroValue, ErrEmptyQueue
	}

	front := (*q)[0]
	*q = (*q)[1:]

	return front, nil
}

func (q *Queue[T]) Enqueue(value T) {
	*q = append(*q, value)
}

func (q *Queue[T]) Peek() (T, error) {
	if len(*q) < 1 {
		var zeroValue T
		return zeroValue, ErrEmptyQueue
	}

	return (*q)[0], nil
}

func (q *Queue[T]) Size() int {
	return len(*q)
}
