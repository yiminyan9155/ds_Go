package main

import "fmt"

// stack follows last in first out
// think about how you remove the top book first from a stack of books

// Stack is a stack that holds a slice
type Stack struct {
	items  []int
	length int
}

// Push appends an element to the top of the stack
func (stack *Stack) Push(item int) {
	stack.items = append(stack.items, item)
	stack.length++
}

// Pop
func (stack *Stack) Pop() {
	// error check
	if stack.length > 0 {
		stack.items = stack.items[:stack.length-1]
		stack.length--
	} else {
		fmt.Println("Can't pop items from empty stack")
	}

}

// main function to test code
func main() {
	// test Push
	testStack := Stack{}
	fmt.Println(testStack)
	testStack.Push(0)
	testStack.Push(1)
	testStack.Push(2)
	fmt.Printf("stack content: %v\n", testStack.items)
	fmt.Printf("length of stack: %v\n", testStack.length)

	// test Pop
	testStack.Pop()
	fmt.Printf("stack content: %v\n", testStack.items)
	fmt.Printf("length of stack: %v\n", testStack.length)
	testStack.Pop()
	testStack.Pop()

	// test invalid input
	testStack.Pop()
}
