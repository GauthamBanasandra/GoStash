package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func main() {
	var s=Salutation{"Count Gabriel Belmonte", "hello"}
	fmt.Print(s.name)
	fmt.Print(s.greeting)
	var t=Salutation{greeting:"Orion", name:"Greetings from Orion"}
	fmt.Println(t)
	var y=Salutation{}
	y.name="Robin hood"
	//y.greeting="hola"
	fmt.Println(y)
}
