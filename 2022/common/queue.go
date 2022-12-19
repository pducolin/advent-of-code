package common

import "errors"

type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		items: []T{},
	}
}

func (queue *Queue[T]) Pop() T {
	if queue.IsEmpty() {
		panic(errors.New("empty queue"))
	}
	ret := queue.items[0]
	if len(queue.items) == 1 {
		queue.items = []T{}
	} else {
		queue.items = queue.items[1:]
	}
	return ret
}

func (queue *Queue[T]) Push(item T) {
	queue.items = append(queue.items, item)
}

func (queue *Queue[T]) IsEmpty() bool {
	return len(queue.items) == 0
}
