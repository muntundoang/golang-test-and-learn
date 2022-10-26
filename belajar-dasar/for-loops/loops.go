package main

import "fmt"

func main()  {
	counter := 1

	for counter <= 10 {
		fmt.Println("looping ke", counter)
		counter++
	}
	fmt.Println("looping 1 End")
	fmt.Println("")

	for counterFor := 1; counterFor <= 10; counterFor++ {
		fmt.Println("looping ke", counterFor)
	} 

	fmt.Println("looping 2 End")
	fmt.Println("")

	slice := []string{"Rio", "Lukman", "Tawekal"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Println("looping 3 End")
	fmt.Println("")

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	fmt.Println("looping 4 End")
	fmt.Println("")

	for _, value := range slice { // i bisa diganti jadi "_" kalau hanya ingin mengambil valuenya saja
		fmt.Println(value)
	}

	fmt.Println("looping 5 End")
	fmt.Println("")

	data := make(map[string]string)
	data["nama"] = "Rio Lukman Tawekal"
	data["alamat"] = "Bandung"

	for key, val := range data {
		fmt.Println("key =", key, "|", "Value", "=", val)
	}

	fmt.Println("looping 6 End")
	fmt.Println("")
}