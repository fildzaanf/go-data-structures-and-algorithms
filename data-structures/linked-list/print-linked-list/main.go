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

func PrintIterative(head *Node) {
	for head != nil {
		fmt.Print(head.Data, " ")

		if head.Next != nil {
		fmt.Print("-> ")
		}

		head = head.Next
	}
}

func Printrecursive(head *Node) {
	if head == nil {
		return
	}

	fmt.Print(head.Data)

	if head.Next != nil {
		fmt.Print(" -> ")
	}

	Printrecursive(head.Next)
}

func main() {
	head := NewNode(1)
	head.Next = NewNode(7)
	head.Next.Next = NewNode(4)
	head.Next.Next.Next = NewNode(2)

	PrintIterative(head)
	fmt.Println()
	Printrecursive(head)
}