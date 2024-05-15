package main

import "fmt"

// Define an interface representing the structure of a Person
type Person interface {
    getName() string
    getAge() int
}

// Define a function that takes any type that matches the Person structure
func printPersonInfo(p Person) {
    fmt.Println("Name:", p.getName())
    fmt.Println("Age:", p.getAge())
}

// Define a struct that satisfies the Person interface
type SimplePerson struct {
    name string
    age  int
}

// Implement methods to satisfy the Person interface for SimplePerson
func (sp SimplePerson) getName() string {
    return sp.name
}

func (sp SimplePerson) getAge() int {
    return sp.age
}

// Define another struct with the same structure but different name
type AnotherPerson struct {
    name string
    age  int
}

// Implement methods to satisfy the Person interface for AnotherPerson
func (ap AnotherPerson) getName() string {
    return ap.name
}

func (ap AnotherPerson) getAge() int {
    return ap.age
}

func xample_typing_1() {
    // Create instances of SimplePerson and AnotherPerson
    simple := SimplePerson{name: "Alice", age: 30}
    another := AnotherPerson{name: "Bob", age: 35}

    // Call the function with instances of different types but matching structure
    printPersonInfo(simple)
    printPersonInfo(another)
}
