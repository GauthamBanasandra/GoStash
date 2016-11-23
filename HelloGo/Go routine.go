package main

import (
	"fmt"
	"strconv"
)

type Salutation struct {
	name     string
	greeting string
}
type Printer func(string) ()

func GetGreetPrefix(name string) (string) {
	prefixMap := map[string]string{
		"Bob":"Mr ",
		"Joe":"Dr ",
		"Amy":"Dr ",
		"Mary":"Lady ",
		"Count Belmonte":"Sir ",
	}

	if value, exists := prefixMap[name]; exists {
		return value
	}

	return "Dude "
}

func GreetSalutation(salutation []Salutation, print Printer) {
	for i, s := range salutation {
		print(s.greeting + " " + GetGreetPrefix(s.name) + s.name + " you are " + strconv.Itoa(i))
	}
}

func main() {
	slice := []Salutation{
		{"Bob", "Hello"},
		{"Amy", "Hi"},
		{"Count Belmonte", "Greetings"},
		{"Mary", "Pleased to meet you"},
		{"Van Helsing", "Good morning"},
	}

	/*
	Put keyword 'go' in front of a function call to make it a go routine.
	It's just that, the call is executed in a separate thread.
	*/
	go GreetSalutation(slice, func(s string) {
		fmt.Println(s + "!!")
	})
	fmt.Println()
	GreetSalutation(slice, func(s string) {
		fmt.Println(s + ":)")
	})
}
