package numeral

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumeral = []RomanNumeral{
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(x int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumeral {
		for x >= numeral.Value {
			result.WriteString(numeral.Symbol)
			x -= numeral.Value
		}
	}

	return result.String()
}
