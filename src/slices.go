package main

import "fmt"

func slices_xample_1() {
	// Creation if an emoty slice
    var numbers []int

	// Initializing a slice with values
    numbers = []int{1, 2, 3, 4, 5}

	// Append elements to a Slice
    numbers = append(numbers, 6, 7, 8)

	// Accessing elements 
    fmt.Println("fitst element", numbers[0])
    fmt.Println("last element", numbers[len(numbers)-1])

	// Subdividing a slice
    fmt.Println("The first three elements:", numbers[:3])

	// Iterating ober a slice
    for i, zahl := range numbers {
        fmt.Println("index:", i, "value:", zahl)
    }

	// Length and capacity
    fmt.Println("length of sliceu:", len(numbers))
    fmt.Println("capacity of slices", cap(numbers))
}
