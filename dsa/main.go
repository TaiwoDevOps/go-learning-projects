package main

import (
	"fmt"

	practiceQuestions "github.com/TaiwoDevOps/dsa/practiseQuestions"
)

func main() {

	// res := twoSumData([]int{2, 7, 11, 15}, 22)

	// fmt.Println(res)

	// nums := []int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9} //{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9} //{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// val := practiceQuestions.RemoveDuplicates(nums)
	// fmt.Println("the length is ", val)

	arRes := practiceQuestions.Generate(5)
	fmt.Println("the generated array is ", arRes)

}

/*

single string swap algorithm

package main

import "../practise_questions/fmt"

func main() {
	res := AreAlmostEqual("bank", "kanb")
	fmt.Println("res:-->", res)

	reverse even length
	head := new(ListNode)

	[5,2,6,3,9,1,7,3,8,4]
	head.InsertNode(2)
	head.InsertNode(5)
	head.InsertNode(6)
	head.InsertNode(3)
	head.InsertNode(9)
	head.InsertNode(1)
	head.InsertNode(7)
	head.InsertNode(3)
	head.InsertNode(8)
	head.InsertNode(4)

	head.printList()
	ReverseEvenLength(head)



}

*/
