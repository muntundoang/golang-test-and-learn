package main

import "fmt"

func main() {
	var name string = "maman"

	if name == "dudung" {
		fmt.Println("Temennya maman")
	} else if name == "maman" {
		fmt.Println("Temennya dudung")
	} else {
		fmt.Println("bukan temennya maman")
	}

	if length := len(name) ; length >= 5 { // <--- bisa deklarasi dalam if statement | variable hanya bisa digunakan dalam if statement
		fmt.Println(length)
		fmt.Println("namanya sama atau lebih dari 5 huruf")
	} else {
		fmt.Println(length)
		fmt.Println("namanya kurang dari 5 huruf")
	}

}
