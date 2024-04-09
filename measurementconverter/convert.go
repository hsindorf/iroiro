package measurementconverter

import (
	"fmt"
	"math"

	"github.com/hsindorf/iroiro/countingconverter"
	"github.com/hsindorf/iroiro/stringutils"
)

const (
	cmToInRate = 0.393701
	mToFtRate  = 3.28084
	kmToMiRate = 0.621371
)

func ConvertDistance(unit string, amount float64, useJPUnits bool) string {
	switch unit {
	case "cm":
		return CMToIn(amount, useJPUnits)
	case "in":
		return InToCM(amount, useJPUnits)
	case "m":
		return MtoFt(amount, useJPUnits)
	case "ft":
		return FtToM(amount, useJPUnits)
	case "km":
		return KMtoMi(amount, useJPUnits)
	case "mi":
		return MiToKM(amount, useJPUnits)
	}
	return ""
}

// CMToIn converts cm to inches. If > 1 ft, also provides ft'in" amount
func CMToIn(cm float64, useJPUnits bool) string {
	in := cm * cmToInRate
	if in > 12 {
		ft := in / 12
		ftFloor := math.Floor(ft)
		inDecimal := 12 * (ft - ftFloor)

		if useJPUnits {
			inString := fmt.Sprintf("%v", countingconverter.ConvertToLargestUnit(in))
			ftString := fmt.Sprintf("%v", countingconverter.ConvertToLargestUnit(ftFloor))
			return fmt.Sprintf("%vin (%vft %vin)", inString, ftString, stringutils.Commafy(inDecimal))
		}

		return fmt.Sprintf("%vin (%vft %vin)", stringutils.Commafy(in), stringutils.Commafy(ftFloor), stringutils.TruncateToTwo(inDecimal))
	}
	return fmt.Sprintf("%vin", stringutils.Commafy(in))
}

// InToCM converts inches to cm. If > 1m, only provides m amount
func InToCM(in float64, useJPUnits bool) string {
	cm := in / cmToInRate
	if cm > 100 {
		m := cm / 100

		if useJPUnits {
			return fmt.Sprintf("%vm", countingconverter.ConvertToLargestUnit(m))
		}

		return fmt.Sprintf("%vm", stringutils.Commafy(m))
	}
	return fmt.Sprintf("%vcm", stringutils.Commafy(cm))
}

// MtoFt converts meters to ft
func MtoFt(m float64, useJPUnits bool) string {
	ft := m * mToFtRate

	if useJPUnits {
		return fmt.Sprintf("%vft", countingconverter.ConvertToLargestUnit(ft))
	}
	return fmt.Sprintf("%vft", stringutils.Commafy(ft))
}

// FtToM converts ft to meters
func FtToM(ft float64, useJPUnits bool) string {
	m := ft / mToFtRate

	if useJPUnits {
		return fmt.Sprintf("%vm", countingconverter.ConvertToLargestUnit(m))
	}
	return fmt.Sprintf("%vm", stringutils.Commafy(m))
}

// KMtoMi converts kilometers to miles
func KMtoMi(km float64, useJPUnits bool) string {
	mi := km * kmToMiRate

	if useJPUnits {
		return fmt.Sprintf("%vmi", countingconverter.ConvertToLargestUnit(mi))
	}
	return fmt.Sprintf("%vmi", stringutils.Commafy(mi))
}

// MiToKM converts miles to kilometers
func MiToKM(mi float64, useJPUnits bool) string {
	km := mi / kmToMiRate

	if useJPUnits {
		return fmt.Sprintf("%vkm", countingconverter.ConvertToLargestUnit(km))
	}
	return fmt.Sprintf("%vkm", stringutils.Commafy(km))
}
