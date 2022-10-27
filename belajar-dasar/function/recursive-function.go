package main

import "fmt"

func factorialLoop(val int) int {
	res := 1
	for i := val; i > 0; i-- {
		res *= i
	}

	return res
}

func factorialRecursive(val int) int {
	if val == 1 {
		return 1
	} else {
		return val * factorialRecursive(val-1)
	}
}

func main() {
	loop := factorialLoop(5)
	recursiveWay := factorialRecursive(5)
	fmt.Println(loop)
	fmt.Println(5 * 4 * 3 * 2 * 1)
	fmt.Println(recursiveWay)
}
