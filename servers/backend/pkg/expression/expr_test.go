package expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		formula string
		x       float64
		want    float64
		wantErr bool
	}{
		{"X + 2", 3, 5, false},
		{"X * 2", 4, 8, false},
		{"X / 2", 6, 3, false},
		{"X - 2", 5, 3, false},
		{"X + Y", 3, 0, true}, // Invalid formula, should return an error
		{
			"X>=0.03&&X<=0.45*0.15^0.5?RectangularWaterFlow(0.5,0.15,0.15,X):0.0",
			0.03,
			0.0856963387712763,
			false,
		},
		{
			"X>=0.03&&X<=0.45*0.15^0.5?RectangularWaterFlow(0.5,0.15,0.15,X):0.0",
			0.02,
			0,
			false,
		},
		{
			"RectangularWaterFlow(0.5,0.15,0.15,X)",
			0.6,
			0.5,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.formula, func(t *testing.T) {
			got, err := Calculate(tt.formula, tt.x)
			require.Equal(t, tt.wantErr, err != nil)
			if err == nil {
				t.Log(got)
				assert.InDelta(t, tt.want, got, 0.001)
			} else {
				t.Logf("Error for formula %s with X=%f: %v", tt.formula, tt.x, err)
			}
		})
	}
}
