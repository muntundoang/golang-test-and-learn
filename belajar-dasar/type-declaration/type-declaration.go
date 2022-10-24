package main

import "fmt"

func main() {
	type NoKTP string
	type status bool

	var noKTPOrang NoKTP = "13212418237549992312"
	var paymentStatus status = true

	fmt.Println(noKTPOrang)
	fmt.Println(paymentStatus)
}