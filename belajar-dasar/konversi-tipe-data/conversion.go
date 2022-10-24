package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 129
		nilai64 int64 = int64(nilai32) // <-- konversi integer32 menjadi integer64
		nilai8 int8 = int8(nilai32) // <-- tidak akan bisa dikonversi karena limit int8 hanya 127 sehingga akan menghasilkan -127 (max: 128) dikarenakan akan kembali ke nilai minimal int8 yaitu -127
	)

	fmt.Println("nilai32 =", nilai32)
	fmt.Println("nilai64 =", nilai64)
	fmt.Println("nilai8 =", nilai8)

	var (
		name = "Rio"
		charR = name[0] // byte a.k.a uint8
		rString = string(charR) // konversi byte ke string sehingga hasilnya ada R
	)

	fmt.Println("rString ->", rString)
}