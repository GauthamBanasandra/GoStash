package main

import "fmt"

const (
	PI = 22 / 7.0
	Language = "Go"

	/* iota represents successively incremented numbers. It's actually the index of the constant.
	Hence A=2, B=3, C=4
	*/
	A = iota
	B = iota
	C = iota
)

const (
	/* Omitting the value causes the previous value to be assigned. Since iota will be assigned,
	it's incremented for each successive declaration.
	*/
	D = iota
	E
	F
)

func main() {
	fmt.Println(PI)
	fmt.Println(Language)
	fmt.Println(A, B, C)
	fmt.Println(D, E, F)
}
