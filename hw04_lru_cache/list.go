package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type (
	ListItem struct {
		Value interface{}
		Next  *ListItem
		Prev  *ListItem
	}

	list struct {
		head   *ListItem
		back   *ListItem
		length int
	}
)

func NewList() List {
	return &list{
		head:   nil,
		back:   nil,
		length: 0,
	}
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	l.length++
	if l.head == nil {
		l.head = &ListItem{Value: v}
		l.back = l.head
	} else {
		prevHead := l.head
		newHead := &ListItem{Value: v, Next: prevHead}
		prevHead.Prev = newHead
		l.head = newHead
	}

	return l.head
}

func (l *list) PushBack(v interface{}) *ListItem {
	l.length++
	if l.head == nil {
		l.head = &ListItem{Value: v}
		l.back = l.head
	} else {
		prevBack := l.back
		newBack := &ListItem{Value: v, Prev: prevBack}
		prevBack.Next = newBack
		l.back = newBack
	}

	return l.back
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}

	if i.Prev == nil {
		l.head = i.Next
	}
	if i.Next == nil {
		l.back = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	i.Next = nil
	i.Prev = nil

	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
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
		l.back = i.Prev
	}

	l.head.Prev = i
	i.Next = l.head

	l.head = i
	l.head.Prev = nil
}

func (l *list) MoveToBack(i *ListItem) {
	if i == nil {
		return
	}

	if i.Next == nil {
		return
	}

	l.back.Next = i
	i.Prev = l.back

	l.back = i
	l.back.Next = nil
}
