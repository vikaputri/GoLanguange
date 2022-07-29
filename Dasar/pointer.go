package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	//address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	//address2 := address1
	//address2.City = "Bandung"

	//fmt.Println(address1) //{Subang Jawa Barat Indonesia}
	//fmt.Println(address2) //{Bandung Jawa Barat Indonesia}

	//var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	//var address2 *Address = &address1
	//address2.City = "Bandung"
	//fmt.Println(address1) //{Bandung Jawa Barat Indonesia}
	//fmt.Println(address2)  //&{Bandung Jawa Barat Indonesia}

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City = "Bandung"

	//address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	//fmt.Println(address1) //{Bandung Jawa Barat Indonesia}
	//fmt.Println(address2) //&{Jakarta DKI Jakarta Indonesia}

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1) //{Jakarta DKI Jakarta Indonesia}
	fmt.Println(address2) //&{Jakarta DKI Jakarta Indonesia}
	fmt.Println(address3) //&{Jakarta DKI Jakarta Indonesia}

	var address4 *Address = new(Address)
	address4.City = "Malang"
	fmt.Println(address4)

}
