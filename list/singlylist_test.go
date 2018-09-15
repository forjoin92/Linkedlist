package list

import (
	"testing"
)

func TestNew(t *testing.T) {
	list := New()
	t.Log(list)
}

type item struct {
	key string
	value string
}

func TestList_Reverse(t *testing.T) {
	list := New()
	list.PushFront(&item{"1","11111"})
	list.PushFront(&item{"2","22222"})
	list.PushFront(&item{"3","33333"})
	list.PushFront(&item{"4","44444"})

	t.Log(list)
	for e := list.head; e != nil; e = e.next {
		t.Log(e.value)
	}

	list = list.Reverse()
	t.Log(list)
	for e := list.head; e != nil; e = e.next {
		t.Log(e.value)
	}
}
