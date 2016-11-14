package main

import "fmt"

type Printer func(string) ()
type Salutation struct {
	name     string
	greeting string
}

func greet3(salutation Salutation, do Printer) {
	do(salutation.name)
	do(salutation.greeting)
}

func CreatePrintFunction(custom string) (Printer) {
	return func(s string) {
		fmt.Println(s + custom)
	}
}
func main() {
	s := Salutation{"Ezio Auditore Da Firenze", "Salute"}
	greet3(s, CreatePrintFunction("!!!"))
}
