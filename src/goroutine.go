package main

import (
	"fmt"
	"time"
)

// Funktion, die in einer Goroutine ausgeführt wird
func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}

func goroutine_xapmle_1() {
	// Goroutine starten, um die Funktion sayHello auszuführen
	go sayHello()

	// Hauptfunktion gibt während die Goroutine ausgeführt wird "Hello" aus
	for i := 0; i < 5; i++ {
		fmt.Println("World")
		time.Sleep(100 * time.Millisecond)
	}

	// Warten, damit die Goroutine beendet werden kann, bevor das Programm endet
	time.Sleep(1 * time.Second)
}
