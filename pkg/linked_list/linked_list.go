package linked_list

import (
	"errors"
	"sync"
)

type DoublyLinkedList[T any] struct {
	head    *Node[T]
	tail    *Node[T]
	current *Node[T]

	mu  sync.RWMutex
	len int
}

type Node[T any] struct {
	next *Node[T]
	prev *Node[T]
	data T
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (l *DoublyLinkedList[T]) Append(data ...T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, v := range data {
		newNode := &Node[T]{data: v, prev: nil, next: nil}

		if l.head == nil {
			l.head = newNode
			l.tail = newNode
			l.current = newNode
			return
		}

		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
		l.len++
	}
}

func (l *DoublyLinkedList[T]) Get() T {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.current.data
}

func (l *DoublyLinkedList[T]) List() []T {
	res := make([]T, l.len)

	i := 0
	current := l.head
	for current != nil {
		res[i] = current.data
		current = current.next
		i++
	}

	return res
}

func (l *DoublyLinkedList[T]) Next() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.len == 0 {
		return errors.New("empty list")
	}

	if l.current.next == nil {
		return errors.New("end of list")
	}

	l.current = l.current.next
	return nil
}

func (l *DoublyLinkedList[T]) Prev() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.len == 0 {
		return errors.New("empty list")
	}

	if l.current.prev == nil {
		return errors.New("end of list")
	}

	l.current = l.current.prev
	return nil
}
