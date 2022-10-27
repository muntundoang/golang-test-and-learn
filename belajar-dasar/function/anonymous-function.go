package main

import "fmt"

type BlackList func(name string) bool

func registerUser(name string, blackList BlackList)  {
	if blackList(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main()  {
	blackList := func(name string) bool {
		return name == "kasar"
	}
	registerUser("Rio", blackList)
	registerUser("kasar", blackList)

	registerUser("masih kasar", func (name string) bool {
		return name == "masih kasar"
	})
}