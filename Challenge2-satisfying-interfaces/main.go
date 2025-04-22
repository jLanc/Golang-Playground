package main

import (
	"fmt"
)

// Objective is to implement the age funciton for employee interface
// Also extend employee interface with something else



type Employee interface {
    Language() string
	Height() string
    Age() int
}

type Engineer struct {
    Name string
}

func (e Engineer) Language() string {
    return e.Name + " programs in Go"
}

func (e Engineer) Height() string {
	return e.Name + "and is 183cm tall"
}

func (e Engineer) Age() int {
	return 30
}

func main() {
    // This will throw an error
    var programmers []Employee
    elliot := Engineer{Name: "Elliot"}
    // Engineer does not implement the Employee interface
    // you'll need to implement Age() and Random()
    programmers = append(programmers, elliot)

	fmt.Println(programmers[0].Language())
	fmt.Println(programmers[0].Age())
	fmt.Println(programmers[0].Height())
}
