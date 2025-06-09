package practiceQuestions

import "fmt"

func RemoveDuplicates(nums []int) int {

	var pointer int = 1

	for i := pointer; i < len(nums); i++ {
		// at index 1, check if the value at index 1 is not equal to the value at index 0

		fmt.Println("the  pointer length ", pointer)
		fmt.Println("the nums index ", i)

		//                    i
		//                    i
		//              p
		// {1, 2, 3, 4, 3, 4, 5, 6, 7, 8, 9}
		if nums[i] != nums[i-1] {
			fmt.Println("yes ", pointer)

			// if it is not, increment the pointer and assign the value at index i to the value at the pointer
			nums[pointer] = nums[i]
			fmt.Println("changed", nums[pointer])
			fmt.Println("num", nums)

			pointer++
		}

	}
	return len(nums[:pointer])
}
