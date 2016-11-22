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

// Declare an interface in the following manner.
type Renamable interface {
	Rename(newName string)        /*
				Salutation type has a method "Rename". Thus, Salutation type implements the
				"Renamable" interface.
				*/
}

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
Declare a method by the following format -
func (type on which it should run) method_name (params)
*/
func (salutations Salutations) GreetMethod(do Printer) {
	for i, salutation := range salutations {
		do(salutation.greeting + ", " + GetAddress(salutation.name) + " " + salutation.name + " you are " + strconv.Itoa(i + 1))
	}
}

/* Pointer to a method.
If the asterisk is not present, then 'salutation" is just a local copy.
Thus, without the asterisk, changes will not be reflected in the calling parameter.
*/
func (salutation *Salutation) Rename(newName string) {
	salutation.name = newName
}

/*
Implementing Write interface.
*/
func (salutation *Salutation) Write(p []byte) (n int, err error) {
	s := string(p)
	salutation.Rename(s)
	n = len(s)
	err = nil
	return
}

func RenameToFrog(r Renamable) {
	r.Rename("Frog")
}

func main() {
	// Declare a variable of the type just defined.
	slice := Salutations{
		{"Count Gabriel Belmonte", "Greetings"},
		{"Joe", "Hi"},
		{"Mary", "What's up?"},
	}

	// Renaming Joe to Bob.
	slice[1].Rename("Bob")

	/*
	Renaming using interface.
	Passing an address here because, RenameToFrog calls Rename, which accepts a pointer.
	*/
	RenameToFrog(&slice[2])

	/*
	Using the Writer interface.
	Fprintf will handle your input if you implement the "Write" interface.
	*/
	fmt.Fprintf(&slice[1], "The count is %d", 10)

	// Call the method on the variable
	slice.GreetMethod(func(s string) {
		fmt.Println(s + "!!")
	})
}
