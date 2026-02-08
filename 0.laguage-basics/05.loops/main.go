package main

import "fmt"

func main() {
	// for {} // creates infinite loop

	// While Loop
	i := 1
	for i <= 5 {
		fmt.Println(i, " ")
		i += 1 // Same as: i = i + 1
	}

	// Classic For Loop
	iterator := 0
	for i := 100; i > 0; i-- {
		if i%2 == 0 {
			continue
		} else if i <= 90 {
			break
		}
		iterator++ // Same as: iterator = iterator + 1
		fmt.Println(iterator, " - Odd number: ", i)
	}

	// Nested loop using range
	for i := range 10 {
		for range i + 1 {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
