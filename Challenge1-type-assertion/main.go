package main

import (
	"fmt"
)

// Objective is to implement type assertion 

type Developer struct {
	Name string
	Age  int
}

func GetDeveloper(name interface{}, age interface{}) Developer {
	var dev Developer

	// Type assertion to convert interface{} to string and int
	dev.Name = name.(string)
	dev.Age = age.(int)

	return dev
}

func main() {
	fmt.Println("Hello World")

	var name interface{} = "Elliot"
	var age interface{} = 26

	dynamicDev := GetDeveloper(name, age)
	fmt.Println(dynamicDev.Name)
}
