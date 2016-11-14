package greeting

import "fmt"

type Salutation struct {
	Name     string	// Variable "Name" must have first letter in caps - needs to be exported.
	Greeting string	// Variable "Greeting" must have first letter in caps - needs to be exported.
}

func GreetingFunc(salutation Salutation) {
	fmt.Print("Greetings fromt the greeting package!!\t" + salutation.Name)
}