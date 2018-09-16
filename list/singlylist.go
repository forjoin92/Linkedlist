package list

type Element struct {
	value interface{}
	next  *Element
	list  *List
}

func (e *Element) Next() *Element {
	if n := e.next; e.list != nil {
		return n
	}
	return nil
}

type List struct {
	head *Element
	tail *Element
	len  int
}

func (l *List) Init() *List {
	l.head = nil
	l.tail = nil
	l.len = 0
	return l
}

func New() *List {
	return new(List).Init()
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.head
}

func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.tail
}

func (l *List) insert(e, at *Element) *Element {
	if at == nil {
		e.next = l.head
		l.head = e
	} else {
		n := at.next
		at.next = e
		e.next = n

	}

	e.list = l
	if e.next == nil {
		l.tail = e
	}
	l.len++
	return e
}

func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{value: v}, at)
}

func (l *List) InsertAfter(v interface{}, at *Element) *Element {
	if at.list != l {
		return nil
	}
	return l.insertValue(v, at)
}

func (l *List) remove(e *Element) *Element {
	if l.head == e {
		l.head = l.head.next
		e.next = nil
		e.list = nil
		if e == l.tail {
			l.tail = nil
		}
		l.len--
		return e
	}

	for pre, n := l.head, l.head.next; n != nil; pre, n = n, n.next {
		if n == e {
			pre.next = n.next
			e.next = nil
			e.list = nil
			if e == l.tail {
				l.tail = pre
			}
			l.len--
			return e
		}
	}
	return nil
}

func (l *List) Remove(e *Element) interface{} {
	if e.list == l {
		l.remove(e)
	}
	return e.value
}

func (l *List) PushFront(v interface{}) *Element {
	return l.insertValue(v, nil)
}

func (l *List) PushBack(v interface{}) *Element {
	return l.insertValue(v, l.tail)
}

func (l *List) Reverse() *List {
	pre, n := l.head, l.head.next
	pre.next = nil
	for n != nil {
		n1 := n.next
		n.next = pre
		pre = n
		n = n1
	}
	l.head, l.tail = l.tail, l.head
	return l
}
