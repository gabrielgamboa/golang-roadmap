package addresses

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func AddressType(address string) string {
	validTypes := []string{"rua", "avenida", "estrada", "rodovia"}

	addressInLowerCase := strings.ToLower(address)
	addressType := strings.Split(addressInLowerCase, " ")[0]

	hasValidType := false

	for _, validType := range validTypes {
		if addressType == validType {
			hasValidType = true
		}
	}

	if hasValidType {
		caser := cases.Title(language.BrazilianPortuguese)
		return caser.String(addressType)
	}

	return "Invalid address type!"
}
