package countingconverter

import (
	"fmt"

	"github.com/hsindorf/iroiro/stringutils"
)

// ConvertToLargestUnit converts a float to a string by Japanese counting system
// by the largest unit possible while still being greater than 1
func ConvertToLargestUnit(input float64) string {
	if chou := input / 1000000000000; chou >= 1 {
		return fmt.Sprintf("%v兆", stringutils.Commafy(chou))
	}

	if oku := input / 100000000; oku >= 1 {
		return fmt.Sprintf("%v億", stringutils.Commafy(oku))
	}

	if man := input / 10000; man >= 1 {
		return fmt.Sprintf("%v万", stringutils.Commafy(man))
	}

	return fmt.Sprintf("%v", stringutils.Commafy(input))
}
