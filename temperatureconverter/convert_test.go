package temperatureconverter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCToF(t *testing.T) {
	tests := []struct {
		name       string
		c          float64
		useJPUnits bool
		want       string
	}{
		{
			name: "converts C to F",
			c:    10000.0,
			want: "18,032°F",
		},
		{
			name:       "converts C to F, use jp units true",
			c:          10000.0,
			useJPUnits: true,
			want:       "1.80万°F",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, cToF(tt.c, tt.useJPUnits))
		})
	}
}

func TestFToC(t *testing.T) {
	tests := []struct {
		name       string
		f          float64
		useJPUnits bool
		want       string
	}{
		{
			name: "converts F to C",
			f:    100000.0,
			want: "55,537.78°C",
		},
		{
			name:       "converts F to C, use jp units true",
			f:          100000.0,
			useJPUnits: true,
			want:       "5.55万°C",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fToC(tt.f, tt.useJPUnits))
		})
	}
}
