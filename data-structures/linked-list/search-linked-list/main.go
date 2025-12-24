package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func SearchKeyIterative(head *Node, key int) bool {
	for head != nil {
		if head.Data == key {
			return true
		}
		head = head.Next
	}
	return false
}

func SearchKeyRecursive(head *Node, key int) bool{
	if head == nil{
		return  false
	}

	if head.Data == key {
		return true
	}

	return SearchKeyRecursive(head.Next, key)
}

func main() {
	head := &Node{Data: 1}
	head.Next = &Node{Data: 3}
	head.Next.Next = &Node{Data: 7}
	head.Next.Next.Next = &Node{Data: 5}
	key := 7

	fmt.Println(SearchKeyIterative(head, key))
	fmt.Println(SearchKeyRecursive(head, key))
}