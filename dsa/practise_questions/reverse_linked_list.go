// Reverse Even Length Groups in a LinkedList
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// inserting a node in the list
func (node *ListNode) InsertNode(val int) {
	//  create a new node
	tempNode := &ListNode{}
	// use the temp node to hold the new value received
	tempNode.Val = val
	tempNode.Next = nil
	// check if the list is empty making this our first node
	if node != nil {
		// if the list is not empty, traverse to the end of the list
		for node.Next != nil {

			// check to see if we would lose the head node in this approach
			node = node.Next
		}
		// if we get here, we are at the end of the list seeing that the last node next property evaluates to nil
		node.Next = tempNode
	}
}

type GroupedNodes struct {
	GroupedVals []int
}

// linked list :  [5,2,6,3,9,1,7,3,8,4]
// Break each node into groups, with the next group's length equal to the previous group's length plus one
func ReverseEvenLength(head *ListNode) *ListNode {
	var temp *ListNode = head
	// current index of the node
	index := 0
	var tempArray []int

	// traverse a linked list
	for temp != nil {
		index++
		tempArray = append(tempArray, temp.Val)
		fmt.Printf("Temp array : %v @ index %d \n\n", tempArray, index)
		if index == 2 {
			break
		}
		// using the index to determine the length of each group
		temp = temp.Next
	}
	return nil
}

// If the leftover nodes are not equal to previous group's length plus one make the last group with the leftoiver nodes
// Next group starts from previous group's last node
// check if each group's length is an even number
// if that is true; reverse the content of the group
// merge groups(including the reveresed data) to make a linkedlist
// return the head of the linked list
