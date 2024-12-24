package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		s1 []int
		s2 []int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 []int
	}{
		{
			name: "test1",
			args: args{
				s1: []int{65, 3, 58, 678, 64},
				s2: []int{64, 2, 3, 43},
			},
			want:  true,
			want1: []int{3, 64},
		},
		{
			name: "test2",
			args: args{
				s1: []int{1, 2, 3, 4, 5},
				s2: []int{6, 7},
			},
			want:  false,
			want1: []int{},
		},
		{
			name: "test2",
			args: args{
				s1: []int{5, 9, 7, 21},
				s2: []int{6, 3, 56, 9, 95},
			},
			want:  true,
			want1: []int{9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := intersection(tt.args.s1, tt.args.s2)
			if got != tt.want {
				t.Errorf("intersection() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				fmt.Println(got1, len(got1), cap(got1))
				fmt.Println(tt.want1, len(tt.want1), cap(tt.want1))

				t.Errorf("intersection() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
