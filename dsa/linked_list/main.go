package main

import (
	"fmt"
	"os"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// inserting a node in the list
func (node *ListNode) insertNode(val int) {
	fmt.Println("====temp node value======", node.Val)
	//  create a new node
	tempNode := &ListNode{}
	// use the temp node to hold the new value received
	tempNode.Val = val
	tempNode.Next = nil
	// check if the list is empty making this our first node
	if node == nil {
		node = tempNode
	} else {
		// if the list is not empty, traverse to the end of the list
		for node.Next != nil {

			// check to see if we would lose the head node in this approach
			node = node.Next

		}

		// if we get here, we are at the end of the list seeing that the last node next property evaluates to nil
		node.Next = tempNode
	}
}

// the delete first node is taking the next node and assigning it to the current node
func (node *ListNode) deleteFirstNode() *ListNode {
	node = node.Next
	return node
}

// deleting last node without using a pointer reference of the last node or keeping a reference to the last node
func (node *ListNode) deleteLastNode() {
	// find the last node by using the nesting
	for node.Next.Next != nil {
		// assign the next node to the current node
		node = node.Next
	}
	// make the next node of the current node nil
	node.Next = nil
}

// print the list
func (node *ListNode) printList() {
	var temp *ListNode
	temp = node
	for temp != nil {
		fmt.Printf("%d ->  with value %d\n\n", temp.Val, temp.Next)
		temp = temp.Next
	}
}

func main() {
	head := new(ListNode)

	fmt.Println("====temp node value======", head.Val)
	// head.insertNode(2)
	// head.insertNode(40)
	// head.insertNode(32)
	// head.insertNode(82)
	// head.insertNode(112)
	// head.insertNode(82)
	// head.insertNode(90)
	// head.printList()

	// lets build a console app to create a linked list
	var choice string

	for {
		fmt.Println("Enter a number |||||")
		fmt.Println("               ⌄⌄⌄⌄⌄")
		fmt.Println("Enter 0 to exit")
		fmt.Println("Enter 1 to insert a node")
		fmt.Println("Enter 2 to delete the first node")
		fmt.Println("Enter 3 to delete the last node")
		fmt.Println("Enter 4 to print the list")
		fmt.Scan(&choice)
		switch choice {
		case "0":
			os.Exit(0)
		case "1":
			var val string
			fmt.Println("Enter a value in numbers only")
			fmt.Scan(&val)
			num, _ := strconv.Atoi(val)
			head.insertNode(num)
		case "2":
			// the new node now gets assigned to the head node
			head = head.deleteFirstNode()
		case "3":
			head.deleteLastNode()
		case "4":
			head.printList()
		}

	}

}
