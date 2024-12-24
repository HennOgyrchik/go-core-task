package main

import (
	"testing"
)

func Test_merge(t *testing.T) {
	type args[T comparable] struct {
		arr []chan T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[int]{
		{
			name: "test 1",
			args: struct{ arr []chan int }{arr: []chan int{make(chan int), make(chan int)}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for i, ch := range tt.args.arr {
				go func() {
					ch <- i
					close(ch)
				}()
			}

			mCh := merge(tt.args.arr)

			count := 0

			for range mCh {
				count++
			}

			if tt.want != count {
				t.Errorf("merge() = %v, want %v", count, tt.want)
			}

		})
	}
}
