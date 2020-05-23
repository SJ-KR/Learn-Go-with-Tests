package numeral

import "strings"

func ConvertToRoman(x int) string {
	var result strings.Builder
	for i := 0; i < x; i++ {
		result.WriteString("I")
	}

	return result.String()
}
