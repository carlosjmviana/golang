package main

import "fmt"

type node struct {
  value int
  next *node
}

type linkedList struct {
  head *node
  length int
}

func (l *linkedList) prepend (n *node) {
  second := l.head 
  l.head = n
  l.head.next = second
  l.length += 1
}

func (l linkedList) printData()  {
  n := l.head
  for n != nil {
    fmt.Printf("%d ", n.value)
    n = n.next
  }
  fmt.Printf("\nLength: %d", l.length)
}

func (l *linkedList) deleteNode(v int) {
  // empty list
  if l.head == nil {
    return
  }

  // Remove first element
  if l.head.value == v {
    l.head = l.head.next
    l.length--
    return 
  }

  current := l.head
  previous := current 
  for current.value != v {
    // check for element not in list
    if previous.next.next == nil {
      return
    }
    previous = current
    current = current.next
  }

  // remove intermediate element
  // remove last element
  previous.next = current.next
  current.next = nil
  l.length--
}

func main () {
  fmt.Println("Main package")  

  myList := linkedList{}
  n1 := &node{value: 30}
  n2 := &node{value: 10}
  n3 := &node{value: 40}
  n4 := &node{value: 18}
  n5 := &node{value: 23}
  n6 := &node{value: 91}

  myList.prepend(n1)
  myList.prepend(n2)
  myList.prepend(n3)
  myList.prepend(n4)
  myList.prepend(n5)
  myList.prepend(n6)
  
  fmt.Printf("\n\nMyList\n")
  myList.printData()
  
  fmt.Printf("\n\nRemoving element not in list: 100\n")
  myList.deleteNode(100)
  myList.printData()

  fmt.Printf("\n\nRemoving intermediate element with value 18\n")
  myList.deleteNode(18)
  myList.printData()
  
  fmt.Printf("\n\nRemoving last element\n")
  myList.deleteNode(30)
  myList.printData()
  
  fmt.Printf("\n\nRemoving first element\n")
  myList.deleteNode(myList.head.value)
  myList.printData()
  
  fmt.Printf("\n\nRemoving all elements")
  myList.deleteNode(myList.head.value)
  myList.deleteNode(myList.head.value)
  myList.deleteNode(myList.head.value)
  myList.printData()
  
  fmt.Printf("\n\nEmpty MyList")
  emptyList := linkedList{}
  emptyList.printData()
  fmt.Printf("\nRemoving from empty list")
  emptyList.deleteNode(0)
  emptyList.printData()

  fmt.Printf("\n")
}
