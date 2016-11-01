package main

import "fmt"

func main() {
	message := "hello go"
	// Pointer declaration (*string is redundant).
	var greeting *string = &message

	fmt.Print(message, "\t", greeting)
}
