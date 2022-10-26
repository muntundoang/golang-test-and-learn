package main
import "fmt"

func main() {

	var (
		ujian int8 = 80
		absensi int8 = 75
		lulusUjian = ujian >= 80
		lulusAbsensi = absensi >= 80
		lulus = lulusAbsensi && lulusUjian
	)

	fmt.Println("lulusUjian", lulusUjian)
	fmt.Println("lulusAbsensi" ,lulusAbsensi)
	fmt.Println("lulus" ,lulus)

}