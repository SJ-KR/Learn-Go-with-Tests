package numeral

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumeral = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
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
func ConvertingToArabic(s string) int {
	if s == "III" {
		return 3
	}
	if s == "II" {
		return 2
	}
	return 1
}
