package distanceconverter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCMToIn(t *testing.T) {
	tests := []struct {
		name       string
		cm         float64
		useJPUnits bool
		want       string
	}{
		{
			name: "converts cm to in",
			cm:   10.0,
			want: "3.94in",
		},
		{
			name: "converts cm to in, appends (ft in) if over 1ft",
			cm:   100.0,
			want: "39.37in (3ft 3.37in)",
		},
		{
			name:       "converts cm to in, appends (ft in) if over 1ft, use jp units true",
			cm:         500000,
			useJPUnits: true,
			want:       "19.69万in (1.64万ft 2.50in)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, CMToIn(tt.cm, tt.useJPUnits))
		})
	}
}

func TestInToCM(t *testing.T) {
	tests := []struct {
		name       string
		in         float64
		useJPUnits bool
		want       string
	}{
		{
			name: "converts in to cm, under 1m returns cm",
			in:   10.0,
			want: "25.40cm",
		},
		{
			name: "converts in to cm, over 1m returns m",
			in:   100.0,
			want: "2.54m",
		},
		{
			name:       "converts in to cm, over 1m returns m, use jp units true",
			in:         1000000,
			useJPUnits: true,
			want:       "2.54万m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, InToCM(tt.in, tt.useJPUnits))
		})
	}
}

func TestMtoFt(t *testing.T) {
	tests := []struct {
		name       string
		m          float64
		useJPUnits bool
		want       string
	}{
		{
			name: "converts m to ft",
			m:    100.0,
			want: "328.08ft",
		},
		{
			name:       "converts m to ft, use jp units flag true",
			m:          10000000,
			useJPUnits: true,
			want:       "3,280.84万ft",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MtoFt(tt.m, tt.useJPUnits))
		})
	}
}

func TestFtToM(t *testing.T) {
	tests := []struct {
		name       string
		ft         float64
		useJPUnits bool
		want       string
	}{
		{
			name: "converts m to ft",
			ft:   100.0,
			want: "30.48m",
		},
		{
			name:       "converts m to ft, use jp units flag true",
			ft:         10000000,
			useJPUnits: true,
			want:       "304.80万m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, FtToM(tt.ft, tt.useJPUnits))
		})
	}
}

func TestKMToMi(t *testing.T) {
	tests := []struct {
		name       string
		km         float64
		useJPUnits bool
		want       string
	}{
		{
			name: "converts m to ft",
			km:   10,
			want: "6.21mi",
		},
		{
			name:       "converts m to ft, use jp units flag true",
			km:         100000,
			useJPUnits: true,
			want:       "6.21万mi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, KMtoMi(tt.km, tt.useJPUnits))
		})
	}
}

func TestMiToKM(t *testing.T) {
	tests := []struct {
		name       string
		mi         float64
		useJPUnits bool
		want       string
	}{
		{
			name: "converts m to ft",
			mi:   10,
			want: "16.09km",
		},
		{
			name:       "converts m to ft, use jp units flag true",
			mi:         100000,
			useJPUnits: true,
			want:       "16.09万km",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MiToKM(tt.mi, tt.useJPUnits))
		})
	}
}
