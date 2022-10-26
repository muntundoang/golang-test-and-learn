package main

import "fmt"

func main()  {
	//Break
	for counterFor := 1; counterFor <= 10; counterFor++ {
		if counterFor == 5 {
			fmt.Println("looping ke", counterFor)
			fmt.Println("end of break loop")
			break
		}
		fmt.Println("looping ke", counterFor)
	}

	//Continue
	for counterFor := 1; counterFor <= 10; counterFor++ {
		if counterFor % 2 == 1 {
			continue
		}
		fmt.Println("looping ke", counterFor)
	}
}