package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}
type Salutations []Salutation

// A method that will populate the channel.
func (salutations Salutations) SalutationsChannel(c chan Salutation) {
	for _, s := range salutations {
		// Put data into the channel.
		c <- s
	}
	// Don't forget to close the channel after all the data has been sent. Otherwise, it'll cause a deadlock.
	close(c)
}

func main() {
	slice := Salutations{
		{"Bob", "Hello"},
		{"Amy", "Hi"},
		{"Count Belmonte", "Greetings"},
		{"Mary", "Pleased to meet you"},
		{"Van Helsing", "Good morning"},
	}
	// Make a channel of type Salutation.
	c := make(chan Salutation)
	go slice.SalutationsChannel(c)

	// Read from the channel using 'range'.
	for s := range c {
		fmt.Println(s.greeting + " " + s.name)
	}
}
