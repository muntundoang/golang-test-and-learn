package main

import "fmt"

func main() {
	var name string = "maman"

	switch name {
	case "dudung":
		fmt.Println("Temennya maman")

	case "maman":
		fmt.Println("Temennya dudung")

	default:
		fmt.Println("bukan anak komplek")
	}

	switch length := len(name); length >= 5 {
	case true:
		fmt.Println(length)
		fmt.Println("namanya sama atau lebih dari 5 huruf")

	case false:
		fmt.Println(length)
		fmt.Println("namanya kurang dari 5 huruf")
	}

	length := len(name)

	switch { // switch bisa tanpa kondisi | kondisi bisa ditaro di case

	case length > 10:
		fmt.Println(length)
		fmt.Println("namanya sama atau lebih dari 10 huruf")

	case length >= 5:
		fmt.Println(length)
		fmt.Println("namanya sama atau lebih dari 10 huruf")

	default:
		fmt.Println("bukan anak komplek")

	}
}
