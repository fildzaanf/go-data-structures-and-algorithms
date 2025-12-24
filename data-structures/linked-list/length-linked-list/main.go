package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func NewNode(data int) *Node {
	return &Node{
		Data: data,
		Next: nil,
	}
}

func LengthIterative(head *Node) int {
	var count int

	for head != nil {
		count++
		head = head.Next
	}

	return count
}

func LengthRecursive(head *Node) int {
	if head == nil {
		return 0
	}

	return 1 + LengthRecursive(head.Next)
}

func main() {
	head := NewNode(1)
	head.Next = NewNode(4)
	head.Next.Next = NewNode(2)
	head.Next.Next.Next = NewNode(3)

	fmt.Println("Length Iterative:", LengthIterative(head))
	fmt.Println("Length Recursive:", LengthRecursive(head))
}
