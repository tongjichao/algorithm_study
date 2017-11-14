/*
	二分查找
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"cmcm.com/algorithm_study/data_structure"
	"cmcm.com/algorithm_study/search"
)

func main() {

	//	test_binnay_search()
	//test_stack()
	test_queue()

}

func test_queue() {
	queue := data_structure.NewQueue()

	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("1")
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
}

func test_stack() {
	stack := data_structure.NewStack()

	stack.Push("a")
	stack.Push("b")
	stack.Push("1")

	fmt.Println(stack.Size())

	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()

}

func test_binnay_search() {
	whitelist := []int{}

	for _, v := range os.Args[1:] {
		value, _ := strconv.Atoi(v)
		whitelist = append(whitelist, value)
	}
	sort.Ints(whitelist)

	fmt.Println(whitelist)

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("请输入一个数字，非数字退出")
		data, _, _ := reader.ReadLine()
		command := string(data)

		value, err := strconv.Atoi(command)
		if err != nil {
			log.Println("input is not a integer,exit...")
			break
		}

		fmt.Println("the index of ", value, " is ", search.Rank(value, whitelist))

	}

}
