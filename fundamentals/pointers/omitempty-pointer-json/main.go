package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
}

// Non pointer fields are not omited.
type Omit struct {
	PersonShowed Person  `json:"person_showed,omitempty"`
	PersonHidden *Person `json:"person_hidden,omitempty"`
}

func (o *Omit) GetUnmarshaled() (Omit, error) {
	omit := Omit{}

	bytes, err := json.Marshal(omit)
	if err != nil {
		fmt.Println("Error when marshal value from bytes!")
		return omit, err
	}

	var newOmit Omit
	err = json.Unmarshal(bytes, &newOmit)
	if err != nil {
		fmt.Println("Error when unmarshal value from bytes!")
		return omit, err
	}

	return newOmit, err
}

func main() {
	omit := Omit{}

	bytes, _ := json.Marshal(omit)
	fmt.Println(string(bytes))
}
