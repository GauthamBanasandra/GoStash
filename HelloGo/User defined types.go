package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func main() {
	// Corresponding initialization.
	var s=Salutation{"Count Gabriel Belmonte", "hello"}
	fmt.Print(s.name)
	fmt.Print(s.greeting)

	// Explicit initialization.
	var t=Salutation{greeting:"Orion", name:"Greetings from Orion"}
	fmt.Println(t)

	// Auto initialization.
	var y=Salutation{}
	y.name="Robin hood"

	//y.greeting="hola"
	fmt.Println(y)
}
