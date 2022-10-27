package main

import "fmt"

type Filter func(string) string

func sayGoodNightFiltered(word string, filter Filter)  {
	wordFiltered := filter(word)
	fmt.Println("Good Night", wordFiltered)
}

func wordFiltering(word string) string {
	if word == "Pak" {
		return "******"
	} else {
		return word
	}
}

func main()  {
	sayGoodNightFiltered("Rio", wordFiltering)
	sayGoodNightFiltered("Pak", wordFiltering)
}