package main

import "fmt"

func sayHello()  {
	fmt.Println("Halo Dunia")
}

func main()  {
	
	for i:=0; i < 5; i++{
		sayHello()
	}
	
}