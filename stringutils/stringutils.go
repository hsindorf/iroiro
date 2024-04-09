package stringutils

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func TruncateToTwo(input float64) string {
	if input == math.Trunc(input) {
		return fmt.Sprintf("%.0f", input)
	}

	if input*10 == math.Trunc(input*10) {
		return fmt.Sprintf("%.1f", input)
	}

	return fmt.Sprintf("%.2f", input)
}

// Commafy takes a float, truncates to two and adds commas to whole number
func Commafy(num float64) string {
	wholeNumber := math.Floor(num)
	wholeNumberInt := int(wholeNumber)
	decimal := num - wholeNumber
	decimalNoZero := strings.TrimLeft(fmt.Sprintf("%v", TruncateToTwo(decimal)), "0")

	p := message.NewPrinter(language.English)
	numWithCommas := p.Sprintf("%v%v", wholeNumberInt, decimalNoZero)

	return numWithCommas
}
