package main

import "fmt"

func getHello(name string, name2 string) (string, string, string) { // <-- akan mengembalikan tipe data string
	return "Hello", name, name2
}

func main()  {

	hello, firstName, lastName := getHello("Rio", "Lukman")
	fmt.Println(hello, firstName, lastName)

	hello1, firstName1, _ := getHello("Rio", "Lukman") // <--- bisa di ignore pakai under score ("_")
	fmt.Println(hello1, firstName1)

}