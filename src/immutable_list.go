package main

import (
	"errors"
	"fmt"
)

// LList represents an immutable list
// LList represents an immutable list
type LList[T any] interface {
    isEmpty() bool
    head() (T, error)
    tail() (LList[T], error)
    length() int
    copy() LList[T]
    append(LList[T]) LList[T]
    transform(func(interface{}) interface{}) LList[T]
}

// Nil represents an empty list
type Nil[T any] struct{}

func (Nil[T]) isEmpty() bool {
    return true
}

func (Nil[T]) head() (T, error) {
    var zeroValue T
    return zeroValue, errors.New("head of empty list")
}


func (Nil[T]) tail() (LList[T], error) {
    return nil, errors.New("tail of empty list")
}

func (Nil[T]) length() int {
    return 0
}

func (Nil[T]) copy() LList[T] {
    return Nil[T]{}
}

func (Nil[T]) append(list LList[T]) LList[T] {
    return list
}

func (Nil[T]) transform(f func(interface{}) interface{}) LList[T] {
    return Nil[T]{}
}


// Examples
func cons() {

    // Create an empty list
    nilList := Nil[int]{}
    fmt.Println("Is nilList empty?", nilList.isEmpty())

    // Create a non-empty list
    list := Cons[int]{1, Cons[int]{2, Nil[int]{}}}
    fmt.Println("Is list empty?", list.isEmpty())

    // Get the head of the list
    head, _ := list.head()
    fmt.Println("Head of list:", head)

    // Get the tail of the list
    tail, _ := list.tail()
    fmt.Println("Tail of list:", tail)

    // Get the length of the list
    fmt.Println("Length of list:", list.length())

    // Copy the list
    copiedList := list.copy()
    fmt.Println("Is copiedList empty?", copiedList.isEmpty())

    // Append two lists
    appendedList := list.append(Cons[int]{3, Cons[int]{4, Nil[int]{}}})
    fmt.Println("Appended list:", appendedList)

    // Transform over the lis
	// Define the transformation function
	transformFunc := func(x interface{}) interface{} {
    // Type assert x to int
  	  value, ok := x.(int)
 	 	  if !ok {
        // Handle the case where the type assertion fails
        // For example, you might return an error or panic
	    }

	    // Perform the transformation
	    return value * 2
	}

	// Apply the transformation function to the list
	transformedList := list.transform(transformFunc)

    fmt.Println("Transformed list:", transformedList)
}
