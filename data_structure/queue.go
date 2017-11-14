package data_structure

import (
	"fmt"
)

type Queue struct {
	First *Node
	Last  *Node
	N     int64
}

func NewQueue() *Queue {
	queue := new(Queue)
	queue.First = nil
	queue.Last = nil
	queue.N = int64(0)

	if queue.IsEmpty() {
		fmt.Println("you created a empty stack")
	}
	return queue
}

func (queue *Queue) IsEmpty() bool {
	return queue.First == nil
}

func (queue *Queue) Size() int64 {
	return queue.N
}

func (queue *Queue) Enqueue(item interface{}) {

	oldlast := queue.Last

	queue.Last = new(Node)
	queue.Last.Item = item
	queue.Last.Next = nil

	if queue.IsEmpty() {
		queue.First = queue.Last
	} else {
		oldlast.Next = queue.Last
	}
	queue.N++

	fmt.Println("add ->", item, queue.Size())
}

func (queue *Queue) Dequeue() interface{} {

	item := queue.First.Item

	queue.First = queue.First.Next

	if queue.IsEmpty() {
		queue.Last = nil
	}

	queue.N--

	fmt.Println("pop ->", item, queue.Size())
	return item
}
