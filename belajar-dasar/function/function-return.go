package main

import "fmt"

func getHello(name string) string { // <-- akan mengembalikan tipe data string
	return "Hello " + name
}

func main()  {
	result := getHello("Rio")
	fmt.Println(result)
}