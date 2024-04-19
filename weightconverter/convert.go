package weightconverter

import (
	"fmt"

	"github.com/hsindorf/iroiro/countingconverter"
	"github.com/hsindorf/iroiro/stringutils"
)

var kgToLbRate = 2.205

func ConvertWeight(unit string, amount float64, useJPUnits bool) string {
	switch unit {
	case "kg":
		return kgToLbs(amount, useJPUnits)
	case "lbs":
		return lbsToKg(amount, useJPUnits)
	}
	return ""
}

func kgToLbs(kg float64, useJPUnits bool) string {
	lbs := kg * kgToLbRate

	if useJPUnits {
		return fmt.Sprintf("%vlbs", countingconverter.ConvertToLargestUnit(lbs))
	}

	return fmt.Sprintf("%vlbs", stringutils.Commafy(lbs))
}

func lbsToKg(lbs float64, useJPUnits bool) string {
	kg := lbs / kgToLbRate

	if useJPUnits {
		return fmt.Sprintf("%vkg", countingconverter.ConvertToLargestUnit(kg))
	}

	return fmt.Sprintf("%vkg", stringutils.Commafy(kg))
}
