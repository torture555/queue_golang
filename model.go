package main

type queue interface {
	Pop() interface{}
	Push(elem interface{})
}

func PopQueue(q queue) interface{} {
	return q.Pop()
}

func PushQueue(q queue, elem interface{}) {
	q.Push(elem)
}
