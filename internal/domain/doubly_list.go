package domain

import (
	"fmt"
)

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	next *Node
	prev *Node
	data *Song
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (l *DoublyLinkedList) Append(data *Song) {
	l.head.data.Meta.Mu.Lock()
	defer l.head.data.Meta.Mu.Unlock()

	newNode := &Node{data: data, prev: nil, next: nil}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	newNode.prev = l.tail
	l.tail.next = newNode
	l.tail = newNode
}

func (l *DoublyLinkedList) Get() SongResponse {
	l.head.data.Meta.Mu.Lock()
	defer l.head.data.Meta.Mu.Unlock()

	res := SongResponse{
		ID:       l.head.data.ID,
		Name:     l.head.data.Name,
		Duration: l.head.data.Duration,
	}

	return res
}

func (l *DoublyLinkedList) AppendMany(data ...*Song) {
	for _, v := range data {
		l.Append(v)
	}
}

func (l *DoublyLinkedList) GetAll() []*Song {
	res := make([]*Song, 0)

	current := l.head
	for current != nil {
		res = append(res, current.data)
		current = current.next
	}

	return res
}

func (l *DoublyLinkedList) MoveBackward() {
	current := l.tail
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.prev
	}
}
