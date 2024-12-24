package main

import (
	"reflect"
	"testing"
)

func Test_conveyer(t *testing.T) {
	tests := []struct {
		name string

		data []uint8
		want []float64
	}{
		{
			name: "test",
			data: []uint8{2, 3, 4, 5},
			want: []float64{8, 27, 64, 125},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := make(chan uint8)
			out := conveyer(in)

			go func() {
				for _, d := range tt.data {
					in <- d
				}
				close(in)
			}()

			var got []float64
			for v := range out {
				got = append(got, v)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("conveyer() = %v, want %v", got, tt.want)
			}

		})
	}
}
