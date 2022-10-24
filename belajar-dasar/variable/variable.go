package main

import "fmt"

func main() {
	var name string
	name = "Rio Lukman Tawekal"
	fmt.Println(name)

	var angka uint8
	angka = 1
	fmt.Println("Satu = ", angka)

	//deklarasi tanpa var
	noVar := "string"
	fmt.Println("noVar = ", noVar)

	//multiple declaration variable
	var (
		firstName string = "Rio Lukman"
		lastName = "Tawekal"
	)
	fmt.Println(firstName,lastName)
}