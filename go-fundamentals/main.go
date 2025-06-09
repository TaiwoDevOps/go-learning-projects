package main

import (
	"fmt"
	// arrSlices "github.com/taiwodevops/go-fundamentals/arrays-slices"
)

func main() {

	// var numbers [6]int

	// var slices []int

	// slices = numbers[2:4]

	// fmt.Println(slices)      // [0 0]
	// fmt.Println(numbers)     // [0 0 0 0 0 0]
	// fmt.Println(len(slices)) // 2
	// fmt.Println(cap(slices)) // 4

	// slices = slices[0:4]
	// fmt.Println(slices) // [0 0 0 0]
	// slices[3] = 10
	// fmt.Println(slices) // [0 0 0 10]

	// slices = numbers[:]
	// fmt.Println(slices) // [0 0 0 0 0 10]

	// slice2 := slices[:]

	// numbers[2] = 20
	// fmt.Println(slice2)  // [0 0 20 0 0 10]
	// fmt.Println(numbers) // [0 0 20 0 0 10]
	// fmt.Println(slices)  // [0 0 20 0 0 10]

	// slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(slice)      // [1 2 3 4 5 6 7 8 9 10]
	// fmt.Println(len(slice)) // 10
	// arrSlices.SubtractOneFromLength(&slice)

	// arrSlices.SubtractOneFromLength(&slice)
	// fmt.Println(slice) // [1 2 3 4 5 6 7 8]

	// newS := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// var sl2 []*int
	// for _, n := range newS {
	// 	sl2 = append(sl2, &n)
	// }

	// fmt.Println(sl2)

	// for _, n := range sl2 {
	// 	fmt.Println(*n) // 10
	// }

	s := make([]int, 0, 3)
	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("cap %v, len %v, %p\n", cap(s), len(s), s)
	}

}
