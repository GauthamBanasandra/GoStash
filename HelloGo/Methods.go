package main

import (
	"strconv"
	"fmt"
)

type Salutation struct {
	name     string
	greeting string
}
// First declare a named type.
type Salutations []Salutation
type Printer func(string) ()

func GetAddress(name string) (string) {
	greetMap := map[string]string{
		"Bob":"Mr. ",
		"Count Gabriel Belmonte":"Sir ",
		"Mary":"Lady ",
	}

	if value, exists := greetMap[name]; exists {
		return value
	}

	return "Dude "
}

/*
Declare a method by -
func (type on which it should run) method_name (params)
*/
func (salutations Salutations) GreetMethod(do Printer) {
	for i, salutation := range salutations {
		do(salutation.greeting + ", " + GetAddress(salutation.name) + " " + salutation.name + " you are " + strconv.Itoa(i + 1))
	}
}

func main() {
	// Declare a variable of the type just defined.
	slice := Salutations{
		{"Count Gabriel Belmonte", "Greetings"},
		{"Joe", "Hi"},
		{"Mary", "What's up?"},
	}

	// Call the method on the variable
	slice.GreetMethod(func(s string) {
		fmt.Println(s + "!!")
	})
}
