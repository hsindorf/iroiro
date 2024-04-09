package countingconverter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr string
	}{
		{
			name:  "parses a number using japanese characters",
			input: "１２３４５６７８９０",
			want:  1234567890,
		},
		{
			name:  "parses an int",
			input: "1000000",
			want:  1_000_000.0,
		},
		{
			name:  "parses a float",
			input: "1.5",
			want:  1.5,
		},
		{
			name:  "parses a float less than 1",
			input: ".5",
			want:  0.5,
		},
		{
			name:  "parses an int with commas",
			input: "1,000,000",
			want:  1_000_000.0,
		},
		{
			name:  "parses a float with commas",
			input: "1,500.5",
			want:  1_500.5,
		},
		{
			name:  "parses an int that uses 'k' to denote thousands",
			input: "100k",
			want:  100_000.0,
		},
		{
			name:  "parses a float that uses 'k' to denote thousands",
			input: "1.5k",
			want:  1_500.0,
		},
		{
			name:  "parses a float less than 1 that uses 'k' to denote thousands",
			input: ".5k",
			want:  500.0,
		},
		{
			name:  "parses an int that uses words - all units present",
			input: "1 trillion 2 billion 3 million 4 thousand",
			want:  1_002_003_004_000.0,
		},
		{
			name:  "parses an int that uses words - some units omitted",
			input: "2 billion 4 thousand",
			want:  2_000_004_000.0,
		},
		{
			name:  "parses a float that uses words",
			input: "2.4 billion",
			want:  2_400_000_000.0,
		},
		{
			name:  "parses a float that uses words, with numbers less than 1",
			input: ".5 billion",
			want:  500_000_000.0,
		},
		{
			name:  "parses an int that uses japanese units - all units present",
			input: "1兆2億3万4千",
			want:  1_000_200_034_000.0,
		},
		{
			name:  "parses an int that uses japanese units - some units omitted",
			input: "2万",
			want:  20_000.0,
		},
		{
			name:  "parses a float that uses japanese units",
			input: "1.5万",
			want:  15_000.0,
		},
		{
			name:  "parses a float that uses japanese units, with numbers less than 1",
			input: ".5万",
			want:  5_000.0,
		},
		{
			name:    "rejects non-numbers",
			input:   "asdfdsa",
			want:    0.0,
			wantErr: "unable to match supported number patterns",
		},
		{
			name:    "rejects k format if multiple decimals",
			input:   "1.5.5k",
			want:    0.0,
			wantErr: "unable to match supported number patterns",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.input)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
