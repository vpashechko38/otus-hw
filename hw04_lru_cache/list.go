package hw04lrucache

import (
	"fmt"
)

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	Print() string
}

type (
	ListItem struct {
		Value interface{}
		Next  *ListItem
		Prev  *ListItem
	}

	list struct{}
)

var (
	head   *ListItem
	back   *ListItem
	length int
)

func NewList() List {
	return list{}
}

func (l list) Len() int {
	return length
}

func (l list) Front() *ListItem {
	return head
}

func (l list) Back() *ListItem {
	return back
}

func (l list) PushFront(v interface{}) *ListItem {
	length++
	if head == nil {
		head = &ListItem{Value: v}
		back = head
	} else {
		prevHead := head
		newHead := &ListItem{Value: v, Next: prevHead}
		prevHead.Prev = newHead
		head = newHead
	}

	return head
}

func (l list) PushBack(v interface{}) *ListItem {
	length++
	if head == nil {
		head = &ListItem{Value: v}
		back = head
	} else {
		prevBack := back
		newBack := &ListItem{Value: v, Prev: prevBack}
		prevBack.Next = newBack
		back = newBack
	}

	return back
}

func (l list) Remove(i *ListItem) {
	if i == nil {
		return
	}

	length--

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	i.Next = nil
	i.Prev = nil
}

func (l list) MoveToFront(i *ListItem) {
	if i == nil {
		return
	}

	if i.Prev == nil {
		return
	}

	i.Prev.Next = i.Next

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		back = i.Prev
	}

	head.Prev = i
	i.Next = head

	head = i
	head.Prev = nil
}

func (l list) MoveToBack(i *ListItem) {
	if i == nil {
		return
	}

	if i.Next == nil {
		return
	}

	back.Next = i
	i.Prev = back

	back = i
	back.Next = nil
}

func (l list) Print() string {
	r := ""
	h := head
	for {
		r += fmt.Sprintf("%v ", h.Value.(*Item).Value)
		if h.Next == nil {
			return r
		}
		h = h.Next
	}
}
