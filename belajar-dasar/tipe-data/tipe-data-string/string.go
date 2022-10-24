package main

import "fmt"

func main() {
	fmt.Println("Rio")
	fmt.Println("Rio Lukman")
	fmt.Println("Rio Lukman Tawekal")
	fmt.Println("Panjang string nya Rio Lukman Tawekal ada", len("Rio Lukman Tawekal"))
	fmt.Println("Rio"[0], "<- Dalam bentuk byte")
	fmt.Println("Rio Lukman"[8], "<- Dalam bentuk byte")
	fmt.Println("Rio Lukman Tawekal"[len("Rio Lukman Tawekal") - 2], "<- Dalam bentuk byte")
}