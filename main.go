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

	"algorithm_study/data_structure"
	"algorithm_study/search"
	"algorithm_study/union_find"
)

func main() {

	//	test_binnay_search()
	//test_stack()
	//test_queue()

	test_union()

}

//测试union
func test_union() {

	data := [...][2]int64{{4, 3}, {3, 7}, {6, 5}, {2, 4}, {2, 1}, {5, 0}, {7, 2}, {6, 1}}

	uf := union_find.NewUF(int64(len(data)))
	for _, v := range data {
		p := v[0]
		q := v[1]

		fmt.Println(uf.Id)
		if uf.Connected(p, q) {
			continue
		}

		uf.UnionFindQuick(p, q)

		fmt.Println(p, "  ", q)
	}

	fmt.Println(uf.Count, "components")

}

//队列测试
func test_queue() {
	queue := data_structure.NewQueue()

	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("1")
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
}

//栈测试
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

//二分查找测试
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
