package main

import (
	"fmt"
)

type Printer func(string) ()
type Salutation struct {
	name     string
	greeting string
}

func greet4(salutation Salutation, do Printer, isFormal bool) {
	// If statement.
	if isFormal {
		do(salutation.name)
	}

	/* Embedding statements. The "prefix" is executed only once and its scope is confined to that line.
	It's like the initialization done in a for loop.
	 */
	if prefix := GetPrefix(salutation.name); isFormal {
		do(prefix + salutation.name)
	} else {
		do(salutation.greeting)
	}
}

func GetPrefix(name string) (prefix string) {
	// Switch - case.
	switch name {
	case "Bob":
		prefix = "Mr. "
		fallthrough // Presence of a "fallthrough" causes the next statement to be chosen.
	case "Joe", "Amy": // Valid for "Joe" or "Amy" as input.
		prefix = "Dr. "
	case "Count":
		prefix = "Sir "
	default:
		prefix = "Dude "
	}

	// Turning switch into an if-else kind of branch.
	switch {
	case name == "Bob":
		prefix = "Mr. "
	case len(name) == 4:
		prefix = "Short name "
	}
	return
}

func TypeSwitchTest(x interface{}) {
	// Switching based on type of the variable.
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Salutation:
		fmt.Println("Salutation")
	default:
		fmt.Println("unknown")
	}
}
func CreatePrintFunction1(custom string) (Printer) {
	return func(s string) {
		fmt.Println(s + custom)
	}
}
func main() {
	s := Salutation{"Amy1", "Salute"}
	greet4(s, CreatePrintFunction1("!!!"), true)
	fmt.Println("\nFiguring out the type")
	TypeSwitchTest(s)
}
