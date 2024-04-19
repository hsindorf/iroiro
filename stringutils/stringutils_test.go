package stringutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommafy(t *testing.T) {
	tests := []struct {
		name string
		num  float64
		want string
	}{
		// {
		// 	name: "adds commas and truncates to 2",
		// 	num:  1234567.8910,
		// 	want: "1,234,567.89",
		// },
		{
			name: "adds commas and truncates to 2, negative number",
			num:  -100.567,
			want: "-100.57",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Commafy(tt.num))
		})
	}
}
