package main

import (
	"fmt"
	"introduction-tests/addresses"
)

func main() {
	addressType := addresses.AddressType("Rua dos bobos")
	fmt.Println(addressType)

	invalidAddressType := addresses.AddressType("Sauna dos bobos")
	fmt.Println(invalidAddressType)
}
