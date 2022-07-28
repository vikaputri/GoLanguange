package main

import "fmt"

func main() {
	var value1 = true
	var value2 = false
	fmt.Println(value1 && value2) //and
	fmt.Println(value1 || value2) //or
	fmt.Println(!value1)          //negasi

	var nilaiAkhir = 98
	var absensi = 88
	var lulusNilaiAkhir bool = nilaiAkhir > 80
	fmt.Println(lulusNilaiAkhir)

	var lulusAbsensi = absensi > 80
	fmt.Println(lulusAbsensi)

	var lulus = lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)

	fmt.Println(nilaiAkhir > 80 && absensi > 80)

}
