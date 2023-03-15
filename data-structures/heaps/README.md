# Data Structures and Algorithms in Go - Heaps

The heap data structured was introduced as a part of the algorithm Heapsort in Feb. 1964. Heaps are used to implement other data structures such as _Priority Queues_

## Heaps Definition
It is a Data Structure that can be expressed as a complete tree. A complete tree is a tree where all levels, except the leaf, are full.

### Types of Heap

Max Heap: The max value is in the root level. 

Min Heap: The min value is in the root level. Each level has a key that is greatter than the key value of the level before.

Max and Min Heaps are implemented as arrays. We can find out the index of a node using the following calculation:

* Left child index: <parent index> x 2 + 1
* Right child index: <parent index> x 2 + 2

### Insert into a Heap

**MaxHeapInsert**

We put the new value at the end of the array and **Hipify** the value. 

### Extract Max

**HeapExtractMax**

We remove the max value from root, get the last value and put it on root. After that, we compare the root with its childs and do this comparison 
until to get leaf nodes. 

### Time Complexity

O(h) where h is the height of the three. O(log n)
