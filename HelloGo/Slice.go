package main

import (
	"fmt"
	"strconv"
)

type Printer func(string) ()
type Salutation struct {
	name     string
	greeting string
}

func GreetSlice(salutation []Salutation, print Printer) {
	/* Using slice.
	'i' will be the index and 's' will be the actual value.
	*/
	for i, s := range salutation {// Use strconv to convert int to string.
		message := s.greeting + " " + s.name + " you are " + strconv.Itoa(i)
		print(message)
	}
}
func main() {
	// Declaring a slice.
	slice := []Salutation{
		{"Count Gabriel Belmonte", "Salute"},
		{"Joe", "Hi"},
		{"Mary", "What's up?"},
	}

	GreetSlice(slice, func(s string) {
		fmt.Println(s)
	})
}
