package main

import "fmt"

func main() {
	// Normal for loop.
	for i := 0; i < 5; i++ {
		fmt.Println("hello")
	}

	// Using for as a while loop.
	i := 0
	for i < 10 {
		fmt.Println("loop number " + string(i))
		i += 1
	}

	// Usage of break and continue is same as C.
}
