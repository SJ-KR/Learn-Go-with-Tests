package numeral

import "strings"

func ConvertToRoman(x int) string {
	var result strings.Builder

	for i := x; i > 0; i-- {
		if i == 4 {
			result.WriteString("IV")
			break
		}
		result.WriteString("I")
	}

	return result.String()
}
