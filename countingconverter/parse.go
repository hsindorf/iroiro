package countingconverter

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	kRegexPattern    = "^(\\d*(\\.\\d+){0,1})k$"
	wordRegexPattern = "((\\d*(\\.\\d+){0,1}) (trillion|billion|million|thousand))"
	jpRegexPattern   = "(\\d*(\\.\\d+){0,1})(兆|億|万|千)"
)

// Parse parses a number from a string to a float. Examples:
//   - Float or int: `"1000000" => 1000000`
//   - Float or int with commas: `"1,000,000" => 1000000`
//   - Number that uses words: `"1 million 2 thousand" => 1000200`
//   - Number that uses 'k' to denote thousands: `"100k" => 100000`
//   - Number that uses Japanese units: `"1.23万" => 12300`
func Parse(input string) (float64, error) {
	noCommas := strings.ReplaceAll(input, ",", "")
	noFunkies := replaceFunkies(noCommas)

	parsed, err := strconv.ParseFloat(noFunkies, 64)
	if err == nil {
		return parsed, nil
	}

	kRegex, _ := regexp.Compile(kRegexPattern)
	if matchesK := kRegex.MatchString(noFunkies); matchesK {
		return parseK(noFunkies), nil
	}

	wordRegex, _ := regexp.Compile(wordRegexPattern)
	if matchesWord := wordRegex.MatchString(noFunkies); matchesWord {
		return parseWord(noFunkies), nil
	}

	jpRegex, _ := regexp.Compile(jpRegexPattern)
	if matchesJP := jpRegex.MatchString(noFunkies); matchesJP {
		return parseJP(noFunkies), nil
	}

	return 0.0, errors.New("unable to match supported number patterns")
}

func replaceFunkies(input string) string {
	funkiesMap := map[string]string{
		"１": "1",
		"２": "2",
		"３": "3",
		"４": "4",
		"５": "5",
		"６": "6",
		"７": "7",
		"８": "8",
		"９": "9",
		"０": "0",
	}

	output := strings.Builder{}

	for _, char := range input {
		charStr := string(char)
		if mapped, ok := funkiesMap[charStr]; ok {
			output.WriteString(mapped)
			continue
		}
		output.WriteRune(char)
	}
	return output.String()
}

func parseK(input string) float64 {
	kRegex, _ := regexp.Compile(kRegexPattern)
	found := kRegex.FindAllStringSubmatch(input, -1)
	parsed, _ := strconv.ParseFloat(found[0][1], 64)
	return parsed * 1000.0
}

func parseWord(input string) float64 {
	units := map[string]float64{
		"trillion": 1_000_000_000_000.0,
		"billion":  1_000_000_000.0,
		"million":  1_000_000.0,
		"thousand": 1_000.0,
		"hundred":  100.0,
	}
	wordRegex, _ := regexp.Compile(wordRegexPattern)
	found := wordRegex.FindAllStringSubmatch(input, -1)
	total := 0.0
	for _, val := range found {
		parsed, _ := strconv.ParseFloat(val[2], 64)
		total += parsed * units[val[4]]
	}

	return total
}

func parseJP(input string) float64 {
	units := map[string]float64{
		"兆": 1_000_000_000_000.0,
		"億": 100_000_000.0,
		"万": 10_000.0,
		"千": 1_000.0,
	}
	jpRegex, _ := regexp.Compile(jpRegexPattern)
	found := jpRegex.FindAllStringSubmatch(input, -1)
	total := 0.0
	for _, val := range found {
		parsed, _ := strconv.ParseFloat(val[1], 64)
		total += parsed * units[val[3]]
	}
	return total
}
