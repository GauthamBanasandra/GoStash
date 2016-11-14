package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

/* Declaring function type. This is mainly done for re-use of the declaration.
Syntax - type <name> func(<param type>) (<return type>). The empty brackets must be specified for void.
*/
type Printer func(string) ()

// The code below is an example for not using the function type declaration.
/*func greet(salutation Salutation, do func(string)) {
	do(salutation.name)
	do(salutation.greeting)
}*/

// Just use the function type 'Printer' that is just declared.
func greet(salutation Salutation, do Printer) {
	do(salutation.name)
	do(salutation.greeting)
}
func Print(s string) {
	fmt.Print(s)
}
func PrintLn(s string) {
	fmt.Println(s)
}
func main() {
	s := Salutation{"Prince", "Hello"}
	greet(s, PrintLn)
}
