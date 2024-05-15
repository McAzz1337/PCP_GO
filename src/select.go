package main

import (
	"fmt"
	"time"
)

func select_xapmle_1() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine, die alle 100ms eine Nachricht auf ch1 sendet
	go func() {
		for {
			ch1 <- "from ch1"
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Goroutine, die alle 200ms eine Nachricht auf ch2 sendet
	go func() {
		for {
			ch2 <- "from ch2"
			time.Sleep(200 * time.Millisecond)
		}
	}()

	// Haupt-Goroutine, die auf die Nachrichten von ch1 und ch2 wartet
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received", msg2)
		}
	}
}
