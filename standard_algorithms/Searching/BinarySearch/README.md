# Binary Search Algorithm

## Problem Statement
Given a sorted array of n integers and a value, determine if the value exists in the array 
in logarithmic time. Print the value index if it exists in the array.

## Implementation

### Iterative
[Iterative Version](./iterative-binary-search.go)

### Recursive
[Recursive Version](./recursive-binary-search.go)

## Time Complexity

In each step, the search space reduces to half. 

step 1:  $n$   
step 2: $n \over 2$  
step 3: $n \over 4$  
...  
step n:  1  


Suppose after $s$ steps we reduce the search space to 1 then:

  ${n \over 2^s} = 1$  
  $n = 2^s$  
  $s = log_2^n$  

Therefore, time complexity is $O(log_2 n)$.
The space complexity is $O(1)$ for iterative version and $O(log_2 n)$ for recursive version because of stack calls.
