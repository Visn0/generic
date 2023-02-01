package list

import "errors"

type Node[T any] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

type List[T any] struct {
	front  *Node[T]
	back   *Node[T]
	length int
}

func New[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Empty() bool {
	return l.length == 0
}

func (l *List[T]) Length() int {
	return l.length
}

func (l *List[T]) Front() *Node[T] {
	return l.front
}

func (l *List[T]) FrontValue() (T, error) {
	if l.front == nil {
		return *new(T), errors.New("empty list")
	}
	return l.front.value, nil
}

func (l *List[T]) Back() *Node[T] {
	return l.back
}

func (l *List[T]) BackValue() (T, error) {
	if l.back == nil {
		return *new(T), errors.New("empty list")
	}
	return l.back.value, nil
}

func (l *List[T]) PushBack(v T) {
	n := &Node[T]{
		value: v,
	}
	if l.back == nil {
		l.front = n
	} else {
		n.prev = l.back
		l.back.next = n
	}
	l.back = n
	l.length++
}

func (l *List[T]) PushFront(v T) {
	n := &Node[T]{
		value: v,
		next:  l.front,
	}
	if l.front == nil {
		l.back = n
	} else {
		l.front.prev = n
	}
	l.front = n
	l.length++
}

func (l *List[T]) PopFront() {
	// list not empy
	if l.front != nil {
		l.front = l.front.next
		// list was length 1
		if l.front == nil {
			l.back = nil
		} else {
			l.front.prev = nil
		}
		l.length--
	}
}

func (l *List[T]) PopBack() {
	if l.back != nil {
		l.back = l.back.prev
		if l.back == nil {
			l.front = nil
		} else {
			l.back.next = nil
		}
		l.length--
	}
}

func (l *List[T]) Remove(n *Node[T]) {
	if l.front == n {
		l.PopFront()
	} else if l.back == n {
		l.PopBack()
	} else {
		n.prev.next = n.next
		n.next.prev = n.prev
		l.length--
	}
}
