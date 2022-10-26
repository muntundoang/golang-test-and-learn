package main
import "fmt"

func main()  {
	var myName [3]string // deklarasi variable array myName dengan kapasitas / length 3 element
	myName[0] = "Rio"
	myName[1] = "Lukman"
	myName[2] = "Tawekal"

	fmt.Println("array myName ->", myName)
	fmt.Println(myName[0])
	fmt.Println(myName[1])
	fmt.Println(myName[2])
	fmt.Println("")
	// fmt.Println(myName[3]) <-- akan error karena kapasitas array hanya ada 3

	var values = [3]int{
		1,2,3,
	}

	var zeroElement = [3]int{
		
	}

	fmt.Println(values)
	fmt.Println("length values ->", len(values))
	fmt.Println("length zeroElement ->", len(zeroElement))
}