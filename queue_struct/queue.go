package queue_struct

type QueueStruct struct {
	first *elemQueue
	last  *elemQueue
}

type elemQueue struct {
	value interface{}
	next  *elemQueue
	prev  *elemQueue
}

func (queue *QueueStruct) Pop() interface{} {
	if queue.first.value == nil {
		return nil
	}
	res := queue.first.value
	queue.first = queue.first.prev
	if queue.first != nil {
		queue.first.next = nil
	}
	return res
}

func (queue *QueueStruct) Push(elem interface{}) {
	newElem := elemQueue{value: elem, next: queue.last, prev: nil}
	if queue.first == nil {
		queue.first = &newElem
	}
	if queue.last == nil {
		queue.last = &newElem
		return
	}
	queue.last.prev = &newElem
	queue.last = &newElem
}
