// unit tests
package addresses

import "testing"

type addressMock struct {
	address             string
	addressTypeExpected string
}

func TestAddressType(t *testing.T) {
	addressesCases := []addressMock{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Invalid address type!"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Invalid address type!"},
	}

	for _, addressCase := range addressesCases {
		if addressType := AddressType(addressCase.address); addressType != addressCase.addressTypeExpected {
			t.Errorf("O tipo recebido '%s' é diferente do tipo esperado '%s'", addressType, addressCase.addressTypeExpected)
		}
	}
}
