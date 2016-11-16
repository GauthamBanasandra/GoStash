package main

import (
	"strconv"
	"fmt"
)

type Salutation struct {
	name     string
	greeting string
}
type Printer func(string) ()

func GetPrefixMap(name string) (string) {
	// Short hand maps.
	prefixMap := map[string]string{
		"Bob":"Mr ",
		"Joe":"Dr ",
		"Amy":"Dr ",
		"Mary":"Lady ",
		"Count Belmonte":"Sir ",
	}

	// Deleting an element.
	//delete(prefixMap, "Mary")
	// Depricated way is - prefixMap["Mary"]="", false

	/* Checking for existence.
	'exists' is a boolean variable. So, if it's true, it means that the key exists.
	 */
	if value, exists := prefixMap[name]; exists {
		return value
	}
	/*// Declare a map in the following way where typeof key=string and typeof value=string.
	var prefixMap map[string]string
	// Need to call make, otherwise, can't use the map.
	prefixMap = make(map[string]string)

	// Use the map just like in Python.
	prefixMap["Bob"] = "Mr. "
	prefixMap["Joe"] = "Dr. "
	prefixMap["Amy"] = "Dr. "
	prefixMap["Count Belmonte"] = "Sir "
	prefixMap["Mary"] = "Lady "*/

	return "Dude "
}

func GreetMap(salutation []Salutation, print Printer) {
	for i, s := range salutation {
		print(s.greeting + " " + GetPrefixMap(s.name) + s.name + " you are " + strconv.Itoa(i))
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
	GreetMap(slice, func(s string) {
		fmt.Println(s)
	})
}
