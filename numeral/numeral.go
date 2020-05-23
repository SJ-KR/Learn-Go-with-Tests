package numeral

import (
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}
type RomanNumerals []RomanNumeral

var allRomanNumerals = RomanNumerals{
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

	for _, numeral := range allRomanNumerals {
		for x >= numeral.Value {
			result.WriteString(numeral.Symbol)
			x -= numeral.Value
		}
	}

	return result.String()
}
func (r RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)

	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}
func CouldBeSubtractive(i int, symbol uint8, roman string) bool {
	subtractiveSymbol := symbol == 'I' || symbol == 'X' || symbol == 'C'
	return i+1 < len(roman) && subtractiveSymbol
}
func ConvertToArabic(roman string) int {
	total := 0
	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if CouldBeSubtractive(i, symbol, roman) {
			if value := allRomanNumerals.ValueOf(symbol, roman[i+1]); value != 0 {
				total += value
				i++
			} else {
				total += allRomanNumerals.ValueOf(symbol)
			}
		} else {
			total += allRomanNumerals.ValueOf(symbol)
		}
	}
	return total
}
