package list

import (
	"testing"
	"fmt"
)

func TestNew(t *testing.T) {
	list := New()
	t.Log(list)
}

type item struct {
	key   string
	value string
}

func (i *item) String() string {
	return fmt.Sprintf("key=%s, value=%s", i.key, i.value)
}

func Init(l *List) *List {
	e := l.PushBack(&item{"3", "33333"})
	l.PushFront(&item{"2", "22222"})
	l.PushBack(&item{"5", "55555"})
	l.PushFront(&item{"1", "11111"})
	l.InsertAfter(&item{"4", "44444"}, e)
	return l
}

func TestList_Front(t *testing.T) {
	list := New()
	list = Init(list)
	t.Log(list.Front().value)
}

func TestList_Back(t *testing.T) {
	list := New()
	list = Init(list)
	t.Log(list.Back().value)
}

func printList(t *testing.T, l *List) {
	t.Log("l.len=", l.Len())
	t.Log("head=", l.head, "tail=", l.tail)
	for e := l.head; e != nil; e = e.Next() {
		t.Log(e.value)
	}
}

func TestList_Remove(t *testing.T) {
	list := New()
	list = Init(list)

	e := list.Front()
	e1 := list.Remove(e)
	t.Log("delete e1:",e1)
	printList(t, list)

	e = list.Front()
	e2 := list.Remove(e.Next())
	t.Log("delete e2:", e2)
	printList(t, list)

	e = list.Back()
	e3 := list.Remove(e)
	t.Log("delete e3:", e3)
	printList(t, list)
}

func TestList_Reverse(t *testing.T) {
	list := New()
	list = Init(list)

	printList(t, list)

	list = list.Reverse()
	printList(t, list)
}
