package main

import "fmt"

// queue follows first come first serve
// think about when you wait in line at Starbucks, they serve in the order of queued

// represents a queue that holds a slice
type Queue struct {
	items  []int
	length int
}

// Enqueue puts in a new item to the end
func (queue *Queue) Enqueue(item int) {
	queue.items = append(queue.items, item)
	queue.length++
}

// Dequeue removes an item from the beginning
func (queue *Queue) Dequeue() {
	if queue.length > 0 {
		queue.items = queue.items[:queue.length-1]
		queue.length--
	} else {
		fmt.Println("Can't Dequeue items from empty queue")
	}

}

// Test implementations
func main() {
	// test Enqueue
	testQueue := Queue{}
	testQueue.Enqueue(1)
	testQueue.Enqueue(2)
	testQueue.Enqueue(3)
	fmt.Printf("queue content: %v\n", testQueue.items)
	fmt.Printf("queue length: %v\n", testQueue.length)

	// test Dequeue
	testQueue.Dequeue()
	fmt.Printf("queue content: %v\n", testQueue.items)
	fmt.Printf("queue length: %v\n", testQueue.length)
	testQueue.Dequeue()
	testQueue.Dequeue()

	// test invalid input
	testQueue.Dequeue()
	fmt.Printf("queue content: %v\n", testQueue.items)
}
