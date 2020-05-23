package numeral

import "strings"

func ConvertToRoman(x int) string {
	var result strings.Builder

	for x > 0 {
		switch {
		case x > 4:
			x -= 5
			result.WriteString("V")
		case x > 3:
			x -= 4
			result.WriteString("IV")
		default:
			x -= 1
			result.WriteString("I")
		}

	}

	return result.String()
}
