package main

import "fmt"

/* Stack -
1. Create a stact of type []interface{}.
2. Follows Last-In-First-Out (LIFO) principle.
3. An element is pushed to the stack with the built-in append function.
4. The element is popped from the stack by slicing off the top element.
*/

// Stack interface
type Stack []interface{}

// Push into the Stack
func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
	fmt.Println("Stack :", *s, &s)
}

// IsEmpty check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Pop the item from Stack
func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	index := len(*s) - 1
	*s = (*s)[:index]
	fmt.Println("Stack After Pop:", *s)
	return
}

// func main() {
// 	var stack Stack
// 	stack.Push("this")
// 	stack.Push(12)
// 	stack.Push(42)
// 	stack.Pop()
// }
