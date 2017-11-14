package data_structure

import (
	"errors"
	"fmt"
)

type Node struct {
	Item interface{}
	Next *Node
}

type Stack struct {
	First *Node
	N     int64
}

func NewStack() *Stack {
	stack := new(Stack)
	stack.First = nil
	stack.N = int64(0)

	if stack.IsEmpty() {
		fmt.Println("you created a empty stack")
	}
	return stack
}

func (stack *Stack) IsEmpty() bool {
	return stack.First == nil
}

func (stack *Stack) Size() int64 {
	return stack.N
}

func (stack *Stack) Push(item interface{}) {
	oldFirst := stack.First
	first := new(Node)
	first.Item = item
	first.Next = oldFirst

	stack.First = first
	stack.N++
	fmt.Println("push == >", item, "长度:", stack.Size())
}

func (stack *Stack) Pop() (interface{}, error) {

	if stack.First == nil {
		return nil, errors.New("stack is empty")
	}

	item := stack.First.Item
	stack.First = stack.First.Next
	stack.N--
	fmt.Println("pop == >", item, "长度:", stack.Size())
	return item, nil
}
