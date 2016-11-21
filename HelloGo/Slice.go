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
	for i, s := range salutation {
		// Use strconv to convert int to string.
		message := s.greeting + " " + s.name + " you are " + strconv.Itoa(i)
		print(message)
	}
}
func main() {
	// Declaring a slice - short version.
	slice := []Salutation{
		{"Count Gabriel Belmonte", "Salute"},
		{"Joe", "Hi"},
		{"Mary", "What's up?"},
	}

	GreetSlice(slice, func(s string) {
		fmt.Println(s)
	})

	// Declaring a slice - elaborate version.
	var s []int        // Format : var <slice_name> []<type>
	// Do not specify the size. It will become an array otherwise.
	// Need to call make(slice_type, length, capacity). If capacity is omitted, then capacity = length.
	s = make([]int, 3)        // Here, capacity=length=3.
	s[0] = 1
	s[1] = 2
	s[2] = 3

	fmt.Println(s)

	// Slicing slices - just like python.
	fmt.Println(s[:2])

	// Appending to a slice - append(prev_slice, new_element). Returns a new slice.
	s = append(s, 5)
	fmt.Println(s)

	// Deleting an element from a slice.
	// For example, deleting the 1st element i.e 2 in the slice 's'.
	s = append(s[:1], s[2:]...)	// ... here enumerates the elements in the slice.
	fmt.Println(s)
}
