package main

import "fmt"

func main() {

	res := twoSumData([]int{2, 7, 11, 15}, 22)

	fmt.Println(res)

}

func twoSumData(numbers []int, target int) []int {

	// create a map to the value and its index as k,v
	valueMap := make(map[int]int)

	// iterate based on the len of the array
	for i := 0; i < len(numbers); i++ {
		subVal := target - numbers[i]

		if j, value := valueMap[subVal]; value {
			return []int{j, i}
		}

		valueMap[numbers[i]] = i
	}

	return []int{}
}
