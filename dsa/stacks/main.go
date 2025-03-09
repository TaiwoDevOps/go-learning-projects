// Stack uses the LIFO (last in first out) principle
package main

import "fmt"

// create a stack using struct

type stack struct {
	values []int // the stacked value holder
	top    int   // top of last stacked value

}

// create a new stack function that returns a new stack with the default values
func NewStack() *stack {
	return &stack{
		values: make([]int, 0),
		top:    -1,
	}
}

// you can only perform two operations on a stack
// push and pop

func (s *stack) Push(val int) {
	// to push to a stack we need to add the value to the top
	s.top++
	s.values = append(s.values, val)

}

// pop a value from the stack
func (s *stack) Pop() int {
	// to pop from a stack we need to remove the value from the top if stack is not empty
	if s.IsEmpty() {
		return -1
	}
	val := s.values[s.top]
	s.top--
	return val
}

// check if the stack is empty
func (s *stack) IsEmpty() bool {
	return s.top == -1
}

func main() {

	// create a new stack
	s := NewStack()

	// push some values to the stack
	s.Push(1)
	s.Push(3)
	s.Push(2)
	fmt.Println("the values in stacks: ", s.values)
	// pop a value from the stack
	fmt.Println(s.Pop())

	// check if the stack is empty
	fmt.Println(s.IsEmpty())
	// remaining values in the stack
	fmt.Println(s.values)

}
