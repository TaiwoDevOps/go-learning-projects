package arrSlices

import "fmt"

// How slices work in Go
// Slices are a reference type, meaning they point to the underlying array.
// When you create a slice, it does not create a new array; instead, it creates a reference to an existing array.
// This means that if you modify the slice, you are modifying the underlying array.
// Slices have a length and a capacity.
// The length is the number of elements in the slice, and the capacity is the number of elements in the underlying array that can be accessed through the slice.
// When you create a slice, you can specify the length and capacity.
// If you do not specify the length and capacity, the slice will have a length of 0 and a capacity of the length of the underlying array.
// Slices can be created using the `make` function, which allows you to specify the length and capacity of the slice.
// Slices can also be created using the `[]` operator, which creates a slice from an existing array.
// Slices can be resized using the `append` function, which adds elements to the end of the slice.
// Slices can also be resized using the `copy` function, which copies elements from one slice to another.
// Slices can be compared using the `==` operator, which compares the elements of the slices.
// Slices can be iterated over using the `for` loop, which allows you to access each element of the slice.
// Slices can be passed to functions as arguments, which allows you to modify the underlying array.
// Slices can be returned from functions, which allows you to return a reference to the underlying array.
// Slices can be used to create multidimensional arrays, which allows you to create arrays of arrays.
// Slices can be used to create dynamic arrays, which allows you to create arrays that can grow and shrink in size.

func SubtractOneFromLength(s *[]int) {
	sl := *s
	sl = sl[:len(sl)-1]
	// put it back to the original slice
	*s = sl

}

func Helper(arr [3]int) {
	fmt.Printf("the array here %d", arr)
}
