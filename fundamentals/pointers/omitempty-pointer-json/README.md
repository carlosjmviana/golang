# Pointers As Omit Fields

In nested structures we can omit a field when marshaling/unmarshaling if we use
pointers to that structure.

The code example in (main)[./main.go] file shows a struct Omit with 2 fields with json omitempty property.
The second field is trully omited because we are using pointer to Person struct. 


