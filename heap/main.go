package main

import "fmt"

// a heap is a complete tree(all levels should be complete except for the lowest)
// min heap has smallest value on the top and values grow when we traverse down
// max heap has largest value on the top and values decrease when we travserse down

// heap in nature is stored in array
// left child of parent is (parentIndex * 2 + 1)
// right child of parent is (parentIndex * 2 + 2)

// Heap structure
type MaxHeap struct {
	heap []int
}

// Insert number
func (h *MaxHeap) InsertNum(num int) {
	h.heap = append(h.heap, num)
	h.maxHeapifyUp(len(h.heap) - 1)
}

// maxHeapifyUp when there's an insertion, as the element is inserted to the end,
// it should heapify upwards
func (h *MaxHeap) maxHeapifyUp(index int) {
	// if the added element is larger than its parent, swap and iterate until it reachses to the root
	for h.heap[index] > h.heap[getParentIndex(index)] {
		h.swap(getParentIndex(index), index)
		index = getParentIndex(index)
	}
}

// extract max
func (h *MaxHeap) ExtractMax() int {
	if len(h.heap) != 0 {
		max := h.heap[0]
		// move the last element of heap to be the new root
		h.heap[0] = h.heap[len(h.heap)-1]
		h.heap = h.heap[:len(h.heap)-1]

		h.maxHeapifyDown(0)
		return max
	}
	fmt.Println("heap is empty")
	return -1
}

// maxHeapifyDown when there's a max removal downwards
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.heap) - 1
	leftIndex, rightIndex := getLeftChildIndex(index), getRightChildIndex(index)
	childToCompareIndex := 0

	// loop as long as it has at least one child because this is a complete tree that fills from left
	for leftIndex <= lastIndex {
		// when the left child is the only child
		// when left child is larger
		if leftIndex == lastIndex || h.heap[leftIndex] > h.heap[rightIndex] {
			childToCompareIndex = leftIndex
		} else {
			// when right child is larger
			childToCompareIndex = rightIndex
		}

		if h.heap[index] < h.heap[childToCompareIndex] {
			h.swap(index, childToCompareIndex)
			index = childToCompareIndex
			leftIndex, rightIndex = getLeftChildIndex(index), getRightChildIndex(index)
		} else {
			return
		}
	}
}

// get parent index
func getParentIndex(childIndex int) int {
	// it doesn't matter left or right child, it will round down as it's int division
	parentIndex := (childIndex - 1) / 2
	return parentIndex
}

// get child index
func getLeftChildIndex(parentIndex int) int {
	leftChildIndex := (parentIndex*2 + 1)
	return leftChildIndex
}
func getRightChildIndex(parentIndex int) int {
	rightChildIndex := (parentIndex*2 + 2)
	return rightChildIndex
}

// swap
func (h *MaxHeap) swap(index1, index2 int) {
	h.heap[index1], h.heap[index2] = h.heap[index2], h.heap[index1]
}

func main() {
	testHeap := &MaxHeap{}

	//test inserting
	testHeap.InsertNum(2)
	testHeap.InsertNum(30)
	testHeap.InsertNum(10)
	testHeap.InsertNum(11)
	// testHeap.InsertNum(23)
	// testHeap.InsertNum(400)
	// testHeap.InsertNum(39)
	// testHeap.InsertNum(15)
	// testHeap.InsertNum(3)
	fmt.Println(testHeap)

	// test heapify up
	testHeap.InsertNum(100)
	fmt.Println(testHeap)

	// test heapify down and extract
	fmt.Println("extracting")
	testHeap.ExtractMax()
	fmt.Println(testHeap)
}
