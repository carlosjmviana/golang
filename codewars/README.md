

# CodeWars Training

## String ends with?

```
Complete the solution so that it returns true if the first argument(string) passed in ends with the 2nd argument (also a string).

Examples:

solution("abc", "bc") // returns true
solution("abc", "d") // returns false
```

### Solution

```
package kata

func solution(str, ending string) bool {
  if len(ending) > len(str) {
    return false
  }
  slice := str[len(str) - len(ending):]
  
  return slice == ending
}
```

## Bit Counting

```
Write a function that takes an integer as input, and returns the number of bits that are equal to one in the binary representation of that number. You can guarantee that input is non-negative.

Example: The binary representation of 1234 is 10011010010, so the function should return 5 in this case

```

### Solution

```
import (
  "strconv"
)
func CountBits(n uint) int {
  binary := strconv.FormatInt(int64(n), 2)
  
  count := 0
  for _,b := range binary {
    if b == '1' {
      count++
    }
  }
  return count
}
```

plan b:
```
func CountBits(i uint) int {
  res := 0
  for i > 0 {
    if i & 1 == 1 {
      res ++
    }
    i = i >> 1
  }
  return res
}
```

#### Other solutions

```
func CountBits(n uint) int {
  var res int = 0
  for (n>0) {
    if (n & 1 == 1) {
      res = res + 1
    }
    n = n >> 1
  }
  return res
}
```

```
import "math/bits"

func CountBits(n uint) int {
  return bits.OnesCount(n)
}
```
