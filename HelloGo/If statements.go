package main

import "fmt"

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
	do(salutation.greeting)
}

func CreatePrintFunction1(custom string) (Printer) {
	return func(s string) {
		fmt.Println(s + custom)
	}
}
func main() {
	s := Salutation{"Ezio Auditore Da Firenze", "Salute"}
	greet4(s, CreatePrintFunction1("!!!"), false)
}
