package weightconverter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKgToLbs(t *testing.T) {
	tests := []struct {
		name       string
		kg         float64
		useJPUnits bool
		want       string
	}{
		{
			name: "converts kg to lbs",
			kg:   10000,
			want: "22,050lbs",
		},
		{
			name:       "converts m to ft, use jp units flag true",
			kg:         10000,
			useJPUnits: true,
			want:       "2.21万lbs",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, kgToLbs(tt.kg, tt.useJPUnits))
		})
	}
}

func TestLbsToKg(t *testing.T) {
	tests := []struct {
		name       string
		lbs        float64
		useJPUnits bool
		want       string
	}{
		{
			name: "converts lbs to kg",
			lbs:  50000,
			want: "22,675.74kg",
		},
		{
			name:       "converts lbs to kg, use jp units flag true",
			lbs:        50000,
			useJPUnits: true,
			want:       "2.27万kg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, lbsToKg(tt.lbs, tt.useJPUnits))
		})
	}
}
