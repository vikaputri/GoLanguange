package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Subang", "Jawa Barat", ""}
	ChangeAddressToIndonesia(&address)
	fmt.Println(address) // tidak berubah tetap {Subang Jawa Barat } agar berubah * pada fungsi dan & pada penggunaan fungsi

}
