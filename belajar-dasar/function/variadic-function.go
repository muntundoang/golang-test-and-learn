package main

import "fmt"

func sumAll(tipeOp string, numbers ...int) (tipe string,res int) {
	res = 0
	tipe = tipeOp
	for _, number := range numbers {
		res += number
	}

	return
}

func main() {
	tipe, res := sumAll("sum", 125, 5498, 8466, 543218, 78, 8)
	fmt.Println( tipe, "=", res)

	slice := []int{10,20,30,40,50}
	tipe, res = sumAll("sum", slice...)
	fmt.Println( tipe, "=", res)
}
