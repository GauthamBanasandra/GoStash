package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	var s string
	for i, arg := range os.Args {
		s += strconv.Itoa(i) + " " + arg + "\n"
	}

	fmt.Println(s)
}
