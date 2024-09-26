package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRectangularWaterFlowFormula_SettingFormula(t *testing.T) {
	tests := []struct {
		name    string
		params  map[string]any
		want    string
		wantErr bool
	}{
		{
			name: "valid parameters",
			params: map[string]any{
				"B": 1.0,
				"b": 0.5,
				"D": 2.0,
				"N": 3,
			},
			want:    "X>=0.03&&X<=0.45*0.50^0.5?RectangularWaterFlow(1.00, 0.50, 2.00, X)*60*3:0.0",
			wantErr: false,
		},
		{
			name: "missing B",
			params: map[string]any{
				"b": 0.5,
				"D": 2.0,
				"N": 3,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "B is not float64",
			params: map[string]any{
				"B": "1.0",
				"b": 0.5,
				"D": 2.0,
				"N": 3,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "B is less than or equal to 0",
			params: map[string]any{
				"B": 0.0,
				"b": 0.5,
				"D": 2.0,
				"N": 3,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "missing b",
			params: map[string]any{
				"B": 1.0,
				"D": 2.0,
				"N": 3,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "missing D",
			params: map[string]any{
				"B": 1.0,
				"b": 0.5,
				"N": 3,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "missing N",
			params: map[string]any{
				"B": 1.0,
				"b": 0.5,
				"D": 2,
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RectangularWaterFlowFormula{}
			got, err := r.SettingFormula(tt.params)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
