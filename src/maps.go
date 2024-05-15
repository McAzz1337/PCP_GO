package main

import "fmt"

func maps_xapmle_1() {
	// Creation of an enpty Mao that contains keys as Strings and intefers as values
    scores := make(map[string]int)

    // Initialisieren von Werten in der Map
    scores["Alice"] = 90
    scores["Bob"] = 85
    scores["Charlie"] = 95

    // Accessing values of a map
    fmt.Println("Alice's Punktzahl:", scores["Alice"])
    fmt.Println("Bob's Punktzahl:", scores["Bob"])
    fmt.Println("Charlie's Punktzahl:", scores["Charlie"])

    // Iterate over the map
    for name, score := range scores {
        fmt.Println(name, "hat", score, "Punkte erzielt")
    }

	// Check if a key in the map exists
    if _, ok := scores["Alice"]; ok {
        fmt.Println("Alice's Punktzahl existiert")
    }
}
