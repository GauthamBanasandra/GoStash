package main
import "fmt"

func main() {
	// Go automatically infers the type, so it's not mandatory to specify it.
	var message string="hello go"
	var a, b, c int

	// Short way
	x, y, z:=1, 2.5, 6

	fmt.Println(message, a, b, c)
	fmt.Println(x, y, z)
}
