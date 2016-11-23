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

	/*// Read from the channel using 'range'.
	for s := range c {
		fmt.Println(s.greeting + " " + s.name)
	}*/

	c2 := make(chan Salutation)
	go slice.SalutationsChannel(c2)

	// An elegant way to loop over both the channels is to use a 'select' block.
	for {
		/*
		It is just like switch - case.
		It selects whichever case(or rather channel) is ready to read/write.
		*/
		select {
		// Variable 'ok' will be false when the channel is closed. Hence, we can return.
		case s, ok := <-c:
			if ok {
				fmt.Println(s.greeting + " " + s.name + "\t channel 1")
			} else {
				return
			}
		case s, ok := <-c2:
			if ok {
				fmt.Println(s.greeting + " " + s.name + "\t channel 2")
			} else {
				return
			}
		default:
			fmt.Println("Waiting for data from channel.")
		}
	}
}
