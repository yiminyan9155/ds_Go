package main

import "fmt"

type LinkedList struct {
	length int
	head   *Node
}

type Node struct {
	value int
	next  *Node
}

// prepend adds the value to the beginning of the list which becomes the new head
func (l *LinkedList) prepend(newNode *Node) {
	// stores old head
	secondNode := l.head
	// makes the new node new head
	l.head = newNode
	// makes old head the next of new head
	l.head.next = secondNode
	l.length++
}

// pintNode prints out all the nodes in the linked list
func (l LinkedList) printNodes() {
	nodeToPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", nodeToPrint.value)
		nodeToPrint = nodeToPrint.next
		l.length--
	}
	fmt.Println("\n")
}

// deleteValue deletes the specified value from the list
// we need to always stay in the previous node to handle/delete the next node as if we delete the current one, we lost the address of chain of list
// and we can't come back the previous node as this is single linked list
func (l *LinkedList) deleteValue(valueToDelete int) {
	// check invalid input
	if l.length == 0 {
		return
	}

	// if the head is the value to delete
	if valueToDelete == l.head.value {
		l.head = l.head.next
		l.length--
		return
	}

	preToDeleteNode := l.head
	for preToDeleteNode.next.value != valueToDelete {
		// if it reaches the end of list and still didn't find the value
		if preToDeleteNode.next.next == nil {
			fmt.Println("value to delete doesn't exist in the list")
			return
		}
		preToDeleteNode = preToDeleteNode.next
	}
	// delete the node by skipping the node in the chain
	preToDeleteNode.next = preToDeleteNode.next.next
	l.length--

}

func main() {
	testLinkedList := LinkedList{}
	node1 := &Node{value: 0}
	node2 := &Node{value: 1}
	node3 := &Node{value: 2}
	node4 := &Node{value: 3}
	node5 := &Node{value: 4}
	node6 := &Node{value: 5}
	testLinkedList.prepend(node1)
	testLinkedList.prepend(node2)
	testLinkedList.prepend(node3)
	testLinkedList.prepend(node4)
	testLinkedList.prepend(node5)
	testLinkedList.prepend(node6)
	fmt.Printf("linkedlist content: %v\n", testLinkedList)

	// test printNodes
	testLinkedList.printNodes()

	// test deleteValue
	// testLinkedList.deleteValue(3)
	// testLinkedList.printNodes()

	// test if the head is the value to delete
	// testLinkedList.deleteValue(5)
	// testLinkedList.printNodes()

	// test if the value doesn't exist in the list
	testLinkedList.deleteValue(10)
	testLinkedList.printNodes()
}
