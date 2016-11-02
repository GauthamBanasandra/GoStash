package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func main() {
	s := Salutation{"Count Gabriel Belmonte", "Salutation"}
	Greet(s)
}

func Greet(salutation Salutation) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting)
	fmt.Println(message)
	fmt.Println(alternate)

	// Go throws an error if any variable is unused. So, use an underscore to ignore that variable
	_, alt1 := CreateMessage(salutation.name, salutation.greeting)
	fmt.Println(alt1)

	message, alternate = CreateMessage2(salutation.name, salutation.greeting)
	fmt.Println(message, alternate)

	// Variadic function.
	message, alternate = CreateMessage3(salutation.name, salutation.greeting, "yo")
	fmt.Println("Variadic function\t" + message)
	fmt.Println("Variadic function\t" + alternate)
}

// Function returning multiple values.
func CreateMessage(name, greeting string) (string, string) {
	return greeting + "\t" + name, "Hey!" + name
}

// Naming the return values.
func CreateMessage2(name, greeting string) (message string, alternate string) {
	message = greeting + "\t" + name
	alternate = "Hey!\t" + name

	// TODO : Figure out why this works.
	return
}

// Variadic function.
func CreateMessage3(name string, greeting ...string) (message string, alternate string) {
	// Getting the length of the greeting.
	fmt.Println(len(greeting))

	message = greeting[1] + "\t" + name
	alternate = "Hey!\t" + name
	return
}