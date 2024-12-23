package main

import (
	"reflect"
	"testing"
)

func Test_sliceExample(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			want: []int{2, 4, 6, 8},
		},
		{
			name: "2",
			args: args{arr: []int{1, 3, 5, 7, 9}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceExample(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceExample() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addElements(t *testing.T) {
	type args struct {
		arr []int
		x   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				x:   105,
			},
			want: []int{1, 2, 3, 4, 5, 105},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addElements(tt.args.arr, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_copySlice(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{arr: []int{1, 2, 3, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copySlice(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeElement(t *testing.T) {
	type args struct {
		arr []int
		i   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				i:   2,
			},
			want: []int{1, 2, 4, 5},
		},
		{
			name: "2",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				i:   -5,
			},
			want: []int{},
		},
		{
			name: "3",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				i:   10,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.arr, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
