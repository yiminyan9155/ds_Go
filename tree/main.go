package main

import "fmt"

// tree is like a tree structure in real word but upside down
// it starts with a node called root node
// each node that has less than 2 children is called a binary tree
// each left child is less than parent and less than right child is called a binary earch tree

var count int = 0

// define a node that represents vertices of a binary search tree
type Node struct {
	middleValue int
	leftChild   *Node
	rightChild  *Node
}

// Insert
func (n *Node) Insert(val int) {
	if val < n.middleValue {
		// move left
		if n.leftChild == nil {
			n.leftChild = &Node{middleValue: val}
		} else {
			// if the left child is occupied, recursively call Insert() for the left child
			n.leftChild.Insert(val)
		}
	} else if val > n.middleValue {
		// move right
		if n.rightChild == nil {
			n.rightChild = &Node{middleValue: val}
		} else {
			n.rightChild.Insert(val)
		}
	}
}

// Search
func (n *Node) Search(val int) bool {
	count++
	if n == nil {
		return false
	}

	if val < n.middleValue {
		return n.leftChild.Search(val)
	} else if val > n.middleValue {
		return n.rightChild.Search(val)
	}
	return true
}

func main() {
	// test Node structure
	tree := &Node{middleValue: 100}

	// test inserting a node
	tree.Insert(50)
	tree.Insert(100)
	tree.Insert(300)
	fmt.Println(tree)

	// test search
	fmt.Println(tree.Search(50)) //search time - 2(depth)
	fmt.Println(count)
	fmt.Println(tree.Search(60))
}
