package numeral

import (
	"fmt"
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
func (r RomanNumerals) ValueOf(symbol string) int {
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}
func CouldBeSubtractive(i int, symbol uint8, roman string) bool {
	return i+1 < len(roman) && symbol == 'I'
}
func ConvertToArabic(roman string) int {
	total := 0
	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if CouldBeSubtractive(i, symbol, roman) {
			nextSymbol := roman[i+1]

			potentialNumber := string([]byte{symbol, nextSymbol})
			fmt.Println(potentialNumber)
			if value := allRomanNumerals.ValueOf(potentialNumber); value != 0 {
				total += value
				i++
			} else {
				total++
			}
		} else {
			total += allRomanNumerals.ValueOf(string([]byte{symbol}))
		}
	}
	return total
}
