package waterflow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRectangularWaterFlow(t *testing.T) {
	tests := []struct {
		B, b, D, h float64
		expected   float64
		expectErr  bool
	}{
		// {1.0, 0.2, 0.5, 0.1, 0, false},
		{0.4, 0.2, 0.5, 0.1, 0, true},            // B out of range
		{1.0, 0.05, 0.5, 0.1, 0, true},           // b out of range
		{1.0, 0.2, 0.1, 0.1, 0, true},            // D out of range
		{1.0, 0.2, 0.5, 0.5, 0, true},            // h out of range
		{0.5, 0.15, 0.15, 0.03, 0.085696, false}, // valid parameters
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("B=%f, b=%f, D=%f, h=%f", tt.B, tt.b, tt.D, tt.h), func(t *testing.T) {
			result, err := rectangularWaterFlow(tt.B, tt.b, tt.D, tt.h)
			require.Equal(t, tt.expectErr, err != nil)
			if err == nil {
				assert.InDelta(t, tt.expected, result, 0.001, "expected %f, got %f", tt.expected, result)
			}
		})
	}
}
