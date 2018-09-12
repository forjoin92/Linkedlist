package common

type Element struct {
	value interface{}
	next  *Element
}

type List struct {
	root Element
	len  int
}
