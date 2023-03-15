package main

import "fmt"

const ALPHABET_SIZE = 26

// Node Struct
type Node struct {
	children [ALPHABET_SIZE]*Node
	isEnd    bool
}

// Trie
type Trie struct {
	root *Node
}

// Build Trie
func NewTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert a word
// 'a' represents the value 97 decimal. We will usew w[i] - 'a' to set the index into children array
func (t *Trie) Insert(w string) {
	wLength := len(w)
	currentNode := t.root
	for i := 0; i < wLength; i++ {
		childrenIdx := w[i] - 'a'
		if currentNode.children[childrenIdx] == nil {
			currentNode.children[childrenIdx] = &Node{}
		}
		currentNode = currentNode.children[childrenIdx]
	}
	currentNode.isEnd = true
}

// Search
func (t *Trie) Search(w string) bool {
	wLength := len(w)
	currentNode := t.root
	for i := 0; i < wLength; i++ {
		childrenIdx := w[i] - 'a'
		if currentNode.children[childrenIdx] == nil {
			return false
		}
		currentNode = currentNode.children[childrenIdx]
	}

	return currentNode.isEnd
}

// Delete
func (t *Trie) Delete(w string) {
	wLength := len(w)
	currentNode := t.root
	for i := 0; i < wLength; i++ {
		childrenIdx := w[i] - 'a'
		if currentNode.children[childrenIdx] == nil {
			return
		}
		currentNode = currentNode.children[childrenIdx]
	}
	// Found w word
	currentNode.isEnd = false
}

func main() {
	trie := NewTrie()

	fmt.Printf("\nEmtpy trie\n")
	fmt.Println(trie.root)

	words := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}

	fmt.Println("\nInserting words...")
	for _, w := range words {
		trie.Insert(w)
	}
	fmt.Println("Print root node")
	fmt.Println(trie.root)

	ws := "oregano"
	fmt.Printf("\nSearch for word %s: %t\n", ws, trie.Search(ws))

	fmt.Printf("\nRemoving word: %s", ws)
	trie.Delete(ws)
	fmt.Printf("\nSearch for world %s: %t\n", ws, trie.Search(ws))
}
