package main

import (
	"fmt"
	"queue_golang/queue_slice"
	"queue_golang/queue_struct"
	"strconv"
	"sync"
	"time"
)

func main() {

	myValues := makeSlice()

	queueSlice := queue_slice.QueueSlice{}
	queueStruct := queue_struct.QueueStruct{}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go popPushValues(&wg, &queueSlice, myValues, "Slice: ")
	go popPushValues(&wg, &queueStruct, myValues, "Struct: ")

	wg.Wait()

}

func makeSlice() []string {

	var res []string

	for i := 0; i < 2*100000000; i++ {
		res = append(res, "dsaddfsfs"+strconv.Itoa(i))
	}

	return res

}

func popPushValues(wg *sync.WaitGroup, queue queue, slice []string, typeQueue string) {

	defer wg.Done()

	_, min, sec := time.Now().Clock()

	for _, v := range slice {
		PushQueue(queue, v)
	}

	for _, _ = range slice {
		_ = PopQueue(queue)
	}

	_, minA, secA := time.Now().Clock()

	fmt.Println(typeQueue + strconv.Itoa((minA*60+secA)-(min*60+sec)))

}
