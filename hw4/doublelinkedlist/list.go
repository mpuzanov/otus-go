package doublelinkedlist

import (
	"fmt"
)

// List Список
type List struct {
	size      int   // размер
	firstItem *Item // points to first node of list
	lastItem  *Item // points to last node of list
}

// NewList ...
func NewList() *List {
	return &List{
		size:      0,
		firstItem: nil,
		lastItem:  nil,
	}
}

// AddValueList Добавляем элемент в список
func (l *List) AddValueList(v interface{}) *Item {
	return l.PushBack(v)
}

// ToString Вывод списка
func (l *List) ToString() string {
	s := ""
	node := l.firstItem
	for node != nil {
		s += fmt.Sprintln(node.value)
		node = node.Next()
	}
	return s
}

// Len длинна списка
func (l *List) Len() int {
	return l.size
}

// First первый Item
func (l *List) First() *Item {
	return l.firstItem
}

// Last последний Item
func (l *List) Last() *Item {
	return l.lastItem
}

// PushFront добавить значение в начало
func (l *List) PushFront(v interface{}) *Item {
	node := &Item{value: v}
	if l.firstItem == nil {
		l.lastItem = node
	} else {
		cur := l.firstItem
		cur.prev = node
		node.next = cur
	}
	l.firstItem = node
	l.size++
	return node
}

// PushBack добавить значение в конец
func (l *List) PushBack(v interface{}) *Item {
	node := &Item{value: v}
	if l.lastItem == nil {
		l.firstItem = node
	} else {
		cur := l.lastItem
		cur.next = node
		node.prev = l.lastItem
	}
	l.lastItem = node
	l.size++
	return node
}

// Remove удалить элемент
func (l *List) Remove(i Item) {

	if i.prev == nil {
		l.firstItem = i.next
	} else {
		i.prev.next = i.next
	}

	if i.next == nil {
		l.lastItem = i.prev
	} else {
		i.next.prev = i.prev
	}
	l.size--
}
