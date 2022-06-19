package main

import "fmt"

// trie is a special case of tree structure
// it has a root node that contains 26 nodes for 26 index(latin alphabet)
// each node then has 26 index, and the node that represents the child contains address
// each node has a bool value isEnd value to indicate whether if it's the last letter of the word

const AlphabetSize = 26

// Trie structure points to the root of trie
type Trie struct {
	root *Node
}

// Node
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Insert a new word to the trie
func (t *Trie) InsertWord(word string) {
	currentNode := t.root
	for i := 0; i < len(word); i++ {
		charIndex := word[i] - 'a'
		// if it doesn't exist, create a node for it
		if charIndex <= 27 && currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	// if we reach the end of word
	currentNode.isEnd = true
}

// Seach checks if a word exists in a trie
func (t *Trie) SearchWord(word string) bool {
	currentNode := t.root
	for i := 0; i < len(word); i++ {
		charIndex := word[i] - 'a'
		// if it doesn't exist, it indicates that the word doesn't exist in the trie
		// catch index out of bound errors if other alphabets are used: (e.g. Caps)
		if charIndex > 26 || currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	// if we reach the end of word
	if currentNode.isEnd == true {
		return true
	}
	return false
}

// initialize a trie structure
func Init() *Trie {
	trie := &Trie{root: &Node{}}

	return trie
}

func main() {
	testTrie := Init()
	testTrie.InsertWord("hi")
	fmt.Println(testTrie.root)
	fmt.Println(testTrie.SearchWord("hi"))
	fmt.Println(testTrie.SearchWord("HI"))

}
