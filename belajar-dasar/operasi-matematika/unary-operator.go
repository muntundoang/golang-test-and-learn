package main

import "fmt"

func main() {
	var (
		a = 12
		b = 25
	)

	 a++ // = a + 1
	 b-- // = b - 1

	fmt.Println("unary +", a)
	fmt.Println("unary -", b)
}