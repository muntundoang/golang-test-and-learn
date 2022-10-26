package main

import "fmt"

func main() {
	const (
		a = 12
		b = 25
	)

	var (
		pertambahan = a + b
		pengurangan = a - b
		perkalian = a * b
		pembagian = b / a
	)

	fmt.Println("pertambahan", pertambahan)
	fmt.Println("pengurangan", pengurangan)
	fmt.Println("perkalian", perkalian)
	fmt.Println("pembagian", pembagian)
}