package main

import "fmt"

func getFullName() (firstName , middleName, lastName string) { // <-- akan mengembalikan tipe data string
	firstName = "Rio"
	middleName = "Lukman"
	lastName = "Tawekal"

	return
}

func main()  {
	a, b, c := getFullName()
	fmt.Println(getFullName())
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
}