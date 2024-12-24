package main

import (
	"reflect"
	"testing"
)

func Test_f(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test",
			args: args{
				s1: []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
				s2: []string{"banana", "date", "fig"},
			},
			want: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("f() = %v, want %v", got, tt.want)
			}
		})
	}
}
