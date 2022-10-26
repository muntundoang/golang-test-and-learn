package main

import "fmt"

func main() {
	var (
		name1 string = "dudung"
		name2 string = "maman"
		val1 = 100
		val2 = 105
	)

	var compare1 bool = name1 > name2 // hasilnya false karena huruf pertama name1 "d" lebih kecil daripada huruf pertama name2 "m" (berdasarkan abjad a b c ...) 
	var compare2 bool = val1 < val2
	fmt.Println("name1 > name2", compare1)
	fmt.Println("val1 < val2", compare2)
}