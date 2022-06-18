package main

import "fmt"

const tableSize = 8

// HashTable - the main table that holds keys and index <- an array
type HashTable struct {
	table [tableSize]*bucket
}

// Bucket - a bucket that holds several bucket nodes <- a linked list
type bucket struct {
	head *bucketNode
}

// Bucket Node <- a node that holds value <- a node of linked list
type bucketNode struct {
	key  string
	next *bucketNode
}

// Hash function - hash input/key into an index to access a bucket
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % tableSize
}

// Insert a bucket
func (ht *HashTable) InsertBucket(key string) {
	tableIndex := hash(key)
	ht.table[tableIndex].insertBucketNode(key)
}

// Search a bucket
func (ht *HashTable) SearchBucket(key string) bool {
	tableIndex := hash(key)
	return ht.table[tableIndex].searchBucketNode(key)
}

// Delete a bucket
func (ht *HashTable) DeleteBucket(key string) {
	tableIndex := hash(key)
	if ht.table[tableIndex] != nil {
		ht.table[tableIndex].deleteBucketNode(key)
	}
}

// insert a bucket node
func (b *bucket) insertBucketNode(k string) {
	if !b.searchBucketNode(k) {
		newBucketNode := &bucketNode{key: k}
		newBucketNode.next = b.head
		b.head = newBucketNode
	} else {
		fmt.Println("item already exists")
	}
}

// search a bucket node
func (b *bucket) searchBucketNode(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete a bucket node
func (b *bucket) deleteBucketNode(k string) {
	// if head is the match, make the next node new head
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	// if head isn't the match, traverse through the list to find the match
	preToCurrentNode := b.head
	for preToCurrentNode.next != nil {
		if preToCurrentNode.next.key == k {
			preToCurrentNode.next = preToCurrentNode.next.next
			return
		}
		preToCurrentNode = preToCurrentNode.next
	}
	fmt.Errorf("item to delete doesn't exist")
}

// Initialize a hashtable
func Init() *HashTable {
	hashTable := &HashTable{}
	for i := range hashTable.table {
		hashTable.table[i] = &bucket{}
	}
	return hashTable
}

func main() {

	// test hashtable structure
	testHashTable := &HashTable{}
	fmt.Println(testHashTable)

	// test init function
	testHashTable2 := Init()
	fmt.Println(testHashTable2)

	// test hash
	testHash := "Rhonda"
	testHash1 := "LiamLoveRhonda"
	fmt.Println(hash(testHash))
	fmt.Println(hash(testHash1))

	// test insertBucketNode to bucket
	testBucket := &bucket{}
	testBucket.insertBucketNode("Rhonda love")
	fmt.Println(testBucket.head.key)
	testBucket.insertBucketNode("Liam")
	fmt.Println(testBucket.head.key)

	// test insertBucket table
	testHashTable2.InsertBucket("Sunny is Rhonda")
	testHashTable2.InsertBucket("Liam")
	// test duplicate insert
	testHashTable2.InsertBucket("Liam")

	// test search
	fmt.Println(testHashTable2.SearchBucket("Sunny is Rhonda"))
	fmt.Println(testHashTable2.SearchBucket("Rhonda is Sunny"))
	fmt.Println(testHashTable2.SearchBucket("Liam"))

	// test delete
	testHashTable2.DeleteBucket("Sunny is Rhonda")
	fmt.Println("after deleting")
	fmt.Println(testHashTable2.SearchBucket("Sunny is Rhonda"))

	// test empty table
	testEmptyTable := &HashTable{}
	testEmptyTable.DeleteBucket("hi")
	fmt.Println(testHashTable2.SearchBucket("hi"))
}
