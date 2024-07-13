package main

import "fmt"

const (
	ARRAY_SIZE = 7
)

// hash table structure
type HashTable struct {
	arr [ARRAY_SIZE]*Bucket
}

// Insert
func (h *HashTable) Insert(key string) {
	idx := hash(key)
	h.arr[idx].insert(key)
}

// Search
func (h *HashTable) Search(key string) bool {
	idx := hash(key)
	return h.arr[idx].search(key)
}

// Delete
func (h *HashTable) Delete(key string) {
	idx := hash(key)
	h.arr[idx].delete(key)
}

// bucket structi
type Bucket struct {
	head *BucketNode
}

// bucket node struct
type BucketNode struct {
	key  string
	next *BucketNode
}

// insert
func (b *Bucket) insert(k string) {
	if !b.search(k) {
		newNode := &BucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Printf("\nkey %s already exists!\n", k)
	}
}

// search
func (b *Bucket) search(key string) bool {
	curNode := b.head
	for curNode != nil {
		if curNode.key == key {
			return true
		}
		curNode = curNode.next
	}
	return false
}

// delete
func (b *Bucket) delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}
	prevNode := b.head
	for prevNode.next != nil {
		if prevNode.next.key == key {
			// delete current node
			prevNode.next = prevNode.next.next
			return
		}
		prevNode = prevNode.next
	}
}

// hash function
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ARRAY_SIZE
}

// init
func Init() *HashTable {
	res := &HashTable{}
	for i := range res.arr {
		res.arr[i] = &Bucket{}
	}
	return res
}

func main() {
	hTable := Init()

	list := []string{
		"CARLO",
		"CARLA",
		"MAC",
		"PC1",
		"CARLE",
	}

	for _, v := range list {
		hTable.Insert(v)
	}
}
