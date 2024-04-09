package currencyconverter

func DollarsToYen(amount float64, rate float64) float64 {
	return amount * rate
}

func YenToDollars(amount float64, rate float64) float64 {
	return amount / rate
}
