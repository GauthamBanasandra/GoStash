package main

import (
	"os"
	"fmt"
)

func main() {
	var s string
	for _, arg := range os.Args {
		s += arg + " "
	}
	fmt.Println(s)
}
