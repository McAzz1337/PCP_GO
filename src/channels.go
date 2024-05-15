package main

import "fmt"

func channels_xapmle_1() {
	// Kanal erstellen
	ch := make(chan string)

	// Goroutine starten, die eine Nachricht sendet
	go func() {
		ch <- "Hello from the other side"
	}()

	// Haupt-Goroutine wartet auf die Nachricht
	msg := <-ch
	fmt.Println(msg) // Ausgabe: Hello from the other side
}
