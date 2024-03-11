package queue_slice

type QueueSlice struct {
	slice []interface{}
}

func (queue *QueueSlice) Pop() interface{} {
	if len(queue.slice) == 0 {
		return nil
	}

	returnElem := queue.slice[0]
	queue.slice = queue.slice[1:]
	return returnElem
}

func (queue *QueueSlice) Push(elem interface{}) {
	queue.slice = append(queue.slice, elem)
}
