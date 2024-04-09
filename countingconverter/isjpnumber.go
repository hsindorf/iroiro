package countingconverter

import "strings"

func IsJapaneseNumber(input string) bool {
	return strings.ContainsAny(input, "兆億万千")
}
