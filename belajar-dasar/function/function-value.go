package main

import "fmt"

func getGoodNight(name string) string {
	return "Good Night " + name
}

func main() {
	sayGoodNight := getGoodNight
	fmt.Println(sayGoodNight("Rio"))
}
