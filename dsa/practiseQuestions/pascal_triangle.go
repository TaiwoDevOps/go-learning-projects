package practiceQuestions

import "fmt"

// https://leetcode.com/problems/pascals-triangle/description/

// Pascal triangle :
// if give an number say : 5
// then return : [[1], [1,1], [1,2,1], [1,3,3,1], [1,4,6,4,1]]
// with indexes    1,    2,      3,        4,          5      respectively
// and the value at index 0 is always 1
// the value at index 1 is always 1
// so it is index plus 1 for getting the size of the current row/ array length in the triangle
func Generate(numRows int) [][]int {
	var subArray [][]int
	if numRows == 1 {
		subArray = [][]int{{1}}
		return subArray
	}

	if numRows == 2 {
		subArray = [][]int{{1}, {1, 1}}
		return subArray
	}

	for i := 0; i < numRows; i++ {
		// create the sub arrays based on the current index + 1
		// and the rest of the elements in the sub array to the sum of the previous two elements in the sub array
		subArray = append(subArray, []int{1})
		for j := 1; j < i; j++ {
			fmt.Printf("sub @ i %d sub @ i -1 %v \n", subArray[i], subArray[i-1][j-1]+subArray[i-1][j])

			subArray[i] = append(subArray[i], subArray[i-1][j-1]+subArray[i-1][j])
		}
		subArray[i] = append(subArray[i], 1)

	}

	return subArray

}
