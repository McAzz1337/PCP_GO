package main

import "fmt"

// Funktion zur Berechnung der Fibonacci-Zahl
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var input int

	// Eingabeaufforderung
	fmt.Print("Geben Sie eine Zahl ein: ")
	// Eingabe lesen
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Ungültige Eingabe:", err)
		return
	}

	// Fibonacci-Zahl berechnen und ausgeben
	result := fibonacci(input)
	fmt.Printf("Die Fibonacci-Zahl für %d ist %d\n", input, result)
}
