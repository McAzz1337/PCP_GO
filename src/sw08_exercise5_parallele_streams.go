package main

import (
	"fmt"
	"sync"
)

func parallel_Goroutines() {
	// Creating a channel to receive the integers
	intChan := make(chan int)

	// WaitGroup to ensure all goroutines finish before exiting main
	var wg sync.WaitGroup
	wg.Add(1)

	// Goroutine to generate integers from 0 to 99
	go func() {
		defer close(intChan)
		for i := 0; i < 100; i++ {
			intChan <- i
		}
	}()

	// Goroutine to print the integers in parallel
	go func() {
		for num := range intChan {
			fmt.Printf("%d, ", num)
		}
		wg.Done()
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
