package main

import "fmt"

func main()  {
	//Continue
	for counterFor := 1; counterFor <= 10; counterFor++ {
		if counterFor % 2 == 1 {
			continue
		}
		fmt.Println("looping ke", counterFor)
	}
}