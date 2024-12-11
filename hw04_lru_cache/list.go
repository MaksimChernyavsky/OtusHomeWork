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

type list struct {
	top *ListItem
	bot *ListItem
	c   int
}

func (l *list) Len() int {
	return l.c
}
func (l *list) Front() *ListItem {
	return l.top
}
func (l *list) Back() *ListItem {
	return l.bot
}
func (l *list) PushFront(v interface{}) *ListItem {
	newListItem := &ListItem{
		Value: v,
		Next:  l.top,
		Prev:  nil,
	}
	if l.top != nil {
		l.top.Prev = newListItem
	}
	l.top = newListItem
	if l.bot == nil {
		l.bot = newListItem
	}
	l.c += 1
	return newListItem
}
func (l *list) PushBack(v interface{}) *ListItem {
	newListItem := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  l.bot,
	}
	if l.bot != nil {
		l.bot.Next = newListItem
	}
	l.bot = newListItem
	if l.top == nil {
		l.top = newListItem
	}
	l.c += 1
	return newListItem
}
func (l *list) Remove(item *ListItem) {
	if item.Prev != nil {
		item.Prev.Next = item.Next
	}
	if item.Next != nil {
		item.Next.Prev = item.Prev
	}
	if l.top == item {
		l.top = item.Next
	}
	if l.bot == item {
		l.bot = item.Prev
	}
	l.c -= 1
}
func (l *list) MoveToFront(item *ListItem) {
	l.Remove(item)
	item.Prev = nil
	item.Next = l.top
	if l.top != nil {
		l.top.Prev = item
	}
	l.top = item
	l.c += 1
}

func NewList() List {
	return new(list)
}
