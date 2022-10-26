package main

import "fmt"

func main() {
	var (
		a = 12
		b = 25
		c = 2
		d = 14
		e = 10
	)

	 a += 10 // = a + 10
	 b -= 5 // = b - 5
	 c *= 3 // = c * 3
	 d /= 14 // = d / 14
	 e %= 3 // = e % 3

	fmt.Println("augmented +", a)
	fmt.Println("augmented -", b)
	fmt.Println("augmented *", c)
	fmt.Println("augmented /", d)
	fmt.Println("augmented %", e)
}