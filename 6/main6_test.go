package main

import (
	"context"
	"testing"
)

func Test_randGen(t *testing.T) {
	type args struct {
		ctx  context.Context
		from int
		to   int
	}
	tests := []struct {
		name  string
		args  args
		count int
	}{

		{
			name: "test",
			args: args{
				ctx:  context.Background(),
				from: 0,
				to:   10,
			},
			count: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := randGen(tt.args.ctx, tt.args.from, tt.args.to)

			for i := 0; i < tt.count; i++ {
				if x := <-ch; x < tt.args.from || x >= tt.args.to {
					t.Errorf("randGen() = %v, want from %v to %v", x, tt.args.from, tt.args.to)
				}
			}

		})
		tt.args.ctx.Done()
	}
}
