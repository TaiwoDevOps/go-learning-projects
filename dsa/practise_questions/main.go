// single string swap algorithm

package main

func main() {
	// res := AreAlmostEqual("bank", "kanb")
	// fmt.Println("res:-->", res)

	// reverse even length
	head := new(ListNode)

	// [5,2,6,3,9,1,7,3,8,4]
	head.InsertNode(5)
	head.InsertNode(2)
	head.InsertNode(6)
	head.InsertNode(3)
	head.InsertNode(9)
	head.InsertNode(1)
	head.InsertNode(7)
	head.InsertNode(3)
	head.InsertNode(8)
	head.InsertNode(4)

	// head.printList()
	ReverseEvenLength(head)
}
