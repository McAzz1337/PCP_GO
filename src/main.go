package main

import "fmt"

func main() {

	fmt.Println("Cons")
	cons()
	fmt.Println()

	fmt.Println("Eval expr")
	eval_expression()
	fmt.Println()

	fmt.Println("Defer")
	defer_xmpl_1()
	defer_xmpl_2()
	fmt.Println()

	fmt.Println("Typing")
	xample_typing_1()
	fmt.Println()

	fmt.Println("Maps")
	maps_xapmle_1()
	fmt.Println()

	fmt.Println("Slices")
	slices_xample_1()
	fmt.Println()

	fmt.Println("Goroutine")
	goroutine_xapmle_1()
	fmt.Println()

	fmt.Println("Channels")
	channels_xapmle_1()
	fmt.Println()

	fmt.Println("Fib")
	fib()
	fmt.Println()

	fmt.Println("Parallel Streams solved with Goroutines")
	parallel_Goroutines()
	fmt.Println()

	var input int
	// Eingabeaufforderung
	fmt.Print("Bereit für Select? (y)")
	// Eingabe lesen
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("Select")
	select_xapmle_1()
	fmt.Println()
}
