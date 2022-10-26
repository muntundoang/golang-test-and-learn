package main

import "fmt"

func main() {
	var (
		bool1 bool = true
		bool2 bool = false
		bool3 bool = true
		bool4 bool = false
	)

	// Operasi &&
	var condition1 = bool1 && bool2 // true && false
	var condition2 = bool1 && bool3 // true && true
	var condition3 = bool4 && bool3 // false && true

	fmt.Println("Operasi &&")
	fmt.Println("condition1",condition1)
	fmt.Println("condition2",condition2)
	fmt.Println("condition3",condition3)
	fmt.Println("")

	var condition4 = bool1 || bool2 // true || false
	var condition5 = bool1 || bool3 // true || true
	var condition6 = bool4 || !bool3 // false || true
	var condition7 = bool2 || !bool4 // false || false

	fmt.Println("Operasi ||")
	fmt.Println("condition4",condition4)
	fmt.Println("condition5",condition5)
	fmt.Println("condition6",condition6)
	fmt.Println("condition7",condition7)
	fmt.Println("")
}
