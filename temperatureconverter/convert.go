package temperatureconverter

import (
	"fmt"

	"github.com/hsindorf/iroiro/countingconverter"
	"github.com/hsindorf/iroiro/stringutils"
)

func ConvertTemperature(unit string, amount float64, useJPUnits bool) string {
	switch unit {
	case "c":
		return cToF(amount, useJPUnits)
	case "f":
		return fToC(amount, useJPUnits)
	}
	return ""
}

func cToF(temp float64, useJPUnits bool) string {
	var f float64 = (temp * 9.0 / 5.0) + 32.0

	if useJPUnits {
		return fmt.Sprintf("%v°F", countingconverter.ConvertToLargestUnit(f))
	}

	return fmt.Sprintf("%v°F", stringutils.Commafy(f))
}

func fToC(temp float64, useJPUnits bool) string {
	var c float64 = (temp - 32.0) * 5.0 / 9.0

	if useJPUnits {
		return fmt.Sprintf("%v°C", countingconverter.ConvertToLargestUnit(c))
	}

	return fmt.Sprintf("%v°C", stringutils.Commafy(c))
}
