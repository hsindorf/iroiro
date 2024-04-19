package unitconverter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name       string
		amount     string
		rate       float64
		useJPUnits bool
		want       string
		wantErr    string
	}{
		{
			name:   "only a number, bare number, should convert to jp",
			amount: "100000",
			want:   "10万",
		},
		{
			name:   "only a number, contains commas, should convert to jp",
			amount: "100,000",
			want:   "10万",
		},
		{
			name:   "only a number, contains k, should convert to jp",
			amount: "100k",
			want:   "10万",
		},
		{
			name:   "only a number, contains words, should convert to jp",
			amount: "100 thousand",
			want:   "10万",
		},
		{
			name:   "only a number, contains jp units, should convert to decimal",
			amount: "10万",
			want:   "100,000",
		},
		{
			name:   "only a number, contains jp units, handles a large number without converting into funky stuff",
			amount: "1兆",
			want:   "1,000,000,000,000",
		},
		{
			name:       "currency, parses $, bare number, use jp units false",
			amount:     "$100",
			rate:       100.0,
			useJPUnits: false,
			want:       "10,000円",
		},
		{
			name:       "currency, parses $, bare number, use jp units true",
			amount:     "$100",
			rate:       100.0,
			useJPUnits: true,
			want:       "1万円",
		},
		{
			name:       "currency, parses $, contains commas, use jp units false",
			amount:     "$1,000",
			rate:       100.0,
			useJPUnits: false,
			want:       "100,000円",
		},
		{
			name:       "currency, parses $, contains commas, use jp units true",
			amount:     "$1,000",
			rate:       100.0,
			useJPUnits: true,
			want:       "10万円",
		},
		{
			name:       "currency, parses $, contains k, use jp units false",
			amount:     "$1k",
			rate:       100.0,
			useJPUnits: false,
			want:       "100,000円",
		},
		{
			name:       "currency, parses $, contains k, use jp units true",
			amount:     "$1k",
			rate:       100.0,
			useJPUnits: true,
			want:       "10万円",
		},
		{
			name:       "currency, parses $, contains words, use jp units false",
			amount:     "1 thousand dollars",
			rate:       100.0,
			useJPUnits: false,
			want:       "100,000円",
		},
		{
			name:       "currency, parses $, contains words, use jp units true",
			amount:     "1 thousand dollars",
			rate:       100.0,
			useJPUnits: true,
			want:       "10万円",
		},
		{
			name:       "currency, parses $, contains jp units, use jp units false",
			amount:     "$1万",
			rate:       100.0,
			useJPUnits: false,
			want:       "1,000,000円",
		},
		{
			name:       "currency, parses $, contains jp units, use jp units true",
			amount:     "$1万",
			rate:       100.0,
			useJPUnits: true,
			want:       "100万円",
		},
		{
			name:       "currency, parses 円, bare number, use jp units false",
			amount:     "1000000 yen",
			rate:       100.0,
			useJPUnits: false,
			want:       "$10,000",
		},
		{
			name:       "currency, parses 円, bare number, use jp units true",
			amount:     "1000000 yen",
			rate:       100.0,
			useJPUnits: true,
			want:       "$1万",
		},
		{
			name:       "currency, parses 円, contains commas, use jp units false",
			amount:     "1,000,000 yen",
			rate:       100.0,
			useJPUnits: false,
			want:       "$10,000",
		},
		{
			name:       "currency, parses 円, contains commas, use jp units true",
			amount:     "1,000,000 yen",
			rate:       100.0,
			useJPUnits: true,
			want:       "$1万",
		},
		{
			name:       "currency, parses 円, contains k, use jp units false",
			amount:     "1000k yen",
			rate:       100.0,
			useJPUnits: false,
			want:       "$10,000",
		},
		{
			name:       "currency, parses 円, contains k, use jp units true",
			amount:     "1000k yen",
			rate:       100.0,
			useJPUnits: true,
			want:       "$1万",
		},
		{
			name:       "currency, parses 円, contains words, use jp units false",
			amount:     "1 million yen",
			rate:       100.0,
			useJPUnits: false,
			want:       "$10,000",
		},
		{
			name:       "currency, parses 円, contains words, use jp units true",
			amount:     "1 million yen",
			rate:       100.0,
			useJPUnits: true,
			want:       "$1万",
		},
		{
			name:       "currency, parses 円, contains jp units, use jp units false",
			amount:     "100万円",
			rate:       100.0,
			useJPUnits: false,
			want:       "$10,000",
		},
		{
			name:       "currency, parses 円, contains jp units, use jp units true",
			amount:     "100万円",
			rate:       100.0,
			useJPUnits: true,
			want:       "$1万",
		},
		{
			name:   "units, parses cm",
			amount: "10cm",
			want:   "3.94in",
		},
		{
			name:   "units, parses cm, output greater than 1ft should provide ft'in\"",
			amount: "100cm",
			want:   "39.37in (3ft 3.37in)",
		},
		{
			name:   "units, parses in",
			amount: "10in",
			want:   "25.40cm",
		},
		{
			name:   "units, parses in, output greater than 1m should be in m",
			amount: "100in",
			want:   "2.54m",
		},
		{
			name:   "units, parses m",
			amount: "100m",
			want:   "328.08ft",
		},
		{
			name:   "units, parses ft",
			amount: "100ft",
			want:   "30.48m",
		},
		{
			name:   "units, parses km",
			amount: "10km",
			want:   "6.21mi",
		},
		{
			name:   "units, parses mi",
			amount: "10mi",
			want:   "16.09km",
		},
		{
			name:   "temperature, parses c",
			amount: "10c",
			want:   "50°F",
		},
		{
			name:   "temperature, parses f",
			amount: "10f",
			want:   "-12.22°C",
		},
		{
			name:   "weight, parses kg",
			amount: "10kg",
			want:   "22.05lbs",
		},
		{
			name:   "weight, parses lbs",
			amount: "10lbs",
			want:   "4.54kg",
		},
		{
			name:    "error if both currencies present",
			amount:  "$100円",
			want:    "",
			wantErr: "unable to match supported number patterns",
		},
		{
			name:    "error if not a number",
			amount:  "asdfas",
			want:    "",
			wantErr: "unable to match supported number patterns",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Convert(tt.amount, tt.rate, tt.useJPUnits)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_ParseAmount(t *testing.T) {
	tests := []struct {
		name         string
		amount       string
		wantCurrency string
		wantAmount   string
	}{
		{
			name:         "no currency",
			amount:       "100",
			wantCurrency: "",
			wantAmount:   "100",
		},
		{
			name:         "dollars as word",
			amount:       "100 dollars",
			wantCurrency: "$",
			wantAmount:   "100",
		},
		{
			name:         "dollars as symbol",
			amount:       "$100",
			wantCurrency: "$",
			wantAmount:   "100",
		},
		{
			name:         "dollars as fat symbol",
			amount:       "＄100",
			wantCurrency: "$",
			wantAmount:   "100",
		},
		{
			name:         "dollars as japanese word",
			amount:       "100ドル",
			wantCurrency: "$",
			wantAmount:   "100",
		},
		{
			name:         "yen as word",
			amount:       "100 yen",
			wantCurrency: "円",
			wantAmount:   "100",
		},
		{
			name:         "yen as kanji",
			amount:       "100円",
			wantCurrency: "円",
			wantAmount:   "100",
		},
		{
			name:         "cm",
			amount:       "1cm",
			wantCurrency: "cm",
			wantAmount:   "1",
		},
		{
			name:         "in",
			amount:       "1in",
			wantCurrency: "in",
			wantAmount:   "1",
		},
		{
			name:         "m",
			amount:       "1m",
			wantCurrency: "m",
			wantAmount:   "1",
		},
		{
			name:         "ft",
			amount:       "1ft",
			wantCurrency: "ft",
			wantAmount:   "1",
		},
		{
			name:         "km",
			amount:       "1km",
			wantCurrency: "km",
			wantAmount:   "1",
		},
		{
			name:         "mi",
			amount:       "1mi",
			wantCurrency: "mi",
			wantAmount:   "1",
		},
		{
			name:         "c",
			amount:       "1c",
			wantCurrency: "c",
			wantAmount:   "1",
		},
		{
			name:         "f",
			amount:       "1f",
			wantCurrency: "f",
			wantAmount:   "1",
		},
		{
			name:         "kg",
			amount:       "1kg",
			wantCurrency: "kg",
			wantAmount:   "1",
		},
		{
			name:         "lbs",
			amount:       "1lbs",
			wantCurrency: "lbs",
			wantAmount:   "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCurrency, gotAmount := ParseAmount(tt.amount)
			assert.Equal(t, tt.wantCurrency, gotCurrency)
			assert.Equal(t, tt.wantAmount, gotAmount)
		})
	}
}
