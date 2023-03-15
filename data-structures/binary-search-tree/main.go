package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	items []*Node
}

// Push function adds a value at the end off slice
func (s *Stack) Push(n *Node) {
	s.items = append(s.items, n)
}

// Pop function removes a value from the end and returns it.
func (s *Stack) Pop() *Node {
	lastIdx := len(s.items) - 1
	removedItem := s.items[lastIdx]
	s.items = s.items[:lastIdx]

	return removedItem
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Node struct
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert func
func (n *Node) Insert(k int) {
	if k > n.Key {
		// move to insert into right branch
		// If leaf node
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if k < n.Key {
		// move to insert into left branch
		// If leaf node
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search func
func (n *Node) Search(k int) bool {
	// If we reach leaf node child, we not found the key K.
	if n == nil {
		return false
	}

	if k > n.Key {
		// Move to the right branch
		return n.Right.Search(k)
	} else if k < n.Key {
		// Move to the left branch
		return n.Left.Search(k)
	}

	// found the key
	return true
}

// While traversing a tree, you need to visit three elements, root node, left subtree, and right subtree.
// The order in which you visit these three nodes determines the type of algorithms.
// Read more: https://www.java67.com/2016/07/how-to-implement-preorder-traversal-of-binary-tree-in-java.html#ixzz7twKS8J2v

// In PreOrder traversal, you visit the root or node first, followed by the left subtree and the right subtree

// In PostOrder traversal, you visit the root node at the last.

// In InOrder traversal, you visit the left subtree, the root and then the right subree.
// IMPORTANT:
// * Ont property of inorder traversal is that it prints the nodes in sorted order if the given binary tree is a binary search tree.

func (n *Node) PrintPreOrder() {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.Key)
	n.Left.PrintPreOrder()
	n.Right.PrintPreOrder()
}

func (n *Node) PrintPostOrder() {
	if n == nil {
		return
	}
	n.Left.PrintPostOrder()
	n.Right.PrintPostOrder()
	fmt.Printf("%d ", n.Key)
}

func (n *Node) PrintInOrder() {
	if n == nil {
		return
	}
	n.Left.PrintInOrder()
	fmt.Printf("%d ", n.Key)
	n.Right.PrintInOrder()
}

func (n *Node) PrintPreOrderIterative() {
	stack := Stack{}
	stack.Push(n)

	for !stack.IsEmpty() {
		currentNode := stack.Pop()
		fmt.Printf("%d ", currentNode.Key)

		if currentNode.Right != nil {
			stack.Push(currentNode.Right)
		}

		if currentNode.Left != nil {
			stack.Push(currentNode.Left)
		}
	}
}

func (n *Node) PrintInOrderIterative() {
	stack := Stack{}
	currentNode := n

	for !stack.IsEmpty() || currentNode != nil {
		if currentNode != nil {
			stack.Push(currentNode)
			currentNode = currentNode.Left
		} else {
			node := stack.Pop()
			fmt.Printf("%d ", node.Key)
			currentNode = node.Right
		}
	}
}

func (n *Node) PrintPreOrderTreeEdges(sb *strings.Builder, padding, pointer string) {
	if n == nil {
		return
	}
	sb.WriteString(padding)
	sb.WriteString(pointer)
	sb.WriteString(fmt.Sprintf("%d", n.Key))
	sb.WriteString("\n")

	var paddingBuilder strings.Builder
	paddingBuilder.WriteString(padding)
	paddingBuilder.WriteString("│  ")

	paddingForBoth := paddingBuilder.String()
	pointerForRight := "└──"
	var pointerForLeft string
	if n.Right != nil {
		pointerForLeft = "├──"
	} else {
		pointerForLeft = "└──"
	}

	n.Left.PrintPreOrderTreeEdges(sb, paddingForBoth, pointerForLeft)
	n.Right.PrintPreOrderTreeEdges(sb, paddingForBoth, pointerForRight)
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(50)

	fmt.Println(tree)

	tree.Insert(12)
	tree.Insert(103)
	tree.Insert(15)
	tree.Insert(87)
	tree.Insert(9)
	tree.Insert(240)
	tree.Insert(720)
	tree.Insert(23)
	tree.Insert(64)

	fmt.Printf("\nPrinting tree in pre-order\n")
	tree.PrintPreOrder()
	fmt.Println()

	fmt.Printf("\nPrinting tree in pre-order Iterative\n")
	tree.PrintPreOrderIterative()
	fmt.Println()

	fmt.Printf("\nPrinting tree in post-order\n")
	tree.PrintPostOrder()
	fmt.Println()

	fmt.Printf("\nPrinting tree in in-order\n")
	tree.PrintInOrder()
	fmt.Println()

	fmt.Printf("\nPrinting tree in in-order Iterative\n")
	tree.PrintInOrderIterative()
	fmt.Println()

	fmt.Printf("\nPrinting tree in pre-order with Tree Edges\n")
	var sb strings.Builder
	tree.PrintPreOrderTreeEdges(&sb, "", "")
	fmt.Println(sb.String())
	fmt.Println()

	fmt.Printf("\nSearching keys...")
	foundK := tree.Search(10)
	fmt.Printf("\nHas key=%d: %t", 10, foundK)

	foundK = tree.Search(9)
	fmt.Printf("\nHas key=%d: %t\n", 9, foundK)
}
