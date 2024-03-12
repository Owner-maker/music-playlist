package domain

import "errors"

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	cursor *Node
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
		l.cursor = newNode
		return
	}

	newNode.prev = l.tail
	l.tail.next = newNode
	l.tail = newNode
}

func (l *DoublyLinkedList) Get() Info {
	l.head.data.Meta.Mu.Lock()
	defer l.head.data.Meta.Mu.Unlock()

	res := Info{
		ID:       l.head.data.Info.ID,
		Name:     l.head.data.Info.Name,
		Duration: l.head.data.Info.Duration,
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

func (l *DoublyLinkedList) Current() *Song {
	return l.cursor.data
}

func (l *DoublyLinkedList) Next() error {
	l.cursor.data.Meta.Mu.Lock()
	defer l.cursor.data.Meta.Mu.Unlock()

	if l.cursor == nil {
		return errors.New("empty list")
	}

	if l.cursor.next == nil {
		return errors.New("end of list")
	}

	l.cursor = l.cursor.next
	return nil
}

func (l *DoublyLinkedList) Prev() error {
	l.cursor.data.Meta.Mu.Lock()
	defer l.cursor.data.Meta.Mu.Unlock()

	if l.cursor == nil {
		return errors.New("empty list")
	}

	if l.cursor.prev == nil {
		return errors.New("end of list")
	}

	l.cursor = l.cursor.prev
	return nil
}
