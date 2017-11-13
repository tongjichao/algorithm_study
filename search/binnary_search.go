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
)

func main() {

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

		fmt.Println("the index of ", value, " is ", rank(value, whitelist))

	}

}

func rank(key int, a []int) int {
	lo := int(0)
	hi := int(len(a) - 1)
	for {
		if lo > hi {
			break
		}

		mid := lo + (hi-lo)/2
		if key < a[mid] {
			hi = mid - 1
			continue
		}

		if key > a[mid] {
			lo = mid + 1
			continue
		}

		return mid
	}

	return -1
}
