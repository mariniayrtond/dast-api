package model

import (
	"reflect"
	"testing"
)

func TestNewAHPMatrix(t *testing.T) {
	type args struct {
		values [][]float64
	}
	tests := []struct {
		name string
		args args
		want AHPMatrix
	}{
		{
			name: "happy-test-1",
			args: args{
				values: [][]float64{
					{
						1.0,
						0.0,
					},
					{
						0.0,
						1.0,
					},
				},
			},
			want: AHPMatrix{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAHPMatrix(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAHPMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}