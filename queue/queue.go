package queue

import (
	"errors"

	"generic/list"
)

type Queue[T any] struct {
	list *list.List[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		list: list.New[T](),
	}
}

func (q Queue[T]) Length() int {
	return q.list.Length()
}

func (q Queue[T]) Push(v T) {
	q.list.PushBack(v)
}

func (q Queue[T]) Front() (T, error) {
	v, err := q.list.FrontValue()
	if err != nil {
		return v, errors.New("queue is empty")
	}
	return v, nil
}

func (q Queue[T]) Pop(v T) {
	q.list.PopFront()
}
