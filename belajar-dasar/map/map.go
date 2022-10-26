package main

import "fmt"

func main()  {
	person := map[string]string{
		"name" : "Rio",
		"address": "Bandung",
	}

	person["job"] = "nganggur"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["job"])
	fmt.Println("-----------------------------")

	delete(person, "job")
	fmt.Println("deleted job", person)
}