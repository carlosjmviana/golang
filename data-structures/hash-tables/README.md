# Hash Tables

The idea is to build an index for the value and access the structure (array) fast. We use a hash function to compute the index (hash code) from the
value that will be inserted on the hash table.


## Hash Function

### Colision 

When we apply a hash function to a two objects, names, we can have the same hash code as the hash function output. This is called hash colision. Collision handeling is how do we deal with collisions.
Collisions can be in some types:

#### Open Addressing

 First inserted 'carlo' in hash table
 | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 
 | ------ | ------ | ------ | ------ | ------ | ------ | ------ |
 |  |  |  |  |  carlo | | |

 After, we insert `carla` and both words has the same `hash code` value 4.

Where to insert `carla` word ???
 | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 
 | ------ | ------ | ------ | ------ | ------ | ------ | ------ |
 |  |  |  |  |  carlo | | |


  We put the value in the next position of hash code.

 | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 
 | ------ | ------ | ------ | ------ | ------ | ------ | ------ |
 |  |  |  |  |  carlo | carla | |

 To search for `carla` we ask first for the hash code positon and walk each array position until to find the expected value

hasFunction(`carla`) = 4

To search `carla` into hash table, we first calculate its hashFunction value and than asks for `carla`. If equals stop and return, else search in next position. 
```javascript
array[4] == 'carla' ?   
array[5] == 'carla' ?   
```

                    start --> end 
 | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 
 | ------ | ------ | ------ | ------ | ------ | ------ | ------ |
 |  |  |  |  |  carlo | carla | |

This solution is better than walk all array to find the specific value. Otherwise, as we increase the number of collisions, the worst case would be something like that:

```bash
hasFunction(carla) = 4
hasFunction(berlin) = 4
hasFunction(cebola) = 4
hasFunction(pepper) = 4
hasFunction(rice) = 4
hasFunction(alho) = 4
```

    -----------------------> end   start --------------->
 | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 
 | ------ | ------ | ------ | ------ | ------ | ------ | ------ |
 | cebola | pepper  | rice  | alho  |  carlo | carla | berlin |

Then to search the word `alho` we must walk every array position, starting by index 4.

#### Separate Chaining

Store multiple values in the same array position. We use `linked list` to store multiple values in the same position.

[Implementation](./main.go) 
