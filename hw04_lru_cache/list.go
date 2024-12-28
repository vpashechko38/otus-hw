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

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

var (
	head   *ListItem = nil
	back   *ListItem = nil
	length int       = 0
)

type list struct {
}

func NewList() List {
	return list{}
}

func (l list) Len() int {
	return length
}

// Front Получить первый элемент списка
func (l list) Front() *ListItem {
	return head
}

// Back Получить последний элемент списка
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

	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev

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
