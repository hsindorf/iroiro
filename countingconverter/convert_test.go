package countingconverter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToLargestUnit(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  string
	}{
		{
			name:  "over 兆, resulting whole number",
			input: 2_000_000_000_000.0,
			want:  "2兆",
		},
		{
			name:  "over 兆, resulting float",
			input: 2_500_000_000_000.0,
			want:  "2.5兆",
		},
		{
			name:  "over 億, resulting whole number",
			input: 200_000_000.0,
			want:  "2億",
		},
		{
			name:  "over 億, resulting float",
			input: 250_000_000.0,
			want:  "2.5億",
		},
		{
			name:  "over 万, resulting whole number",
			input: 20_000.0,
			want:  "2万",
		},
		{
			name:  "over 万, resulting float",
			input: 25_000.0,
			want:  "2.5万",
		},
		{
			name:  "under 万, resulting whole number",
			input: 1_000,
			want:  "1,000",
		},
		{
			name:  "under 万, resulting float",
			input: 1_000.5,
			want:  "1,000.5",
		},
		{
			name:  "truncates a long float",
			input: 111_111,
			want:  "11.11万",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertToLargestUnit(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
